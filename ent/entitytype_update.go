// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"notification-service/ent/entitytype"
	"notification-service/ent/notificationobject"
	"notification-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntityTypeUpdate is the builder for updating EntityType entities.
type EntityTypeUpdate struct {
	config
	hooks    []Hook
	mutation *EntityTypeMutation
}

// Where appends a list predicates to the EntityTypeUpdate builder.
func (etu *EntityTypeUpdate) Where(ps ...predicate.EntityType) *EntityTypeUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetUpdatedAt sets the "updated_at" field.
func (etu *EntityTypeUpdate) SetUpdatedAt(t time.Time) *EntityTypeUpdate {
	etu.mutation.SetUpdatedAt(t)
	return etu
}

// SetDeletedAt sets the "deleted_at" field.
func (etu *EntityTypeUpdate) SetDeletedAt(t time.Time) *EntityTypeUpdate {
	etu.mutation.SetDeletedAt(t)
	return etu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etu *EntityTypeUpdate) SetNillableDeletedAt(t *time.Time) *EntityTypeUpdate {
	if t != nil {
		etu.SetDeletedAt(*t)
	}
	return etu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (etu *EntityTypeUpdate) ClearDeletedAt() *EntityTypeUpdate {
	etu.mutation.ClearDeletedAt()
	return etu
}

// SetEntityName sets the "entityName" field.
func (etu *EntityTypeUpdate) SetEntityName(s string) *EntityTypeUpdate {
	etu.mutation.SetEntityName(s)
	return etu
}

// SetNillableEntityName sets the "entityName" field if the given value is not nil.
func (etu *EntityTypeUpdate) SetNillableEntityName(s *string) *EntityTypeUpdate {
	if s != nil {
		etu.SetEntityName(*s)
	}
	return etu
}

// SetNotificationDescription sets the "notificationDescription" field.
func (etu *EntityTypeUpdate) SetNotificationDescription(s string) *EntityTypeUpdate {
	etu.mutation.SetNotificationDescription(s)
	return etu
}

// SetNillableNotificationDescription sets the "notificationDescription" field if the given value is not nil.
func (etu *EntityTypeUpdate) SetNillableNotificationDescription(s *string) *EntityTypeUpdate {
	if s != nil {
		etu.SetNotificationDescription(*s)
	}
	return etu
}

// AddNotificationObjectIDs adds the "notificationObjects" edge to the NotificationObject entity by IDs.
func (etu *EntityTypeUpdate) AddNotificationObjectIDs(ids ...int) *EntityTypeUpdate {
	etu.mutation.AddNotificationObjectIDs(ids...)
	return etu
}

// AddNotificationObjects adds the "notificationObjects" edges to the NotificationObject entity.
func (etu *EntityTypeUpdate) AddNotificationObjects(n ...*NotificationObject) *EntityTypeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return etu.AddNotificationObjectIDs(ids...)
}

// Mutation returns the EntityTypeMutation object of the builder.
func (etu *EntityTypeUpdate) Mutation() *EntityTypeMutation {
	return etu.mutation
}

// ClearNotificationObjects clears all "notificationObjects" edges to the NotificationObject entity.
func (etu *EntityTypeUpdate) ClearNotificationObjects() *EntityTypeUpdate {
	etu.mutation.ClearNotificationObjects()
	return etu
}

// RemoveNotificationObjectIDs removes the "notificationObjects" edge to NotificationObject entities by IDs.
func (etu *EntityTypeUpdate) RemoveNotificationObjectIDs(ids ...int) *EntityTypeUpdate {
	etu.mutation.RemoveNotificationObjectIDs(ids...)
	return etu
}

// RemoveNotificationObjects removes "notificationObjects" edges to NotificationObject entities.
func (etu *EntityTypeUpdate) RemoveNotificationObjects(n ...*NotificationObject) *EntityTypeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return etu.RemoveNotificationObjectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EntityTypeUpdate) Save(ctx context.Context) (int, error) {
	if err := etu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EntityTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EntityTypeUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EntityTypeUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etu *EntityTypeUpdate) defaults() error {
	if _, ok := etu.mutation.UpdatedAt(); !ok {
		if entitytype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entitytype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entitytype.UpdateDefaultUpdatedAt()
		etu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (etu *EntityTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(entitytype.Table, entitytype.Columns, sqlgraph.NewFieldSpec(entitytype.FieldID, field.TypeInt))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.UpdatedAt(); ok {
		_spec.SetField(entitytype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := etu.mutation.DeletedAt(); ok {
		_spec.SetField(entitytype.FieldDeletedAt, field.TypeTime, value)
	}
	if etu.mutation.DeletedAtCleared() {
		_spec.ClearField(entitytype.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := etu.mutation.EntityName(); ok {
		_spec.SetField(entitytype.FieldEntityName, field.TypeString, value)
	}
	if value, ok := etu.mutation.NotificationDescription(); ok {
		_spec.SetField(entitytype.FieldNotificationDescription, field.TypeString, value)
	}
	if etu.mutation.NotificationObjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.RemovedNotificationObjectsIDs(); len(nodes) > 0 && !etu.mutation.NotificationObjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.NotificationObjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// EntityTypeUpdateOne is the builder for updating a single EntityType entity.
type EntityTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityTypeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (etuo *EntityTypeUpdateOne) SetUpdatedAt(t time.Time) *EntityTypeUpdateOne {
	etuo.mutation.SetUpdatedAt(t)
	return etuo
}

// SetDeletedAt sets the "deleted_at" field.
func (etuo *EntityTypeUpdateOne) SetDeletedAt(t time.Time) *EntityTypeUpdateOne {
	etuo.mutation.SetDeletedAt(t)
	return etuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (etuo *EntityTypeUpdateOne) SetNillableDeletedAt(t *time.Time) *EntityTypeUpdateOne {
	if t != nil {
		etuo.SetDeletedAt(*t)
	}
	return etuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (etuo *EntityTypeUpdateOne) ClearDeletedAt() *EntityTypeUpdateOne {
	etuo.mutation.ClearDeletedAt()
	return etuo
}

// SetEntityName sets the "entityName" field.
func (etuo *EntityTypeUpdateOne) SetEntityName(s string) *EntityTypeUpdateOne {
	etuo.mutation.SetEntityName(s)
	return etuo
}

// SetNillableEntityName sets the "entityName" field if the given value is not nil.
func (etuo *EntityTypeUpdateOne) SetNillableEntityName(s *string) *EntityTypeUpdateOne {
	if s != nil {
		etuo.SetEntityName(*s)
	}
	return etuo
}

// SetNotificationDescription sets the "notificationDescription" field.
func (etuo *EntityTypeUpdateOne) SetNotificationDescription(s string) *EntityTypeUpdateOne {
	etuo.mutation.SetNotificationDescription(s)
	return etuo
}

// SetNillableNotificationDescription sets the "notificationDescription" field if the given value is not nil.
func (etuo *EntityTypeUpdateOne) SetNillableNotificationDescription(s *string) *EntityTypeUpdateOne {
	if s != nil {
		etuo.SetNotificationDescription(*s)
	}
	return etuo
}

// AddNotificationObjectIDs adds the "notificationObjects" edge to the NotificationObject entity by IDs.
func (etuo *EntityTypeUpdateOne) AddNotificationObjectIDs(ids ...int) *EntityTypeUpdateOne {
	etuo.mutation.AddNotificationObjectIDs(ids...)
	return etuo
}

// AddNotificationObjects adds the "notificationObjects" edges to the NotificationObject entity.
func (etuo *EntityTypeUpdateOne) AddNotificationObjects(n ...*NotificationObject) *EntityTypeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return etuo.AddNotificationObjectIDs(ids...)
}

// Mutation returns the EntityTypeMutation object of the builder.
func (etuo *EntityTypeUpdateOne) Mutation() *EntityTypeMutation {
	return etuo.mutation
}

// ClearNotificationObjects clears all "notificationObjects" edges to the NotificationObject entity.
func (etuo *EntityTypeUpdateOne) ClearNotificationObjects() *EntityTypeUpdateOne {
	etuo.mutation.ClearNotificationObjects()
	return etuo
}

// RemoveNotificationObjectIDs removes the "notificationObjects" edge to NotificationObject entities by IDs.
func (etuo *EntityTypeUpdateOne) RemoveNotificationObjectIDs(ids ...int) *EntityTypeUpdateOne {
	etuo.mutation.RemoveNotificationObjectIDs(ids...)
	return etuo
}

// RemoveNotificationObjects removes "notificationObjects" edges to NotificationObject entities.
func (etuo *EntityTypeUpdateOne) RemoveNotificationObjects(n ...*NotificationObject) *EntityTypeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return etuo.RemoveNotificationObjectIDs(ids...)
}

// Where appends a list predicates to the EntityTypeUpdate builder.
func (etuo *EntityTypeUpdateOne) Where(ps ...predicate.EntityType) *EntityTypeUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EntityTypeUpdateOne) Select(field string, fields ...string) *EntityTypeUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EntityType entity.
func (etuo *EntityTypeUpdateOne) Save(ctx context.Context) (*EntityType, error) {
	if err := etuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EntityTypeUpdateOne) SaveX(ctx context.Context) *EntityType {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EntityTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EntityTypeUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etuo *EntityTypeUpdateOne) defaults() error {
	if _, ok := etuo.mutation.UpdatedAt(); !ok {
		if entitytype.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized entitytype.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := entitytype.UpdateDefaultUpdatedAt()
		etuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (etuo *EntityTypeUpdateOne) sqlSave(ctx context.Context) (_node *EntityType, err error) {
	_spec := sqlgraph.NewUpdateSpec(entitytype.Table, entitytype.Columns, sqlgraph.NewFieldSpec(entitytype.FieldID, field.TypeInt))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntityType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitytype.FieldID)
		for _, f := range fields {
			if !entitytype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entitytype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entitytype.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := etuo.mutation.DeletedAt(); ok {
		_spec.SetField(entitytype.FieldDeletedAt, field.TypeTime, value)
	}
	if etuo.mutation.DeletedAtCleared() {
		_spec.ClearField(entitytype.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := etuo.mutation.EntityName(); ok {
		_spec.SetField(entitytype.FieldEntityName, field.TypeString, value)
	}
	if value, ok := etuo.mutation.NotificationDescription(); ok {
		_spec.SetField(entitytype.FieldNotificationDescription, field.TypeString, value)
	}
	if etuo.mutation.NotificationObjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.RemovedNotificationObjectsIDs(); len(nodes) > 0 && !etuo.mutation.NotificationObjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.NotificationObjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   entitytype.NotificationObjectsTable,
			Columns: []string{entitytype.NotificationObjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntityType{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitytype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}
