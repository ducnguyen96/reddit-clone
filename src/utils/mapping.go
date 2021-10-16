package utils

import (
	"github.com/ducnguyen96/reddit-clone/ent"
	"github.com/ducnguyen96/reddit-clone/graph/model"
)

func MapEntGoUserToGraphUser(u *ent.User) *model.User {
	return &model.User{
		ID:        Uint64ToString(u.ID),
		Username: u.Username,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
	}
}