package user_services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/user_repository"
	"github.com/ducnguyen96/reddit-clone/graph/services/auth_services"
	"github.com/ducnguyen96/reddit-clone/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *user_repository.UserRepository
}

func NewUserService(repository *user_repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUserTransaction(ctx context.Context, userRegisterInput model.UserRegisterInput) (*ent.User, *ent.Tx, error) {
	// validate username exist
	u := s.repository.FindByUserName(ctx, userRegisterInput.Username)

	if u != nil {
		return nil, nil, fmt.Errorf("username is existed")
	}

	usr, transaction, err := s.repository.CreateTx(ctx, userRegisterInput)
	return usr, transaction, err
}

func (s *UserService) GetCurrentUser(ctx context.Context) (*ent.User, error) {
	token := utils.GetAuthToken(ctx)
	usr, err := s.GetUserByToken(ctx, token)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	if usr == nil {
		return nil, errors.New("unauthorized")
	}
	return usr, nil
}

func (s *UserService) GetCurrentUserNoTokenValid(ctx context.Context) (*ent.User, error) {
	token := utils.GetAuthToken(ctx)
	if len(token) < 10 {
		return nil, nil
	}
	return s.GetUserByToken(ctx, token)
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (*ent.User, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		return nil, errors.New("invalid token")
	}
	token = strings.TrimPrefix(token, "Bearer ")

	parsedToken, err := jwt.ParseWithClaims(token, &auth_services.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return auth_services.MySigningKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	if claims, ok := parsedToken.Claims.(*auth_services.MyCustomClaims); ok && parsedToken.Valid {
		usr := s.repository.FindById(ctx, claims.UserId)
		return usr, nil
	} else {
		return nil, err
	}
}

func (s *UserService) Login(ctx context.Context, input model.UserLoginInput) (*ent.User, error) {
	usr := s.repository.FindByUserName(ctx, input.Username)
	if usr == nil {
		return nil, errors.New("invalid Input")
	}

	hashedPw := []byte(usr.Password)
	err := bcrypt.CompareHashAndPassword(hashedPw, []byte(input.Password))

	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid Input")
	}
	return usr, nil
}

func (s *UserService) CreateAction(ctx context.Context, usr ent.User, input model.UserCreateActionInput) (*ent.Action, error) {
	return s.repository.CreateAction(ctx, usr, input)
}

func (s *UserService) GetUserActionStatusForPost(ctx context.Context, postId uint64, usr *ent.User) (bool, bool) {
	return s.repository.GetUserActionStatusForPost(ctx, postId, usr)
}

func (s *UserService) FindUserByUsername(ctx context.Context, username string) *ent.User {
	return s.repository.FindByUserName(ctx, username)
}

// Rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func (s *UserService) Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
