// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAvatarURL holds the string denoting the avatar_url field in the database.
	FieldAvatarURL = "avatar_url"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeCommunities holds the string denoting the communities edge name in mutations.
	EdgeCommunities = "communities"
	// EdgeMyCommunities holds the string denoting the my_communities edge name in mutations.
	EdgeMyCommunities = "my_communities"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CommunitiesTable is the table that holds the communities relation/edge. The primary key declared below.
	CommunitiesTable = "community_users"
	// CommunitiesInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	CommunitiesInverseTable = "communities"
	// MyCommunitiesTable is the table that holds the my_communities relation/edge. The primary key declared below.
	MyCommunitiesTable = "community_admins"
	// MyCommunitiesInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	MyCommunitiesInverseTable = "communities"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "user_posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
	// CommentsTable is the table that holds the comments relation/edge. The primary key declared below.
	CommentsTable = "user_comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldEmail,
	FieldAvatarURL,
	FieldPassword,
}

var (
	// CommunitiesPrimaryKey and CommunitiesColumn2 are the table columns denoting the
	// primary key for the communities relation (M2M).
	CommunitiesPrimaryKey = []string{"community_id", "user_id"}
	// MyCommunitiesPrimaryKey and MyCommunitiesColumn2 are the table columns denoting the
	// primary key for the my_communities relation (M2M).
	MyCommunitiesPrimaryKey = []string{"community_id", "user_id"}
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"user_id", "post_id"}
	// CommentsPrimaryKey and CommentsColumn2 are the table columns denoting the
	// primary key for the comments relation (M2M).
	CommentsPrimaryKey = []string{"user_id", "comment_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/ducnguyen96/reddit-clone/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
