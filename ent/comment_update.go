// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/reddit-clone/ent/comment"
	"github.com/ducnguyen96/reddit-clone/ent/post"
	"github.com/ducnguyen96/reddit-clone/ent/predicate"
	"github.com/ducnguyen96/reddit-clone/ent/schema/enums"
	"github.com/ducnguyen96/reddit-clone/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetContentMode sets the "content_mode" field.
func (cu *CommentUpdate) SetContentMode(ecm enums.InputContentMode) *CommentUpdate {
	cu.mutation.SetContentMode(ecm)
	return cu
}

// SetUpVotes sets the "up_votes" field.
func (cu *CommentUpdate) SetUpVotes(i int) *CommentUpdate {
	cu.mutation.ResetUpVotes()
	cu.mutation.SetUpVotes(i)
	return cu
}

// SetNillableUpVotes sets the "up_votes" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUpVotes(i *int) *CommentUpdate {
	if i != nil {
		cu.SetUpVotes(*i)
	}
	return cu
}

// AddUpVotes adds i to the "up_votes" field.
func (cu *CommentUpdate) AddUpVotes(i int) *CommentUpdate {
	cu.mutation.AddUpVotes(i)
	return cu
}

// SetDownVotes sets the "down_votes" field.
func (cu *CommentUpdate) SetDownVotes(i int) *CommentUpdate {
	cu.mutation.ResetDownVotes()
	cu.mutation.SetDownVotes(i)
	return cu
}

// SetNillableDownVotes sets the "down_votes" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDownVotes(i *int) *CommentUpdate {
	if i != nil {
		cu.SetDownVotes(*i)
	}
	return cu
}

// AddDownVotes adds i to the "down_votes" field.
func (cu *CommentUpdate) AddDownVotes(i int) *CommentUpdate {
	cu.mutation.AddDownVotes(i)
	return cu
}

// SetPostID sets the "post_id" field.
func (cu *CommentUpdate) SetPostID(u uint64) *CommentUpdate {
	cu.mutation.ResetPostID()
	cu.mutation.SetPostID(u)
	return cu
}

// AddPostID adds u to the "post_id" field.
func (cu *CommentUpdate) AddPostID(u uint64) *CommentUpdate {
	cu.mutation.AddPostID(u)
	return cu
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cu *CommentUpdate) AddPostIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.AddPostIDs(ids...)
	return cu
}

// AddPosts adds the "posts" edges to the Post entity.
func (cu *CommentUpdate) AddPosts(p ...*Post) *CommentUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (cu *CommentUpdate) AddUserIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUser adds the "user" edges to the User entity.
func (cu *CommentUpdate) AddUser(u ...*User) *CommentUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cu *CommentUpdate) SetParentID(id uint64) *CommentUpdate {
	cu.mutation.SetParentID(id)
	return cu
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cu *CommentUpdate) SetNillableParentID(id *uint64) *CommentUpdate {
	if id != nil {
		cu = cu.SetParentID(*id)
	}
	return cu
}

// SetParent sets the "parent" edge to the Comment entity.
func (cu *CommentUpdate) SetParent(c *Comment) *CommentUpdate {
	return cu.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cu *CommentUpdate) AddChildIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.AddChildIDs(ids...)
	return cu
}

// AddChildren adds the "children" edges to the Comment entity.
func (cu *CommentUpdate) AddChildren(c ...*Comment) *CommentUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChildIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (cu *CommentUpdate) ClearPosts() *CommentUpdate {
	cu.mutation.ClearPosts()
	return cu
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (cu *CommentUpdate) RemovePostIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.RemovePostIDs(ids...)
	return cu
}

// RemovePosts removes "posts" edges to Post entities.
func (cu *CommentUpdate) RemovePosts(p ...*Post) *CommentUpdate {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePostIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (cu *CommentUpdate) ClearUser() *CommentUpdate {
	cu.mutation.ClearUser()
	return cu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (cu *CommentUpdate) RemoveUserIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUser removes "user" edges to User entities.
func (cu *CommentUpdate) RemoveUser(u ...*User) *CommentUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// ClearParent clears the "parent" edge to the Comment entity.
func (cu *CommentUpdate) ClearParent() *CommentUpdate {
	cu.mutation.ClearParent()
	return cu
}

// ClearChildren clears all "children" edges to the Comment entity.
func (cu *CommentUpdate) ClearChildren() *CommentUpdate {
	cu.mutation.ClearChildren()
	return cu
}

// RemoveChildIDs removes the "children" edge to Comment entities by IDs.
func (cu *CommentUpdate) RemoveChildIDs(ids ...uint64) *CommentUpdate {
	cu.mutation.RemoveChildIDs(ids...)
	return cu
}

// RemoveChildren removes "children" edges to Comment entities.
func (cu *CommentUpdate) RemoveChildren(c ...*Comment) *CommentUpdate {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.ContentMode(); ok {
		if err := comment.ContentModeValidator(v); err != nil {
			return &ValidationError{Name: "content_mode", err: fmt.Errorf("ent: validator failed for field \"content_mode\": %w", err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cu.mutation.ContentMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: comment.FieldContentMode,
		})
	}
	if value, ok := cu.mutation.UpVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUpVotes,
		})
	}
	if value, ok := cu.mutation.AddedUpVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUpVotes,
		})
	}
	if value, ok := cu.mutation.DownVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldDownVotes,
		})
	}
	if value, ok := cu.mutation.AddedDownVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldDownVotes,
		})
	}
	if value, ok := cu.mutation.PostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldPostID,
		})
	}
	if value, ok := cu.mutation.AddedPostID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldPostID,
		})
	}
	if cu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPostsIDs(); len(nodes) > 0 && !cu.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedUserIDs(); len(nodes) > 0 && !cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetContentMode sets the "content_mode" field.
func (cuo *CommentUpdateOne) SetContentMode(ecm enums.InputContentMode) *CommentUpdateOne {
	cuo.mutation.SetContentMode(ecm)
	return cuo
}

// SetUpVotes sets the "up_votes" field.
func (cuo *CommentUpdateOne) SetUpVotes(i int) *CommentUpdateOne {
	cuo.mutation.ResetUpVotes()
	cuo.mutation.SetUpVotes(i)
	return cuo
}

// SetNillableUpVotes sets the "up_votes" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUpVotes(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetUpVotes(*i)
	}
	return cuo
}

// AddUpVotes adds i to the "up_votes" field.
func (cuo *CommentUpdateOne) AddUpVotes(i int) *CommentUpdateOne {
	cuo.mutation.AddUpVotes(i)
	return cuo
}

// SetDownVotes sets the "down_votes" field.
func (cuo *CommentUpdateOne) SetDownVotes(i int) *CommentUpdateOne {
	cuo.mutation.ResetDownVotes()
	cuo.mutation.SetDownVotes(i)
	return cuo
}

// SetNillableDownVotes sets the "down_votes" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDownVotes(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetDownVotes(*i)
	}
	return cuo
}

// AddDownVotes adds i to the "down_votes" field.
func (cuo *CommentUpdateOne) AddDownVotes(i int) *CommentUpdateOne {
	cuo.mutation.AddDownVotes(i)
	return cuo
}

// SetPostID sets the "post_id" field.
func (cuo *CommentUpdateOne) SetPostID(u uint64) *CommentUpdateOne {
	cuo.mutation.ResetPostID()
	cuo.mutation.SetPostID(u)
	return cuo
}

// AddPostID adds u to the "post_id" field.
func (cuo *CommentUpdateOne) AddPostID(u uint64) *CommentUpdateOne {
	cuo.mutation.AddPostID(u)
	return cuo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (cuo *CommentUpdateOne) AddPostIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.AddPostIDs(ids...)
	return cuo
}

// AddPosts adds the "posts" edges to the Post entity.
func (cuo *CommentUpdateOne) AddPosts(p ...*Post) *CommentUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (cuo *CommentUpdateOne) AddUserIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUser adds the "user" edges to the User entity.
func (cuo *CommentUpdateOne) AddUser(u ...*User) *CommentUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// SetParentID sets the "parent" edge to the Comment entity by ID.
func (cuo *CommentUpdateOne) SetParentID(id uint64) *CommentUpdateOne {
	cuo.mutation.SetParentID(id)
	return cuo
}

// SetNillableParentID sets the "parent" edge to the Comment entity by ID if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableParentID(id *uint64) *CommentUpdateOne {
	if id != nil {
		cuo = cuo.SetParentID(*id)
	}
	return cuo
}

// SetParent sets the "parent" edge to the Comment entity.
func (cuo *CommentUpdateOne) SetParent(c *Comment) *CommentUpdateOne {
	return cuo.SetParentID(c.ID)
}

// AddChildIDs adds the "children" edge to the Comment entity by IDs.
func (cuo *CommentUpdateOne) AddChildIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.AddChildIDs(ids...)
	return cuo
}

// AddChildren adds the "children" edges to the Comment entity.
func (cuo *CommentUpdateOne) AddChildren(c ...*Comment) *CommentUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChildIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (cuo *CommentUpdateOne) ClearPosts() *CommentUpdateOne {
	cuo.mutation.ClearPosts()
	return cuo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (cuo *CommentUpdateOne) RemovePostIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.RemovePostIDs(ids...)
	return cuo
}

// RemovePosts removes "posts" edges to Post entities.
func (cuo *CommentUpdateOne) RemovePosts(p ...*Post) *CommentUpdateOne {
	ids := make([]uint64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePostIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (cuo *CommentUpdateOne) ClearUser() *CommentUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (cuo *CommentUpdateOne) RemoveUserIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUser removes "user" edges to User entities.
func (cuo *CommentUpdateOne) RemoveUser(u ...*User) *CommentUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// ClearParent clears the "parent" edge to the Comment entity.
func (cuo *CommentUpdateOne) ClearParent() *CommentUpdateOne {
	cuo.mutation.ClearParent()
	return cuo
}

// ClearChildren clears all "children" edges to the Comment entity.
func (cuo *CommentUpdateOne) ClearChildren() *CommentUpdateOne {
	cuo.mutation.ClearChildren()
	return cuo
}

// RemoveChildIDs removes the "children" edge to Comment entities by IDs.
func (cuo *CommentUpdateOne) RemoveChildIDs(ids ...uint64) *CommentUpdateOne {
	cuo.mutation.RemoveChildIDs(ids...)
	return cuo
}

// RemoveChildren removes "children" edges to Comment entities.
func (cuo *CommentUpdateOne) RemoveChildren(c ...*Comment) *CommentUpdateOne {
	ids := make([]uint64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.ContentMode(); ok {
		if err := comment.ContentModeValidator(v); err != nil {
			return &ValidationError{Name: "content_mode", err: fmt.Errorf("ent: validator failed for field \"content_mode\": %w", err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cuo.mutation.ContentMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: comment.FieldContentMode,
		})
	}
	if value, ok := cuo.mutation.UpVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUpVotes,
		})
	}
	if value, ok := cuo.mutation.AddedUpVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUpVotes,
		})
	}
	if value, ok := cuo.mutation.DownVotes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldDownVotes,
		})
	}
	if value, ok := cuo.mutation.AddedDownVotes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldDownVotes,
		})
	}
	if value, ok := cuo.mutation.PostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldPostID,
		})
	}
	if value, ok := cuo.mutation.AddedPostID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: comment.FieldPostID,
		})
	}
	if cuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !cuo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.PostsTable,
			Columns: comment.PostsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   comment.UserTable,
			Columns: comment.UserPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.ParentTable,
			Columns: []string{comment.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !cuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.ChildrenTable,
			Columns: []string{comment.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
