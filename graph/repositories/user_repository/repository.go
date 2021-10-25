package user_repository

import (
	"context"
	"fmt"
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/ent/action"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
	"github.com/ducnguyen96/reddit-clone/ent/user"
	"github.com/ducnguyen96/reddit-clone/graph/model"
	"github.com/ducnguyen96/reddit-clone/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	readDB  *ent.Client
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

func (r *UserRepository) FindById(ctx context.Context, id uint64) *ent.User {
	return r.readDB.User.Query().Where().Where(user.ID(id)).FirstX(ctx)
}

func (r *UserRepository) CreateTx(ctx context.Context, input model.UserRegisterInput) (*ent.User, *ent.Tx, error) {
	tx, err := r.writeDB.Tx(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// create hashed password
	// Tham khảo tại https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
	pwd := []byte(input.Password)
	hashedPwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return nil, nil, fmt.Errorf("failed hashing password: %w", err)
	}

	u, err := tx.User.Create().SetUsername(input.Username).SetPassword(string(hashedPwd)).Save(ctx)

	if err != nil {
		return nil, nil, err
	}
	return u, tx, nil
}

func (r *UserRepository) CreateAction(ctx context.Context, usr ent.User, input model.UserCreateActionInput) (*ent.Action, error) {
	var isUpVoted, isDownVoted = false, false
	postId := utils.StringToUint64(input.Target)
	po := r.readDB.Post.Query().Where(post.ID(postId)).FirstX(ctx)

	// validate action
	at := r.readDB.User.QueryActions(&usr).
		Where(action.TargetTypeEQ(enums.POST)).
		Where(action.Target(postId)).
		FirstX(ctx)

	if at != nil && at.Type == enums.UpVote {
		isUpVoted = true
	}

	if at != nil && at.Type == enums.DownVote {
		isDownVoted = true
	}

	// if action existed
	if isUpVoted && input.Type == model.UserActionTypeUpVote || isDownVoted && input.Type == model.UserActionTypeDownVote {
		return nil, nil
	}

	// starting a tx
	tx, err := r.writeDB.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	// Delete old action
	if at != nil {
		err = tx.Action.DeleteOne(at).Exec(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed deleting old action: %w", err)
		}
	}

	// Update post
	if isUpVoted == true {
		po, err = tx.Post.UpdateOne(po).SetUpVotes(po.UpVotes - 1).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating post: %w", err)
		}
	}

	if isDownVoted == true {
		po, err = tx.Post.UpdateOne(po).SetDownVotes(po.DownVotes - 1).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating post: %w", err)
		}
	}

	// create new action
	newAt, err := tx.Action.Create().
		SetType(utils.GraphUserActionTypeToEnt(input.Type)).
		SetTargetType(utils.GraphUserActionTargetTypeToEnt(input.TargetType)).
		SetTarget(utils.StringToUint64(input.Target)).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating new action: %w", err)
	}

	// update post
	if input.Type == model.UserActionTypeUpVote {
		_, err = tx.Post.UpdateOne(po).SetUpVotes(po.UpVotes + 1).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating post after action: %w", err)
		}
	}

	if input.Type == model.UserActionTypeDownVote {
		_, err = tx.Post.UpdateOne(po).SetDownVotes(po.DownVotes + 1).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed updating post after action: %w", err)
		}
	}

	// add action to user
	_, err = tx.User.Update().Where(user.ID(usr.ID)).AddActions(newAt).Save(ctx)
	

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("failed adding action to user tx: %w", err)
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed commiting tx: %w", err)
	}

	return newAt, nil
}

func (r *UserRepository) GetUserActionStatusForPost(ctx context.Context, postId uint64, usr *ent.User) (bool, bool) {

	at := r.readDB.User.QueryActions(usr).
		Where(action.TargetTypeEQ(enums.POST)).
		Where(action.Target(postId)).
		FirstX(ctx)

	if at == nil {
		return false, false
	} else {
		if at.Type == enums.UpVote {
			return true, false
		} else {
			return false, true
		}
	}
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
