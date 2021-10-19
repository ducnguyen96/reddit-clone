package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(2).MaxLen(300).Annotations(entsql.Annotation{
			Size: 300,
		}).
			Validate(MaxRuneCount(300)),
		field.String("slug").MinLen(2).MaxLen(300).Annotations(entsql.Annotation{
			Size: 300,
		}).
			Validate(MaxRuneCount(300)).Unique(),
		field.Enum("type").GoType(enums.CommunityType(0)),
		field.Bool("is_adult").Default(false),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("admins", User.Type).Required(),
		edge.To("posts", Post.Type),
	}
}

// Mixin of the Community.
func (Community) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}
