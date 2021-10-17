package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

type Post struct {
	ent.Schema
}


// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MinLen(2).MaxLen(300).Annotations(entsql.Annotation{
			Size: 300,
		}).
			Validate(MaxRuneCount(300)).Unique(),
		field.String("slug").MinLen(2).MaxLen(400).Annotations(entsql.Annotation{
			Size: 400,
		}).Validate(MaxRuneCount(400)).Unique(),
		field.String("content"),
		field.Enum("type").GoType(enums.PostType(0)),
		field.Enum("content_mode").GoType(enums.InputContentMode(0)),
		field.Int("up_votes").Default(0),
		field.Int("down_votes").Default(0),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("posts").Required(),
		edge.From("community", Community.Type).Ref("posts").Required(),
		edge.To("tags", Tag.Type),
		edge.To("comments", Comment.Type),
	}
}

// Mixin of the Post.
func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}