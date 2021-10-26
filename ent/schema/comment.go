package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

type Comment struct {
	ent.Schema
}


// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.Enum("content_mode").GoType(enums.InputContentMode(0)),
		field.Int("up_votes").Default(0),
		field.Int("down_votes").Default(0),
		field.Uint64("post_id"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("comments").Required(),
		edge.From("user", User.Type).Ref("comments").Required(),
		edge.To("children", Comment.Type).
			From("parent").
			Unique(),
	}
}

// Mixin of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}