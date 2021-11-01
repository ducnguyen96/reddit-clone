package auth_services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	_ "github.com/joho/godotenv/autoload"
)

var MySigningKey = []byte(os.Getenv("AUTH_SECRET_KEY"))

type MyCustomClaims struct {
	UserId uint64 `json:"user_id"`
	jwt.StandardClaims
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) CreateToken(userId uint64) (*model.TokenPayloadDto, error) {
	// Create the Claims
	// Tham khảo https://pkg.go.dev/github.com/golang-jwt/jwt#example-NewWithClaims-CustomClaimsType
	expiresIn, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRES_IN"))
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiresIn)).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)

	if err != nil {
		return nil, fmt.Errorf("failed signing token: %w", err)
	}

	return &model.TokenPayloadDto{
		ExpiresIn:   &expiresIn,
		AccessToken: &ss,
	}, nil
}
