// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"notification-service/ent/notificationobjectid"
	"notification-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationObjectIDDelete is the builder for deleting a NotificationObjectID entity.
type NotificationObjectIDDelete struct {
	config
	hooks    []Hook
	mutation *NotificationObjectIDMutation
}

// Where appends a list predicates to the NotificationObjectIDDelete builder.
func (noid *NotificationObjectIDDelete) Where(ps ...predicate.NotificationObjectID) *NotificationObjectIDDelete {
	noid.mutation.Where(ps...)
	return noid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (noid *NotificationObjectIDDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, noid.sqlExec, noid.mutation, noid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (noid *NotificationObjectIDDelete) ExecX(ctx context.Context) int {
	n, err := noid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (noid *NotificationObjectIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(notificationobjectid.Table, sqlgraph.NewFieldSpec(notificationobjectid.FieldID, field.TypeInt))
	if ps := noid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, noid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	noid.mutation.done = true
	return affected, err
}

// NotificationObjectIDDeleteOne is the builder for deleting a single NotificationObjectID entity.
type NotificationObjectIDDeleteOne struct {
	noid *NotificationObjectIDDelete
}

// Where appends a list predicates to the NotificationObjectIDDelete builder.
func (noido *NotificationObjectIDDeleteOne) Where(ps ...predicate.NotificationObjectID) *NotificationObjectIDDeleteOne {
	noido.noid.mutation.Where(ps...)
	return noido
}

// Exec executes the deletion query.
func (noido *NotificationObjectIDDeleteOne) Exec(ctx context.Context) error {
	n, err := noido.noid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{notificationobjectid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (noido *NotificationObjectIDDeleteOne) ExecX(ctx context.Context) {
	if err := noido.Exec(ctx); err != nil {
		panic(err)
	}
}
