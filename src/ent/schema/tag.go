package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}


// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").MinLen(2).MaxLen(50).Annotations(entsql.Annotation{
			Size: 50,
		}).
			Validate(MaxRuneCount(50)),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("tags"),
	}
}


// Mixin of the Tag.
func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}