package mymixin

import (
	"context"
	"fmt"
	"notification-service/ent/hook"
	"notification-service/ent/intercept"
	"time"

	gen "notification-service/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Time implements the soft delete pattern for schemas.
type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return []ent.Field{
		// field.Time("created_at").
		// 	Immutable().
		// 	Default(time.Now),
		// field.Time("updated_at").
		// 	Default(time.Now).
		// 	UpdateDefault(time.Now),
		field.Time("delete_time").
			Optional(),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the Time.
func (t Time) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			t.P(q)
			return nil
		}),
	}
}

// Hooks of the Time.
func (t Time) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					// Skip soft-delete, means delete the entity permanently.
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}

					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetDeleteTime(time.Time)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					t.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetDeleteTime(time.Now())
					return mx.Client().Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (t Time) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(t.Fields()[0].Descriptor().Name),
	)
}
