package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

type Media struct {
	ent.Schema
}


// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").Unique(),
		field.Enum("type").GoType(enums.MediaType(0)),
	}
}

// Mixin of the Media.
func (Media) Mixin() []ent.Mixin {
	return []ent.Mixin{
		// Embed the BaseMixin in the user schema.
		BaseMixin{},
	}
}