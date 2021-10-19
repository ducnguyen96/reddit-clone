// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ducnguyen96/reddit-clone/ent/community"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

// Community is the model entity for the Community schema.
type Community struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Type holds the value of the "type" field.
	Type enums.CommunityType `json:"type,omitempty"`
	// IsAdult holds the value of the "is_adult" field.
	IsAdult bool `json:"is_adult,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommunityQuery when eager-loading is set.
	Edges CommunityEdges `json:"edges"`
}

// CommunityEdges holds the relations/edges for other nodes in the graph.
type CommunityEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Admins holds the value of the admins edge.
	Admins []*User `json:"admins,omitempty"`
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e CommunityEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[2] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Community) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case community.FieldType:
			values[i] = new(enums.CommunityType)
		case community.FieldIsAdult:
			values[i] = new(sql.NullBool)
		case community.FieldID:
			values[i] = new(sql.NullInt64)
		case community.FieldName, community.FieldSlug:
			values[i] = new(sql.NullString)
		case community.FieldCreatedAt, community.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Community", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Community fields.
func (c *Community) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case community.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case community.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case community.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case community.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case community.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				c.Slug = value.String
			}
		case community.FieldType:
			if value, ok := values[i].(*enums.CommunityType); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value != nil {
				c.Type = *value
			}
		case community.FieldIsAdult:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_adult", values[i])
			} else if value.Valid {
				c.IsAdult = value.Bool
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Community entity.
func (c *Community) QueryUsers() *UserQuery {
	return (&CommunityClient{config: c.config}).QueryUsers(c)
}

// QueryAdmins queries the "admins" edge of the Community entity.
func (c *Community) QueryAdmins() *UserQuery {
	return (&CommunityClient{config: c.config}).QueryAdmins(c)
}

// QueryPosts queries the "posts" edge of the Community entity.
func (c *Community) QueryPosts() *PostQuery {
	return (&CommunityClient{config: c.config}).QueryPosts(c)
}

// Update returns a builder for updating this Community.
// Note that you need to call Community.Unwrap() before calling this method if this Community
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Community) Update() *CommunityUpdateOne {
	return (&CommunityClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Community entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Community) Unwrap() *Community {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Community is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Community) String() string {
	var builder strings.Builder
	builder.WriteString("Community(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", slug=")
	builder.WriteString(c.Slug)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteString(", is_adult=")
	builder.WriteString(fmt.Sprintf("%v", c.IsAdult))
	builder.WriteByte(')')
	return builder.String()
}

// Communities is a parsable slice of Community.
type Communities []*Community

func (c Communities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}