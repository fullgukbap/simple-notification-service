// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"notification-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipStatusUpdate is the builder for updating FriendshipStatus entities.
type FriendshipStatusUpdate struct {
	config
	hooks    []Hook
	mutation *FriendshipStatusMutation
}

// Where appends a list predicates to the FriendshipStatusUpdate builder.
func (fsu *FriendshipStatusUpdate) Where(ps ...predicate.FriendshipStatus) *FriendshipStatusUpdate {
	fsu.mutation.Where(ps...)
	return fsu
}

// SetUpdatedAt sets the "updated_at" field.
func (fsu *FriendshipStatusUpdate) SetUpdatedAt(t time.Time) *FriendshipStatusUpdate {
	fsu.mutation.SetUpdatedAt(t)
	return fsu
}

// SetDeletedAt sets the "deleted_at" field.
func (fsu *FriendshipStatusUpdate) SetDeletedAt(t time.Time) *FriendshipStatusUpdate {
	fsu.mutation.SetDeletedAt(t)
	return fsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fsu *FriendshipStatusUpdate) SetNillableDeletedAt(t *time.Time) *FriendshipStatusUpdate {
	if t != nil {
		fsu.SetDeletedAt(*t)
	}
	return fsu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fsu *FriendshipStatusUpdate) ClearDeletedAt() *FriendshipStatusUpdate {
	fsu.mutation.ClearDeletedAt()
	return fsu
}

// SetDescription sets the "description" field.
func (fsu *FriendshipStatusUpdate) SetDescription(s string) *FriendshipStatusUpdate {
	fsu.mutation.SetDescription(s)
	return fsu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fsu *FriendshipStatusUpdate) SetNillableDescription(s *string) *FriendshipStatusUpdate {
	if s != nil {
		fsu.SetDescription(*s)
	}
	return fsu
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (fsu *FriendshipStatusUpdate) AddFriendshipIDs(ids ...int) *FriendshipStatusUpdate {
	fsu.mutation.AddFriendshipIDs(ids...)
	return fsu
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (fsu *FriendshipStatusUpdate) AddFriendships(f ...*Friendship) *FriendshipStatusUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.AddFriendshipIDs(ids...)
}

// Mutation returns the FriendshipStatusMutation object of the builder.
func (fsu *FriendshipStatusUpdate) Mutation() *FriendshipStatusMutation {
	return fsu.mutation
}

// ClearFriendships clears all "friendships" edges to the Friendship entity.
func (fsu *FriendshipStatusUpdate) ClearFriendships() *FriendshipStatusUpdate {
	fsu.mutation.ClearFriendships()
	return fsu
}

// RemoveFriendshipIDs removes the "friendships" edge to Friendship entities by IDs.
func (fsu *FriendshipStatusUpdate) RemoveFriendshipIDs(ids ...int) *FriendshipStatusUpdate {
	fsu.mutation.RemoveFriendshipIDs(ids...)
	return fsu
}

// RemoveFriendships removes "friendships" edges to Friendship entities.
func (fsu *FriendshipStatusUpdate) RemoveFriendships(f ...*Friendship) *FriendshipStatusUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.RemoveFriendshipIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fsu *FriendshipStatusUpdate) Save(ctx context.Context) (int, error) {
	if err := fsu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, fsu.sqlSave, fsu.mutation, fsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsu *FriendshipStatusUpdate) SaveX(ctx context.Context) int {
	affected, err := fsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fsu *FriendshipStatusUpdate) Exec(ctx context.Context) error {
	_, err := fsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsu *FriendshipStatusUpdate) ExecX(ctx context.Context) {
	if err := fsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fsu *FriendshipStatusUpdate) defaults() error {
	if _, ok := fsu.mutation.UpdatedAt(); !ok {
		if friendshipstatus.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized friendshipstatus.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := friendshipstatus.UpdateDefaultUpdatedAt()
		fsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fsu *FriendshipStatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendshipstatus.Table, friendshipstatus.Columns, sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt))
	if ps := fsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fsu.mutation.UpdatedAt(); ok {
		_spec.SetField(friendshipstatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fsu.mutation.DeletedAt(); ok {
		_spec.SetField(friendshipstatus.FieldDeletedAt, field.TypeTime, value)
	}
	if fsu.mutation.DeletedAtCleared() {
		_spec.ClearField(friendshipstatus.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fsu.mutation.Description(); ok {
		_spec.SetField(friendshipstatus.FieldDescription, field.TypeString, value)
	}
	if fsu.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.RemovedFriendshipsIDs(); len(nodes) > 0 && !fsu.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendshipstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fsu.mutation.done = true
	return n, nil
}

// FriendshipStatusUpdateOne is the builder for updating a single FriendshipStatus entity.
type FriendshipStatusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendshipStatusMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fsuo *FriendshipStatusUpdateOne) SetUpdatedAt(t time.Time) *FriendshipStatusUpdateOne {
	fsuo.mutation.SetUpdatedAt(t)
	return fsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fsuo *FriendshipStatusUpdateOne) SetDeletedAt(t time.Time) *FriendshipStatusUpdateOne {
	fsuo.mutation.SetDeletedAt(t)
	return fsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fsuo *FriendshipStatusUpdateOne) SetNillableDeletedAt(t *time.Time) *FriendshipStatusUpdateOne {
	if t != nil {
		fsuo.SetDeletedAt(*t)
	}
	return fsuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fsuo *FriendshipStatusUpdateOne) ClearDeletedAt() *FriendshipStatusUpdateOne {
	fsuo.mutation.ClearDeletedAt()
	return fsuo
}

// SetDescription sets the "description" field.
func (fsuo *FriendshipStatusUpdateOne) SetDescription(s string) *FriendshipStatusUpdateOne {
	fsuo.mutation.SetDescription(s)
	return fsuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fsuo *FriendshipStatusUpdateOne) SetNillableDescription(s *string) *FriendshipStatusUpdateOne {
	if s != nil {
		fsuo.SetDescription(*s)
	}
	return fsuo
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (fsuo *FriendshipStatusUpdateOne) AddFriendshipIDs(ids ...int) *FriendshipStatusUpdateOne {
	fsuo.mutation.AddFriendshipIDs(ids...)
	return fsuo
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (fsuo *FriendshipStatusUpdateOne) AddFriendships(f ...*Friendship) *FriendshipStatusUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.AddFriendshipIDs(ids...)
}

// Mutation returns the FriendshipStatusMutation object of the builder.
func (fsuo *FriendshipStatusUpdateOne) Mutation() *FriendshipStatusMutation {
	return fsuo.mutation
}

// ClearFriendships clears all "friendships" edges to the Friendship entity.
func (fsuo *FriendshipStatusUpdateOne) ClearFriendships() *FriendshipStatusUpdateOne {
	fsuo.mutation.ClearFriendships()
	return fsuo
}

// RemoveFriendshipIDs removes the "friendships" edge to Friendship entities by IDs.
func (fsuo *FriendshipStatusUpdateOne) RemoveFriendshipIDs(ids ...int) *FriendshipStatusUpdateOne {
	fsuo.mutation.RemoveFriendshipIDs(ids...)
	return fsuo
}

// RemoveFriendships removes "friendships" edges to Friendship entities.
func (fsuo *FriendshipStatusUpdateOne) RemoveFriendships(f ...*Friendship) *FriendshipStatusUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.RemoveFriendshipIDs(ids...)
}

// Where appends a list predicates to the FriendshipStatusUpdate builder.
func (fsuo *FriendshipStatusUpdateOne) Where(ps ...predicate.FriendshipStatus) *FriendshipStatusUpdateOne {
	fsuo.mutation.Where(ps...)
	return fsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fsuo *FriendshipStatusUpdateOne) Select(field string, fields ...string) *FriendshipStatusUpdateOne {
	fsuo.fields = append([]string{field}, fields...)
	return fsuo
}

// Save executes the query and returns the updated FriendshipStatus entity.
func (fsuo *FriendshipStatusUpdateOne) Save(ctx context.Context) (*FriendshipStatus, error) {
	if err := fsuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fsuo.sqlSave, fsuo.mutation, fsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsuo *FriendshipStatusUpdateOne) SaveX(ctx context.Context) *FriendshipStatus {
	node, err := fsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fsuo *FriendshipStatusUpdateOne) Exec(ctx context.Context) error {
	_, err := fsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsuo *FriendshipStatusUpdateOne) ExecX(ctx context.Context) {
	if err := fsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fsuo *FriendshipStatusUpdateOne) defaults() error {
	if _, ok := fsuo.mutation.UpdatedAt(); !ok {
		if friendshipstatus.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized friendshipstatus.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := friendshipstatus.UpdateDefaultUpdatedAt()
		fsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fsuo *FriendshipStatusUpdateOne) sqlSave(ctx context.Context) (_node *FriendshipStatus, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendshipstatus.Table, friendshipstatus.Columns, sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt))
	id, ok := fsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FriendshipStatus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendshipstatus.FieldID)
		for _, f := range fields {
			if !friendshipstatus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendshipstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(friendshipstatus.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fsuo.mutation.DeletedAt(); ok {
		_spec.SetField(friendshipstatus.FieldDeletedAt, field.TypeTime, value)
	}
	if fsuo.mutation.DeletedAtCleared() {
		_spec.ClearField(friendshipstatus.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fsuo.mutation.Description(); ok {
		_spec.SetField(friendshipstatus.FieldDescription, field.TypeString, value)
	}
	if fsuo.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.RemovedFriendshipsIDs(); len(nodes) > 0 && !fsuo.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   friendshipstatus.FriendshipsTable,
			Columns: []string{friendshipstatus.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FriendshipStatus{config: fsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendshipstatus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fsuo.mutation.done = true
	return _node, nil
}
