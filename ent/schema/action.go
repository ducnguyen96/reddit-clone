package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
)

type Action struct {
	ent.Schema
}

func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("target"),
		field.Enum("type").GoType(enums.UserActionType(0)),
		field.Enum("target_type").GoType(enums.UserActionTargetType(0)),
	}
}

func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("actions"),
	}
}

func (Action) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
