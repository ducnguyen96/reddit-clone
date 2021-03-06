// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/community"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
	"github.com/ducnguyen96/reddit-clone/ent/user"
)

// CommunityCreate is the builder for creating a Community entity.
type CommunityCreate struct {
	config
	mutation *CommunityMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommunityCreate) SetCreatedAt(t time.Time) *CommunityCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableCreatedAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommunityCreate) SetUpdatedAt(t time.Time) *CommunityCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableUpdatedAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CommunityCreate) SetName(s string) *CommunityCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetSlug sets the "slug" field.
func (cc *CommunityCreate) SetSlug(s string) *CommunityCreate {
	cc.mutation.SetSlug(s)
	return cc
}

// SetType sets the "type" field.
func (cc *CommunityCreate) SetType(et enums.CommunityType) *CommunityCreate {
	cc.mutation.SetType(et)
	return cc
}

// SetIsAdult sets the "is_adult" field.
func (cc *CommunityCreate) SetIsAdult(b bool) *CommunityCreate {
	cc.mutation.SetIsAdult(b)
	return cc
}

// SetNillableIsAdult sets the "is_adult" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableIsAdult(b *bool) *CommunityCreate {
	if b != nil {
		cc.SetIsAdult(*b)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommunityCreate) SetID(u uint64) *CommunityCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (cc *CommunityCreate) AddUserIDs(ids ...uint64) *CommunityCreate {
	cc.mutation.AddUserIDs(ids...)
	return cc
}

// AddUsers adds the "users" edges to the User entity.
func (cc *CommunityCreate) AddUsers(u ...*User) *CommunityCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddUserIDs(ids...)
}

// AddAdminIDs adds the "admins" edge to the User entity by IDs.
func (cc *CommunityCreate) AddAdminIDs(ids ...uint64) *CommunityCreate {
	cc.mutation.AddAdminIDs(ids...)
	return cc
}

// AddAdmins adds the "admins" edges to the User entity.
func (cc *CommunityCreate) AddAdmins(u ...*User) *CommunityCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddAdminIDs(ids...)
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cc *CommunityCreate) AddPostIDs(ids ...uint64) *CommunityCreate {
	cc.mutation.AddPostIDs(ids...)
	return cc
}

// AddPosts adds the "posts" edges to the Post entity.
func (cc *CommunityCreate) AddPosts(p ...*Post) *CommunityCreate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPostIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cc *CommunityCreate) Mutation() *CommunityMutation {
	return cc.mutation
}

// Save creates the Community in the database.
func (cc *CommunityCreate) Save(ctx context.Context) (*Community, error) {
	var (
		err  error
		node *Community
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommunityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommunityCreate) SaveX(ctx context.Context) *Community {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommunityCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommunityCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommunityCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if community.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized community.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := community.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if community.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized community.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := community.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.IsAdult(); !ok {
		v := community.DefaultIsAdult
		cc.mutation.SetIsAdult(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommunityCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := community.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "slug"`)}
	}
	if v, ok := cc.mutation.Slug(); ok {
		if err := community.SlugValidator(v); err != nil {
			return &ValidationError{Name: "slug", err: fmt.Errorf(`ent: validator failed for field "slug": %w`, err)}
		}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := community.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsAdult(); !ok {
		return &ValidationError{Name: "is_adult", err: errors.New(`ent: missing required field "is_adult"`)}
	}
	if len(cc.mutation.AdminsIDs()) == 0 {
		return &ValidationError{Name: "admins", err: errors.New("ent: missing required edge \"admins\"")}
	}
	return nil
}

func (cc *CommunityCreate) sqlSave(ctx context.Context) (*Community, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (cc *CommunityCreate) createSpec() (*Community, *sqlgraph.CreateSpec) {
	var (
		_node = &Community{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: community.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: community.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: community.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: community.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: community.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Slug(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: community.FieldSlug,
		})
		_node.Slug = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: community.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.IsAdult(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: community.FieldIsAdult,
		})
		_node.IsAdult = value
	}
	if nodes := cc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.UsersTable,
			Columns: community.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.AdminsTable,
			Columns: community.AdminsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   community.PostsTable,
			Columns: community.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommunityCreateBulk is the builder for creating many Community entities in bulk.
type CommunityCreateBulk struct {
	config
	builders []*CommunityCreate
}

// Save creates the Community entities in the database.
func (ccb *CommunityCreateBulk) Save(ctx context.Context) ([]*Community, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Community, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommunityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommunityCreateBulk) SaveX(ctx context.Context) []*Community {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommunityCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommunityCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
