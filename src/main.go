package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ducnguyen96/reddit-clone/ent"
	_ "github.com/ducnguyen96/reddit-clone/ent/runtime"
	"github.com/ducnguyen96/reddit-clone/graph/directives"
	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/user_repository"
	graph "github.com/ducnguyen96/reddit-clone/graph/resolver"
	auth_services "github.com/ducnguyen96/reddit-clone/graph/services/auth_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/user_services"
	"github.com/ducnguyen96/reddit-clone/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func init() {
	defaultTranslation()
}

func defaultTranslation() {
	//directives.ValidateAddTranslation("email", " not a valid email (custom message)")
}

func main() {
	// pg database
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort) //Build connection string

	readClient, err := ent.Open("postgres", dbUri)
	if err != nil {
		log.Fatal(err)
	}

	//writeClient, err := ent.Open("postgres", dbUri)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Run the auto migration tool.
	if err := readClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(readClient)

	// Set up a http server.
	r := gin.Default()
	r.Use(utils.GinContextToContextMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to reddit clone")
	})

	// Repositories
	userRepository := user_repository.NewUserRepository(readClient, readClient)

	// Services
	userService := user_services.NewUserService(userRepository)
	authService := auth_services.NewAuthService()

	gr := generated.Config{Resolvers: &graph.Resolver{
		UerService: userService,
		AuthService: authService,
	}}

	gr.Directives.Binding = directives.Binding
	h := handler.NewDefaultServer(generated.NewExecutableSchema(gr))

	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/graphql", func() gin.HandlerFunc {
		h := playground.Handler("GraphQL", "/graphql")
		return func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		}
	}())

	if err := r.Run(); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}