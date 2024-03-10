// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification-service/ent/predicate"
	"notification-service/ent/toktok"
	"notification-service/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TokTokUpdate is the builder for updating TokTok entities.
type TokTokUpdate struct {
	config
	hooks    []Hook
	mutation *TokTokMutation
}

// Where appends a list predicates to the TokTokUpdate builder.
func (ttu *TokTokUpdate) Where(ps ...predicate.TokTok) *TokTokUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetDeleteTime sets the "delete_time" field.
func (ttu *TokTokUpdate) SetDeleteTime(t time.Time) *TokTokUpdate {
	ttu.mutation.SetDeleteTime(t)
	return ttu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ttu *TokTokUpdate) SetNillableDeleteTime(t *time.Time) *TokTokUpdate {
	if t != nil {
		ttu.SetDeleteTime(*t)
	}
	return ttu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ttu *TokTokUpdate) ClearDeleteTime() *TokTokUpdate {
	ttu.mutation.ClearDeleteTime()
	return ttu
}

// SetReceiverIDID sets the "receiverID" edge to the User entity by ID.
func (ttu *TokTokUpdate) SetReceiverIDID(id int) *TokTokUpdate {
	ttu.mutation.SetReceiverIDID(id)
	return ttu
}

// SetNillableReceiverIDID sets the "receiverID" edge to the User entity by ID if the given value is not nil.
func (ttu *TokTokUpdate) SetNillableReceiverIDID(id *int) *TokTokUpdate {
	if id != nil {
		ttu = ttu.SetReceiverIDID(*id)
	}
	return ttu
}

// SetReceiverID sets the "receiverID" edge to the User entity.
func (ttu *TokTokUpdate) SetReceiverID(u *User) *TokTokUpdate {
	return ttu.SetReceiverIDID(u.ID)
}

// SetSenderIDID sets the "senderID" edge to the User entity by ID.
func (ttu *TokTokUpdate) SetSenderIDID(id int) *TokTokUpdate {
	ttu.mutation.SetSenderIDID(id)
	return ttu
}

// SetNillableSenderIDID sets the "senderID" edge to the User entity by ID if the given value is not nil.
func (ttu *TokTokUpdate) SetNillableSenderIDID(id *int) *TokTokUpdate {
	if id != nil {
		ttu = ttu.SetSenderIDID(*id)
	}
	return ttu
}

// SetSenderID sets the "senderID" edge to the User entity.
func (ttu *TokTokUpdate) SetSenderID(u *User) *TokTokUpdate {
	return ttu.SetSenderIDID(u.ID)
}

// Mutation returns the TokTokMutation object of the builder.
func (ttu *TokTokUpdate) Mutation() *TokTokMutation {
	return ttu.mutation
}

// ClearReceiverID clears the "receiverID" edge to the User entity.
func (ttu *TokTokUpdate) ClearReceiverID() *TokTokUpdate {
	ttu.mutation.ClearReceiverID()
	return ttu
}

// ClearSenderID clears the "senderID" edge to the User entity.
func (ttu *TokTokUpdate) ClearSenderID() *TokTokUpdate {
	ttu.mutation.ClearSenderID()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TokTokUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ttu.sqlSave, ttu.mutation, ttu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TokTokUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TokTokUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TokTokUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ttu *TokTokUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(toktok.Table, toktok.Columns, sqlgraph.NewFieldSpec(toktok.FieldID, field.TypeInt))
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.DeleteTime(); ok {
		_spec.SetField(toktok.FieldDeleteTime, field.TypeTime, value)
	}
	if ttu.mutation.DeleteTimeCleared() {
		_spec.ClearField(toktok.FieldDeleteTime, field.TypeTime)
	}
	if ttu.mutation.ReceiverIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.ReceiverIDTable,
			Columns: []string{toktok.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.ReceiverIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.ReceiverIDTable,
			Columns: []string{toktok.ReceiverIDColumn},
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
	if ttu.mutation.SenderIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.SenderIDTable,
			Columns: []string{toktok.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.SenderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.SenderIDTable,
			Columns: []string{toktok.SenderIDColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{toktok.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ttu.mutation.done = true
	return n, nil
}

// TokTokUpdateOne is the builder for updating a single TokTok entity.
type TokTokUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokTokMutation
}

// SetDeleteTime sets the "delete_time" field.
func (ttuo *TokTokUpdateOne) SetDeleteTime(t time.Time) *TokTokUpdateOne {
	ttuo.mutation.SetDeleteTime(t)
	return ttuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ttuo *TokTokUpdateOne) SetNillableDeleteTime(t *time.Time) *TokTokUpdateOne {
	if t != nil {
		ttuo.SetDeleteTime(*t)
	}
	return ttuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ttuo *TokTokUpdateOne) ClearDeleteTime() *TokTokUpdateOne {
	ttuo.mutation.ClearDeleteTime()
	return ttuo
}

// SetReceiverIDID sets the "receiverID" edge to the User entity by ID.
func (ttuo *TokTokUpdateOne) SetReceiverIDID(id int) *TokTokUpdateOne {
	ttuo.mutation.SetReceiverIDID(id)
	return ttuo
}

// SetNillableReceiverIDID sets the "receiverID" edge to the User entity by ID if the given value is not nil.
func (ttuo *TokTokUpdateOne) SetNillableReceiverIDID(id *int) *TokTokUpdateOne {
	if id != nil {
		ttuo = ttuo.SetReceiverIDID(*id)
	}
	return ttuo
}

// SetReceiverID sets the "receiverID" edge to the User entity.
func (ttuo *TokTokUpdateOne) SetReceiverID(u *User) *TokTokUpdateOne {
	return ttuo.SetReceiverIDID(u.ID)
}

// SetSenderIDID sets the "senderID" edge to the User entity by ID.
func (ttuo *TokTokUpdateOne) SetSenderIDID(id int) *TokTokUpdateOne {
	ttuo.mutation.SetSenderIDID(id)
	return ttuo
}

// SetNillableSenderIDID sets the "senderID" edge to the User entity by ID if the given value is not nil.
func (ttuo *TokTokUpdateOne) SetNillableSenderIDID(id *int) *TokTokUpdateOne {
	if id != nil {
		ttuo = ttuo.SetSenderIDID(*id)
	}
	return ttuo
}

// SetSenderID sets the "senderID" edge to the User entity.
func (ttuo *TokTokUpdateOne) SetSenderID(u *User) *TokTokUpdateOne {
	return ttuo.SetSenderIDID(u.ID)
}

// Mutation returns the TokTokMutation object of the builder.
func (ttuo *TokTokUpdateOne) Mutation() *TokTokMutation {
	return ttuo.mutation
}

// ClearReceiverID clears the "receiverID" edge to the User entity.
func (ttuo *TokTokUpdateOne) ClearReceiverID() *TokTokUpdateOne {
	ttuo.mutation.ClearReceiverID()
	return ttuo
}

// ClearSenderID clears the "senderID" edge to the User entity.
func (ttuo *TokTokUpdateOne) ClearSenderID() *TokTokUpdateOne {
	ttuo.mutation.ClearSenderID()
	return ttuo
}

// Where appends a list predicates to the TokTokUpdate builder.
func (ttuo *TokTokUpdateOne) Where(ps ...predicate.TokTok) *TokTokUpdateOne {
	ttuo.mutation.Where(ps...)
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TokTokUpdateOne) Select(field string, fields ...string) *TokTokUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TokTok entity.
func (ttuo *TokTokUpdateOne) Save(ctx context.Context) (*TokTok, error) {
	return withHooks(ctx, ttuo.sqlSave, ttuo.mutation, ttuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TokTokUpdateOne) SaveX(ctx context.Context) *TokTok {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TokTokUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TokTokUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ttuo *TokTokUpdateOne) sqlSave(ctx context.Context) (_node *TokTok, err error) {
	_spec := sqlgraph.NewUpdateSpec(toktok.Table, toktok.Columns, sqlgraph.NewFieldSpec(toktok.FieldID, field.TypeInt))
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TokTok.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, toktok.FieldID)
		for _, f := range fields {
			if !toktok.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != toktok.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.DeleteTime(); ok {
		_spec.SetField(toktok.FieldDeleteTime, field.TypeTime, value)
	}
	if ttuo.mutation.DeleteTimeCleared() {
		_spec.ClearField(toktok.FieldDeleteTime, field.TypeTime)
	}
	if ttuo.mutation.ReceiverIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.ReceiverIDTable,
			Columns: []string{toktok.ReceiverIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.ReceiverIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.ReceiverIDTable,
			Columns: []string{toktok.ReceiverIDColumn},
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
	if ttuo.mutation.SenderIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.SenderIDTable,
			Columns: []string{toktok.SenderIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.SenderIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   toktok.SenderIDTable,
			Columns: []string{toktok.SenderIDColumn},
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
	_node = &TokTok{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{toktok.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ttuo.mutation.done = true
	return _node, nil
}