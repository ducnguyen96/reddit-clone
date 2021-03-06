// Code generated by entc, DO NOT EDIT.

package community

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

const (
	// Label holds the string label denoting the community type in the database.
	Label = "community"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsAdult holds the string denoting the is_adult field in the database.
	FieldIsAdult = "is_adult"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the community in the database.
	Table = "communities"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "community_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// AdminsTable is the table that holds the admins relation/edge. The primary key declared below.
	AdminsTable = "community_admins"
	// AdminsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminsInverseTable = "users"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "community_posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
)

// Columns holds all SQL columns for community fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldSlug,
	FieldType,
	FieldIsAdult,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"community_id", "user_id"}
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"community_id", "user_id"}
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"community_id", "post_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// DefaultIsAdult holds the default value on creation for the "is_adult" field.
	DefaultIsAdult bool
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.CommunityType) error {
	switch _type.String() {
	case "Public", "Restricted", "Private":
		return nil
	default:
		return fmt.Errorf("community: invalid enum value for type field: %q", _type)
	}
}
