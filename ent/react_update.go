// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"newsfeed/ent/comment"
	"newsfeed/ent/post"
	"newsfeed/ent/predicate"
	"newsfeed/ent/react"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReactUpdate is the builder for updating React entities.
type ReactUpdate struct {
	config
	hooks     []Hook
	mutation  *ReactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ReactUpdate builder.
func (ru *ReactUpdate) Where(ps ...predicate.React) *ReactUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReactUpdate) SetUpdatedAt(t time.Time) *ReactUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetReactType sets the "react_type" field.
func (ru *ReactUpdate) SetReactType(s string) *ReactUpdate {
	ru.mutation.SetReactType(s)
	return ru
}

// SetNillableReactType sets the "react_type" field if the given value is not nil.
func (ru *ReactUpdate) SetNillableReactType(s *string) *ReactUpdate {
	if s != nil {
		ru.SetReactType(*s)
	}
	return ru
}

// ClearReactType clears the value of the "react_type" field.
func (ru *ReactUpdate) ClearReactType() *ReactUpdate {
	ru.mutation.ClearReactType()
	return ru
}

// SetPostType sets the "post_type" field.
func (ru *ReactUpdate) SetPostType(s string) *ReactUpdate {
	ru.mutation.SetPostType(s)
	return ru
}

// SetNillablePostType sets the "post_type" field if the given value is not nil.
func (ru *ReactUpdate) SetNillablePostType(s *string) *ReactUpdate {
	if s != nil {
		ru.SetPostType(*s)
	}
	return ru
}

// ClearPostType clears the value of the "post_type" field.
func (ru *ReactUpdate) ClearPostType() *ReactUpdate {
	ru.mutation.ClearPostType()
	return ru
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (ru *ReactUpdate) AddPostIDs(ids ...int) *ReactUpdate {
	ru.mutation.AddPostIDs(ids...)
	return ru
}

// AddPosts adds the "posts" edges to the Post entity.
func (ru *ReactUpdate) AddPosts(p ...*Post) *ReactUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPostIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (ru *ReactUpdate) AddCommentIDs(ids ...int) *ReactUpdate {
	ru.mutation.AddCommentIDs(ids...)
	return ru
}

// AddComments adds the "comments" edges to the Comment entity.
func (ru *ReactUpdate) AddComments(c ...*Comment) *ReactUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddCommentIDs(ids...)
}

// AddReactedUserIDs adds the "reacted_user" edge to the Comment entity by IDs.
func (ru *ReactUpdate) AddReactedUserIDs(ids ...int) *ReactUpdate {
	ru.mutation.AddReactedUserIDs(ids...)
	return ru
}

// AddReactedUser adds the "reacted_user" edges to the Comment entity.
func (ru *ReactUpdate) AddReactedUser(c ...*Comment) *ReactUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddReactedUserIDs(ids...)
}

// Mutation returns the ReactMutation object of the builder.
func (ru *ReactUpdate) Mutation() *ReactMutation {
	return ru.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (ru *ReactUpdate) ClearPosts() *ReactUpdate {
	ru.mutation.ClearPosts()
	return ru
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (ru *ReactUpdate) RemovePostIDs(ids ...int) *ReactUpdate {
	ru.mutation.RemovePostIDs(ids...)
	return ru
}

// RemovePosts removes "posts" edges to Post entities.
func (ru *ReactUpdate) RemovePosts(p ...*Post) *ReactUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePostIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (ru *ReactUpdate) ClearComments() *ReactUpdate {
	ru.mutation.ClearComments()
	return ru
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (ru *ReactUpdate) RemoveCommentIDs(ids ...int) *ReactUpdate {
	ru.mutation.RemoveCommentIDs(ids...)
	return ru
}

// RemoveComments removes "comments" edges to Comment entities.
func (ru *ReactUpdate) RemoveComments(c ...*Comment) *ReactUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveCommentIDs(ids...)
}

// ClearReactedUser clears all "reacted_user" edges to the Comment entity.
func (ru *ReactUpdate) ClearReactedUser() *ReactUpdate {
	ru.mutation.ClearReactedUser()
	return ru
}

// RemoveReactedUserIDs removes the "reacted_user" edge to Comment entities by IDs.
func (ru *ReactUpdate) RemoveReactedUserIDs(ids ...int) *ReactUpdate {
	ru.mutation.RemoveReactedUserIDs(ids...)
	return ru
}

// RemoveReactedUser removes "reacted_user" edges to Comment entities.
func (ru *ReactUpdate) RemoveReactedUser(c ...*Comment) *ReactUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveReactedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReactUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReactUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReactUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReactUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReactUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := react.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *ReactUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReactUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *ReactUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(react.Table, react.Columns, sqlgraph.NewFieldSpec(react.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(react.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.ReactType(); ok {
		_spec.SetField(react.FieldReactType, field.TypeString, value)
	}
	if ru.mutation.ReactTypeCleared() {
		_spec.ClearField(react.FieldReactType, field.TypeString)
	}
	if value, ok := ru.mutation.PostType(); ok {
		_spec.SetField(react.FieldPostType, field.TypeString, value)
	}
	if ru.mutation.PostTypeCleared() {
		_spec.ClearField(react.FieldPostType, field.TypeString)
	}
	if ru.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPostsIDs(); len(nodes) > 0 && !ru.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !ru.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ReactedUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedReactedUserIDs(); len(nodes) > 0 && !ru.mutation.ReactedUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ReactedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{react.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReactUpdateOne is the builder for updating a single React entity.
type ReactUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ReactMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReactUpdateOne) SetUpdatedAt(t time.Time) *ReactUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetReactType sets the "react_type" field.
func (ruo *ReactUpdateOne) SetReactType(s string) *ReactUpdateOne {
	ruo.mutation.SetReactType(s)
	return ruo
}

// SetNillableReactType sets the "react_type" field if the given value is not nil.
func (ruo *ReactUpdateOne) SetNillableReactType(s *string) *ReactUpdateOne {
	if s != nil {
		ruo.SetReactType(*s)
	}
	return ruo
}

// ClearReactType clears the value of the "react_type" field.
func (ruo *ReactUpdateOne) ClearReactType() *ReactUpdateOne {
	ruo.mutation.ClearReactType()
	return ruo
}

// SetPostType sets the "post_type" field.
func (ruo *ReactUpdateOne) SetPostType(s string) *ReactUpdateOne {
	ruo.mutation.SetPostType(s)
	return ruo
}

// SetNillablePostType sets the "post_type" field if the given value is not nil.
func (ruo *ReactUpdateOne) SetNillablePostType(s *string) *ReactUpdateOne {
	if s != nil {
		ruo.SetPostType(*s)
	}
	return ruo
}

// ClearPostType clears the value of the "post_type" field.
func (ruo *ReactUpdateOne) ClearPostType() *ReactUpdateOne {
	ruo.mutation.ClearPostType()
	return ruo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (ruo *ReactUpdateOne) AddPostIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.AddPostIDs(ids...)
	return ruo
}

// AddPosts adds the "posts" edges to the Post entity.
func (ruo *ReactUpdateOne) AddPosts(p ...*Post) *ReactUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPostIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (ruo *ReactUpdateOne) AddCommentIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.AddCommentIDs(ids...)
	return ruo
}

// AddComments adds the "comments" edges to the Comment entity.
func (ruo *ReactUpdateOne) AddComments(c ...*Comment) *ReactUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddCommentIDs(ids...)
}

// AddReactedUserIDs adds the "reacted_user" edge to the Comment entity by IDs.
func (ruo *ReactUpdateOne) AddReactedUserIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.AddReactedUserIDs(ids...)
	return ruo
}

// AddReactedUser adds the "reacted_user" edges to the Comment entity.
func (ruo *ReactUpdateOne) AddReactedUser(c ...*Comment) *ReactUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddReactedUserIDs(ids...)
}

// Mutation returns the ReactMutation object of the builder.
func (ruo *ReactUpdateOne) Mutation() *ReactMutation {
	return ruo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (ruo *ReactUpdateOne) ClearPosts() *ReactUpdateOne {
	ruo.mutation.ClearPosts()
	return ruo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (ruo *ReactUpdateOne) RemovePostIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.RemovePostIDs(ids...)
	return ruo
}

// RemovePosts removes "posts" edges to Post entities.
func (ruo *ReactUpdateOne) RemovePosts(p ...*Post) *ReactUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePostIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (ruo *ReactUpdateOne) ClearComments() *ReactUpdateOne {
	ruo.mutation.ClearComments()
	return ruo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (ruo *ReactUpdateOne) RemoveCommentIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.RemoveCommentIDs(ids...)
	return ruo
}

// RemoveComments removes "comments" edges to Comment entities.
func (ruo *ReactUpdateOne) RemoveComments(c ...*Comment) *ReactUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveCommentIDs(ids...)
}

// ClearReactedUser clears all "reacted_user" edges to the Comment entity.
func (ruo *ReactUpdateOne) ClearReactedUser() *ReactUpdateOne {
	ruo.mutation.ClearReactedUser()
	return ruo
}

// RemoveReactedUserIDs removes the "reacted_user" edge to Comment entities by IDs.
func (ruo *ReactUpdateOne) RemoveReactedUserIDs(ids ...int) *ReactUpdateOne {
	ruo.mutation.RemoveReactedUserIDs(ids...)
	return ruo
}

// RemoveReactedUser removes "reacted_user" edges to Comment entities.
func (ruo *ReactUpdateOne) RemoveReactedUser(c ...*Comment) *ReactUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveReactedUserIDs(ids...)
}

// Where appends a list predicates to the ReactUpdate builder.
func (ruo *ReactUpdateOne) Where(ps ...predicate.React) *ReactUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReactUpdateOne) Select(field string, fields ...string) *ReactUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated React entity.
func (ruo *ReactUpdateOne) Save(ctx context.Context) (*React, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReactUpdateOne) SaveX(ctx context.Context) *React {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReactUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReactUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReactUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := react.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *ReactUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ReactUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *ReactUpdateOne) sqlSave(ctx context.Context) (_node *React, err error) {
	_spec := sqlgraph.NewUpdateSpec(react.Table, react.Columns, sqlgraph.NewFieldSpec(react.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "React.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, react.FieldID)
		for _, f := range fields {
			if !react.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != react.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(react.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.ReactType(); ok {
		_spec.SetField(react.FieldReactType, field.TypeString, value)
	}
	if ruo.mutation.ReactTypeCleared() {
		_spec.ClearField(react.FieldReactType, field.TypeString)
	}
	if value, ok := ruo.mutation.PostType(); ok {
		_spec.SetField(react.FieldPostType, field.TypeString, value)
	}
	if ruo.mutation.PostTypeCleared() {
		_spec.ClearField(react.FieldPostType, field.TypeString)
	}
	if ruo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !ruo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.PostsTable,
			Columns: react.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !ruo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.CommentsTable,
			Columns: react.CommentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ReactedUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedReactedUserIDs(); len(nodes) > 0 && !ruo.mutation.ReactedUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ReactedUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   react.ReactedUserTable,
			Columns: react.ReactedUserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &React{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{react.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
