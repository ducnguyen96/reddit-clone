// Code generated by entc, DO NOT EDIT.

package media

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

const (
	// Label holds the string label denoting the media type in the database.
	Label = "media"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the media in the database.
	Table = "media"
)

// Columns holds all SQL columns for media fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldURL,
	FieldType,
}

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
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enums.MediaType) error {
	switch _type.String() {
	case "Image", "Video":
		return nil
	default:
		return fmt.Errorf("media: invalid enum value for type field: %q", _type)
	}
}