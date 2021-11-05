package main

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ducnguyen96/reddit-clone/ent"
	_ "github.com/ducnguyen96/reddit-clone/ent/runtime"
	"github.com/ducnguyen96/reddit-clone/graph/directives"
	"github.com/ducnguyen96/reddit-clone/graph/generated"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/comment_repository"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/community_repository"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/media_repository"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/post_repository"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/user_repository"
	graph "github.com/ducnguyen96/reddit-clone/graph/resolver"
	"github.com/ducnguyen96/reddit-clone/graph/services/auth_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/comment_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/community_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/media_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/post_services"
	"github.com/ducnguyen96/reddit-clone/graph/services/user_services"
	"github.com/ducnguyen96/reddit-clone/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func init() {
	defaultTranslation()
}

func defaultTranslation() {
	//directives.ValidateAddTranslation("email", " not a valid email (custom message)")
}

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://redditclone.ducnguyen96.xyz", "http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "Authorization")
	r.Use(cors.New(corsConfig))

	// Repositories
	userRepo := user_repository.NewUserRepository(readClient, readClient)
	communityRepo := community_repository.NewCommunityRepository(readClient, readClient)
	postRepo := post_repository.NewPostRepository(readClient, readClient)
	commentRepo := comment_repository.NewCommentRepository(readClient, readClient)
	mediaRepo := media_repository.NewMediaRepository(readClient, readClient)

	// Services
	userService := user_services.NewUserService(userRepo)
	authService := auth_services.NewAuthService()
	communityService := community_services.NewCommunityService(communityRepo)
	postService := post_services.NewPostService(postRepo)
	commentService := comment_services.NewCommentService(commentRepo)
	mediaService := media_services.NewMediaService(mediaRepo)

	gr := generated.Config{Resolvers: &graph.Resolver{
		UserService:      userService,
		AuthService:      authService,
		CommunityService: communityService,
		PostService:      postService,
		CommentService:   commentService,
		MediaService:     mediaService,
	}}

	gr.Directives.Binding = directives.Binding
	h := handler.NewDefaultServer(generated.NewExecutableSchema(gr))

	r.POST("/", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.PUT("/succeed", func(c *gin.Context) {
		uri := c.Request.RequestURI
		id := uri[12:]
		cmd := exec.Command("/ffmpeg.sh", id)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Print the output
		fmt.Println(string(stdout))
	})

	r.POST("/upload", func(c *gin.Context) {
		var form Form
		_ = c.ShouldBind(&form)

		// Validate inputs
		file := form.File

		valid, message := utils.ValidateUploadFile(file)
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
				"path":    "validateUploadFiles(form.Files)",
			})
			return
		}

		contentType := file.Header.Get("Content-Type")

		isImage := strings.HasPrefix(contentType, "image")
		isVideo := strings.HasPrefix(contentType, "video")

		unique := uuid.New().String()

		fn := strings.ReplaceAll(file.Filename, " ", "_")

		var path string
		if isImage {
			path = "/images/" + unique + "-" + fn
		}

		if isVideo {
			path = "/videos/raw/" + unique + "-" + fn
		}

		_ = c.SaveUploadedFile(file, "/media"+path)

		c.JSON(http.StatusOK, gin.H{
			"url":          path,
			"content-type": contentType,
		})
	})

	r.GET("/", func() gin.HandlerFunc {
		h := playground.Handler("GraphQL", "/")
		return func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		}
	}())

	if err := r.Run(":5000"); err != nil { // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		panic("Error")
	}
}
