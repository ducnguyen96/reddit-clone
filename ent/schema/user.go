package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("email").MinLen(10).MaxLen(100).Annotations(entsql.Annotation{
			Size: 150,
		}).
			Validate(MaxRuneCount(100)).Nillable().Optional(),
		field.String("avatar_url").Nillable().Optional(),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("communities", Community.Type).Ref("users"),
		edge.From("my_communities", Community.Type).Ref("admins"),
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("actions", Action.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}
