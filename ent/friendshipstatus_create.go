// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipStatusCreate is the builder for creating a FriendshipStatus entity.
type FriendshipStatusCreate struct {
	config
	mutation *FriendshipStatusMutation
	hooks    []Hook
}

// SetDeleteTime sets the "delete_time" field.
func (fsc *FriendshipStatusCreate) SetDeleteTime(t time.Time) *FriendshipStatusCreate {
	fsc.mutation.SetDeleteTime(t)
	return fsc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (fsc *FriendshipStatusCreate) SetNillableDeleteTime(t *time.Time) *FriendshipStatusCreate {
	if t != nil {
		fsc.SetDeleteTime(*t)
	}
	return fsc
}

// SetDescription sets the "description" field.
func (fsc *FriendshipStatusCreate) SetDescription(s string) *FriendshipStatusCreate {
	fsc.mutation.SetDescription(s)
	return fsc
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (fsc *FriendshipStatusCreate) AddFriendshipIDs(ids ...int) *FriendshipStatusCreate {
	fsc.mutation.AddFriendshipIDs(ids...)
	return fsc
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (fsc *FriendshipStatusCreate) AddFriendships(f ...*Friendship) *FriendshipStatusCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsc.AddFriendshipIDs(ids...)
}

// Mutation returns the FriendshipStatusMutation object of the builder.
func (fsc *FriendshipStatusCreate) Mutation() *FriendshipStatusMutation {
	return fsc.mutation
}

// Save creates the FriendshipStatus in the database.
func (fsc *FriendshipStatusCreate) Save(ctx context.Context) (*FriendshipStatus, error) {
	return withHooks(ctx, fsc.sqlSave, fsc.mutation, fsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fsc *FriendshipStatusCreate) SaveX(ctx context.Context) *FriendshipStatus {
	v, err := fsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fsc *FriendshipStatusCreate) Exec(ctx context.Context) error {
	_, err := fsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsc *FriendshipStatusCreate) ExecX(ctx context.Context) {
	if err := fsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsc *FriendshipStatusCreate) check() error {
	if _, ok := fsc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "FriendshipStatus.description"`)}
	}
	return nil
}

func (fsc *FriendshipStatusCreate) sqlSave(ctx context.Context) (*FriendshipStatus, error) {
	if err := fsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fsc.mutation.id = &_node.ID
	fsc.mutation.done = true
	return _node, nil
}

func (fsc *FriendshipStatusCreate) createSpec() (*FriendshipStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &FriendshipStatus{config: fsc.config}
		_spec = sqlgraph.NewCreateSpec(friendshipstatus.Table, sqlgraph.NewFieldSpec(friendshipstatus.FieldID, field.TypeInt))
	)
	if value, ok := fsc.mutation.DeleteTime(); ok {
		_spec.SetField(friendshipstatus.FieldDeleteTime, field.TypeTime, value)
		_node.DeleteTime = value
	}
	if value, ok := fsc.mutation.Description(); ok {
		_spec.SetField(friendshipstatus.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := fsc.mutation.FriendshipsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FriendshipStatusCreateBulk is the builder for creating many FriendshipStatus entities in bulk.
type FriendshipStatusCreateBulk struct {
	config
	err      error
	builders []*FriendshipStatusCreate
}

// Save creates the FriendshipStatus entities in the database.
func (fscb *FriendshipStatusCreateBulk) Save(ctx context.Context) ([]*FriendshipStatus, error) {
	if fscb.err != nil {
		return nil, fscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fscb.builders))
	nodes := make([]*FriendshipStatus, len(fscb.builders))
	mutators := make([]Mutator, len(fscb.builders))
	for i := range fscb.builders {
		func(i int, root context.Context) {
			builder := fscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipStatusMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fscb *FriendshipStatusCreateBulk) SaveX(ctx context.Context) []*FriendshipStatus {
	v, err := fscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fscb *FriendshipStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := fscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fscb *FriendshipStatusCreateBulk) ExecX(ctx context.Context) {
	if err := fscb.Exec(ctx); err != nil {
		panic(err)
	}
}