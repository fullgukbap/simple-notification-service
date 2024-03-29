// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"notification-service/ent/notificationobject"
	"notification-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationObjectDelete is the builder for deleting a NotificationObject entity.
type NotificationObjectDelete struct {
	config
	hooks    []Hook
	mutation *NotificationObjectMutation
}

// Where appends a list predicates to the NotificationObjectDelete builder.
func (nod *NotificationObjectDelete) Where(ps ...predicate.NotificationObject) *NotificationObjectDelete {
	nod.mutation.Where(ps...)
	return nod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nod *NotificationObjectDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nod.sqlExec, nod.mutation, nod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nod *NotificationObjectDelete) ExecX(ctx context.Context) int {
	n, err := nod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nod *NotificationObjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notificationobject.Table, sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt))
	if ps := nod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nod.mutation.done = true
	return affected, err
}

// NotificationObjectDeleteOne is the builder for deleting a single NotificationObject entity.
type NotificationObjectDeleteOne struct {
	nod *NotificationObjectDelete
}

// Where appends a list predicates to the NotificationObjectDelete builder.
func (nodo *NotificationObjectDeleteOne) Where(ps ...predicate.NotificationObject) *NotificationObjectDeleteOne {
	nodo.nod.mutation.Where(ps...)
	return nodo
}

// Exec executes the deletion query.
func (nodo *NotificationObjectDeleteOne) Exec(ctx context.Context) error {
	n, err := nodo.nod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationobject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nodo *NotificationObjectDeleteOne) ExecX(ctx context.Context) {
	if err := nodo.Exec(ctx); err != nil {
		panic(err)
	}
}
