// Code generated by entc, DO NOT EDIT.

package post

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldContentMode holds the string denoting the content_mode field in the database.
	FieldContentMode = "content_mode"
	// FieldUpVotes holds the string denoting the up_votes field in the database.
	FieldUpVotes = "up_votes"
	// FieldDownVotes holds the string denoting the down_votes field in the database.
	FieldDownVotes = "down_votes"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeCommunity holds the string denoting the community edge name in mutations.
	EdgeCommunity = "community"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// OwnerTable is the table that holds the owner relation/edge. The primary key declared below.
	OwnerTable = "user_posts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// CommunityTable is the table that holds the community relation/edge. The primary key declared below.
	CommunityTable = "community_posts"
	// CommunityInverseTable is the table name for the Community entity.
	// It exists in this package in order to avoid circular dependency with the "community" package.
	CommunityInverseTable = "communities"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "post_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// CommentsTable is the table that holds the comments relation/edge. The primary key declared below.
	CommentsTable = "post_comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldSlug,
	FieldContent,
	FieldType,
	FieldContentMode,
	FieldUpVotes,
	FieldDownVotes,
}

var (
	// OwnerPrimaryKey and OwnerColumn2 are the table columns denoting the
	// primary key for the owner relation (M2M).
	OwnerPrimaryKey = []string{"user_id", "post_id"}
	// CommunityPrimaryKey and CommunityColumn2 are the table columns denoting the
	// primary key for the community relation (M2M).
	CommunityPrimaryKey = []string{"community_id", "post_id"}
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"post_id", "tag_id"}
	// CommentsPrimaryKey and CommentsColumn2 are the table columns denoting the
	// primary key for the comments relation (M2M).
	CommentsPrimaryKey = []string{"post_id", "comment_id"}
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// DefaultUpVotes holds the default value on creation for the "up_votes" field.
	DefaultUpVotes int
	// DefaultDownVotes holds the default value on creation for the "down_votes" field.
	DefaultDownVotes int
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.PostType) error {
	switch _type.String() {
	case "Post", "Image_Video", "Link":
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for type field: %q", _type)
	}
}

// ContentModeValidator is a validator for the "content_mode" field enum values. It is called by the builders before save.
func ContentModeValidator(cm enums.InputContentMode) error {
	switch cm.String() {
	case "MarkDown", "TextEditor":
		return nil
	default:
		return fmt.Errorf("post: invalid enum value for content_mode field: %q", cm)
	}
}