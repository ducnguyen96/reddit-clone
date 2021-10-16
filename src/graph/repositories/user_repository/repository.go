package user_repository

import (
	"context"
	"fmt"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/user"
	"github.com/ducnguyen96/reddit-clone/graph/model"
)

type UserRepository struct {
	readDB *ent.Client
	writeDB *ent.Client
}

func NewUserRepository(readDB *ent.Client, writeDB *ent.Client) *UserRepository {
	return &UserRepository{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (r *UserRepository) Create(ctx context.Context, input model.UserRegisterInput) (*ent.User, error) {
	u, err := r.writeDB.User.Create().SetUsername(input.Username).Save(ctx)
	return u, err
}

func (r *UserRepository) FindByUserName(ctx context.Context, username string) *ent.User {
	return r.readDB.User.Query().Where().Where(user.Username(username)).FirstX(ctx)
}

func (r *UserRepository) CreateTx(ctx context.Context, input model.UserRegisterInput) (*ent.User, *ent.Tx, error) {
	tx, err := r.writeDB.Tx(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("starting a transaction: %w", err)
	}

	u, err := tx.User.Create().SetUsername(input.Username).Save(ctx)

	if err != nil {
		return nil, nil, err
	}
	return u, tx, nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}