package schema

import (
	"context"
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/ent/hook"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

// tham khảo thêm tại https://entgo.io/docs/faq/

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Hooks of the Mixin.
func (BaseMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(IDHook(), ent.OpCreate),
	}
}

func IDHook() ent.Hook {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	type IDSetter interface {
		SetID(uint64)
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			is, ok := m.(IDSetter)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation %T", m)
			}
			id, err := sf.NextID()
			if err != nil {
				return nil, err
			}
			is.SetID(id)
			return next.Mutate(ctx, m)
		})
	}
}

type Gender int

const (
	Male Gender = iota
	Female
)

func (p Gender) String() string {
	switch p {
	case Female:
		return "FEMALE"
	default:
		return "MALE"
	}
}

// Values provides list valid values for Enum.
func (Gender) Values() []string {
	return []string{Male.String(), Female.String()}
}

// Value provides the DB a string from int.
func (p Gender) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *Gender) Scan(val interface{}) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}
	switch s {
	case "FEMALE":
		*p = Female
	default:
		*p = Male
	}
	return nil
}