package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
)

func (r *mutationResolver) Register(ctx context.Context, input model.UserRegisterInput) (model.RegisterResult, error) {
	pwd, rpwd := input.Password, input.RepeatPassword

	if pwd != rpwd {
		customErr := &model.CustomError{
			Message: "password and repeatPassword not match",
			Path:    "Register",
		}
		return &model.RegisterBadRequest{Errors: []*model.CustomError{customErr}}, nil
	}

	user, transaction, err := r.UserService.CreateUserTransaction(ctx, input)
	if err != nil {
		customErr := &model.CustomError{
			Message: fmt.Sprintf("%v", err),
			Path:    "Register",
		}
		return &model.RegisterBadRequest{Errors: []*model.CustomError{customErr}}, nil
	}

	token, err := r.AuthService.CreateToken(user.ID)

	if err != nil {
		err := r.UserService.Rollback(transaction, fmt.Errorf("failed creating token: %w", err))
		return &model.RegisterInternalServerError{
			Error: &model.CustomError{
				Message: fmt.Sprintf("%v", err),
				Path:    "Register",
			},
		}, nil
	}

	err = transaction.Commit()

	if err != nil {
		return &model.RegisterInternalServerError{
			Error: &model.CustomError{
				Message: fmt.Sprintf("%v", err),
				Path:    "Register",
			},
		}, nil
	}

	return &model.RegisterPayload{
		User: utils.MapEntGoUserToGraphUser(user),
		Token: &model.TokenPayloadDto{
			ExpiresIn:   token.ExpiresIn,
			AccessToken: token.AccessToken,
		},
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.UserLoginInput) (*model.TokenPayloadDto, error) {
	usr, err := r.UserService.Login(ctx, input)

	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.CreateToken(usr.ID)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *mutationResolver) SignOut(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginGoogle(ctx context.Context, gmail string) (*model.TokenPayloadDto, error) {
	// find if existed
	username := strings.Split(gmail, "@gmail.com")[0]
	usr := r.UserService.FindUserByUsername(ctx, username)

	if usr != nil {
		token, err := r.AuthService.CreateToken(usr.ID)

		if err != nil {
			return nil, err
		}

		return token, nil
	}

	// usr == nil
	pw := utils.GeneratePassword(20, 2, 2, 2)

	user, transaction, err := r.UserService.CreateUserTransaction(ctx, model.UserRegisterInput{Username: username, Password: pw, RepeatPassword: pw})
	if err != nil {
		return nil, err
	}

	token, err := r.AuthService.CreateToken(user.ID)

	if err != nil {
		err := r.UserService.Rollback(transaction, fmt.Errorf("failed creating token: %w", err))
		return nil, err
	}

	err = transaction.Commit()

	if err != nil {
		return nil, err
	}

	return token, nil
}
