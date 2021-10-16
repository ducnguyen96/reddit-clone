package user_services

import (
	"context"
	"fmt"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/graph/repositories/user_repository"
)

type UserService struct {
	repository *user_repository.UserRepository
}

func NewUserService(repository *user_repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func(s *UserService) CreateUserTransaction(ctx context.Context, userRegisterInput model.UserRegisterInput) (*ent.User, *ent.Tx, error) {
	// validate username exist
	u := s.repository.FindByUserName(ctx, userRegisterInput.Username)

	if u != nil {
		return  nil, nil, fmt.Errorf("username is existed")
	}

	user, transaction, err := s.repository.CreateTx(ctx, userRegisterInput)
	return user, transaction, err
}

// Rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func(s *UserService) Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}