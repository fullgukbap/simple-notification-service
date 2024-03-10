// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"notification-service/ent/predicate"
	"notification-service/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipUpdate is the builder for updating Friendship entities.
type FriendshipUpdate struct {
	config
	hooks    []Hook
	mutation *FriendshipMutation
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fu *FriendshipUpdate) Where(ps ...predicate.Friendship) *FriendshipUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetDeleteTime sets the "delete_time" field.
func (fu *FriendshipUpdate) SetDeleteTime(t time.Time) *FriendshipUpdate {
	fu.mutation.SetDeleteTime(t)
	return fu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableDeleteTime(t *time.Time) *FriendshipUpdate {
	if t != nil {
		fu.SetDeleteTime(*t)
	}
	return fu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (fu *FriendshipUpdate) ClearDeleteTime() *FriendshipUpdate {
	fu.mutation.ClearDeleteTime()
	return fu
}

// SetSenderIDID sets the "senderID" edge to the User entity by ID.
func (fu *FriendshipUpdate) SetSenderIDID(id int) *FriendshipUpdate {
	fu.mutation.SetSenderIDID(id)
	return fu
}

// SetNillableSenderIDID sets the "senderID" edge to the User entity by ID if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableSenderIDID(id *int) *FriendshipUpdate {
	if id != nil {
		fu = fu.SetSenderIDID(*id)
	}
	return fu
}

// SetSenderID sets the "senderID" edge to the User entity.
func (fu *FriendshipUpdate) SetSenderID(u *User) *FriendshipUpdate {
	return fu.SetSenderIDID(u.ID)
}

// SetReceiverIDID sets the "receiverID" edge to the User entity by ID.
func (fu *FriendshipUpdate) SetReceiverIDID(id int) *FriendshipUpdate {
	fu.mutation.SetReceiverIDID(id)
	return fu
}

// SetNillableReceiverIDID sets the "receiverID" edge to the User entity by ID if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableReceiverIDID(id *int) *FriendshipUpdate {
	if id != nil {
		fu = fu.SetReceiverIDID(*id)
	}
	return fu
}

// SetReceiverID sets the "receiverID" edge to the User entity.
func (fu *FriendshipUpdate) SetReceiverID(u *User) *FriendshipUpdate {
	return fu.SetReceiverIDID(u.ID)
}

// SetFriendshipStatusIDID sets the "friendshipStatusID" edge to the FriendshipStatus entity by ID.
func (fu *FriendshipUpdate) SetFriendshipStatusIDID(id int) *FriendshipUpdate {
	fu.mutation.SetFriendshipStatusIDID(id)
	return fu
}

// SetNillableFriendshipStatusIDID sets the "friendshipStatusID" edge to the FriendshipStatus entity by ID if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableFriendshipStatusIDID(id *int) *FriendshipUpdate {
	if id != nil {
		fu = fu.SetFriendshipStatusIDID(*id)
	}
	return fu
}

// SetFriendshipStatusID sets the "friendshipStatusID" edge to the FriendshipStatus entity.
func (fu *FriendshipUpdate) SetFriendshipStatusID(f *FriendshipStatus) *FriendshipUpdate {
	return fu.SetFriendshipStatusIDID(f.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fu *FriendshipUpdate) Mutation() *FriendshipMutation {
	return fu.mutation
}

// ClearSenderID clears the "senderID" edge to the User entity.
func (fu *FriendshipUpdate) ClearSenderID() *FriendshipUpdate {
	fu.mutation.ClearSenderID()
	return fu
}

// ClearReceiverID clears the "receiverID" edge to the User entity.
func (fu *FriendshipUpdate) ClearReceiverID() *FriendshipUpdate {
	fu.mutation.ClearReceiverID()
	return fu
}

// ClearFriendshipStatusID clears the "friendshipStatusID" edge to the FriendshipStatus entity.
func (fu *FriendshipUpdate) ClearFriendshipStatusID() *FriendshipUpdate {
	fu.mutation.ClearFriendshipStatusID()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendshipUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendshipUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendshipUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FriendshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.DeleteTime(); ok {
		_spec.SetField(friendship.FieldDeleteTime, field.TypeTime, value)
	}
	if fu.mutation.DeleteTimeCleared() {
		_spec.ClearField(friendship.FieldDeleteTime, field.TypeTime)
	}
	if fu.mutation.SenderIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.SenderIDTable,
			Columns: []string{friendship.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.SenderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.SenderIDTable,
			Columns: []string{friendship.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.ReceiverIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.ReceiverIDTable,
			Columns: []string{friendship.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.ReceiverIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.ReceiverIDTable,
			Columns: []string{friendship.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FriendshipStatusIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.FriendshipStatusIDTable,
			Columns: []string{friendship.FriendshipStatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FriendshipStatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.FriendshipStatusIDTable,
			Columns: []string{friendship.FriendshipStatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendshipUpdateOne is the builder for updating a single Friendship entity.
type FriendshipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendshipMutation
}

// SetDeleteTime sets the "delete_time" field.
func (fuo *FriendshipUpdateOne) SetDeleteTime(t time.Time) *FriendshipUpdateOne {
	fuo.mutation.SetDeleteTime(t)
	return fuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableDeleteTime(t *time.Time) *FriendshipUpdateOne {
	if t != nil {
		fuo.SetDeleteTime(*t)
	}
	return fuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (fuo *FriendshipUpdateOne) ClearDeleteTime() *FriendshipUpdateOne {
	fuo.mutation.ClearDeleteTime()
	return fuo
}

// SetSenderIDID sets the "senderID" edge to the User entity by ID.
func (fuo *FriendshipUpdateOne) SetSenderIDID(id int) *FriendshipUpdateOne {
	fuo.mutation.SetSenderIDID(id)
	return fuo
}

// SetNillableSenderIDID sets the "senderID" edge to the User entity by ID if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableSenderIDID(id *int) *FriendshipUpdateOne {
	if id != nil {
		fuo = fuo.SetSenderIDID(*id)
	}
	return fuo
}

// SetSenderID sets the "senderID" edge to the User entity.
func (fuo *FriendshipUpdateOne) SetSenderID(u *User) *FriendshipUpdateOne {
	return fuo.SetSenderIDID(u.ID)
}

// SetReceiverIDID sets the "receiverID" edge to the User entity by ID.
func (fuo *FriendshipUpdateOne) SetReceiverIDID(id int) *FriendshipUpdateOne {
	fuo.mutation.SetReceiverIDID(id)
	return fuo
}

// SetNillableReceiverIDID sets the "receiverID" edge to the User entity by ID if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableReceiverIDID(id *int) *FriendshipUpdateOne {
	if id != nil {
		fuo = fuo.SetReceiverIDID(*id)
	}
	return fuo
}

// SetReceiverID sets the "receiverID" edge to the User entity.
func (fuo *FriendshipUpdateOne) SetReceiverID(u *User) *FriendshipUpdateOne {
	return fuo.SetReceiverIDID(u.ID)
}

// SetFriendshipStatusIDID sets the "friendshipStatusID" edge to the FriendshipStatus entity by ID.
func (fuo *FriendshipUpdateOne) SetFriendshipStatusIDID(id int) *FriendshipUpdateOne {
	fuo.mutation.SetFriendshipStatusIDID(id)
	return fuo
}

// SetNillableFriendshipStatusIDID sets the "friendshipStatusID" edge to the FriendshipStatus entity by ID if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableFriendshipStatusIDID(id *int) *FriendshipUpdateOne {
	if id != nil {
		fuo = fuo.SetFriendshipStatusIDID(*id)
	}
	return fuo
}

// SetFriendshipStatusID sets the "friendshipStatusID" edge to the FriendshipStatus entity.
func (fuo *FriendshipUpdateOne) SetFriendshipStatusID(f *FriendshipStatus) *FriendshipUpdateOne {
	return fuo.SetFriendshipStatusIDID(f.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fuo *FriendshipUpdateOne) Mutation() *FriendshipMutation {
	return fuo.mutation
}

// ClearSenderID clears the "senderID" edge to the User entity.
func (fuo *FriendshipUpdateOne) ClearSenderID() *FriendshipUpdateOne {
	fuo.mutation.ClearSenderID()
	return fuo
}

// ClearReceiverID clears the "receiverID" edge to the User entity.
func (fuo *FriendshipUpdateOne) ClearReceiverID() *FriendshipUpdateOne {
	fuo.mutation.ClearReceiverID()
	return fuo
}

// ClearFriendshipStatusID clears the "friendshipStatusID" edge to the FriendshipStatus entity.
func (fuo *FriendshipUpdateOne) ClearFriendshipStatusID() *FriendshipUpdateOne {
	fuo.mutation.ClearFriendshipStatusID()
	return fuo
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fuo *FriendshipUpdateOne) Where(ps ...predicate.Friendship) *FriendshipUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendshipUpdateOne) Select(field string, fields ...string) *FriendshipUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friendship entity.
func (fuo *FriendshipUpdateOne) Save(ctx context.Context) (*Friendship, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) SaveX(ctx context.Context) *Friendship {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendshipUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FriendshipUpdateOne) sqlSave(ctx context.Context) (_node *Friendship, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friendship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for _, f := range fields {
			if !friendship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.DeleteTime(); ok {
		_spec.SetField(friendship.FieldDeleteTime, field.TypeTime, value)
	}
	if fuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(friendship.FieldDeleteTime, field.TypeTime)
	}
	if fuo.mutation.SenderIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.SenderIDTable,
			Columns: []string{friendship.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.SenderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.SenderIDTable,
			Columns: []string{friendship.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.ReceiverIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.ReceiverIDTable,
			Columns: []string{friendship.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.ReceiverIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.ReceiverIDTable,
			Columns: []string{friendship.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FriendshipStatusIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.FriendshipStatusIDTable,
			Columns: []string{friendship.FriendshipStatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FriendshipStatusIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   friendship.FriendshipStatusIDTable,
			Columns: []string{friendship.FriendshipStatusIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Friendship{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}