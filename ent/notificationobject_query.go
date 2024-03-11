// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"notification-service/ent/entitytype"
	"notification-service/ent/notification"
	"notification-service/ent/notificationchange"
	"notification-service/ent/notificationobject"
	"notification-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NotificationObjectQuery is the builder for querying NotificationObject entities.
type NotificationObjectQuery struct {
	config
	ctx                     *QueryContext
	order                   []notificationobject.OrderOption
	inters                  []Interceptor
	predicates              []predicate.NotificationObject
	withNotifications       *NotificationQuery
	withNotificationChanges *NotificationChangeQuery
	withEntityType          *EntityTypeQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotificationObjectQuery builder.
func (noq *NotificationObjectQuery) Where(ps ...predicate.NotificationObject) *NotificationObjectQuery {
	noq.predicates = append(noq.predicates, ps...)
	return noq
}

// Limit the number of records to be returned by this query.
func (noq *NotificationObjectQuery) Limit(limit int) *NotificationObjectQuery {
	noq.ctx.Limit = &limit
	return noq
}

// Offset to start from.
func (noq *NotificationObjectQuery) Offset(offset int) *NotificationObjectQuery {
	noq.ctx.Offset = &offset
	return noq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (noq *NotificationObjectQuery) Unique(unique bool) *NotificationObjectQuery {
	noq.ctx.Unique = &unique
	return noq
}

// Order specifies how the records should be ordered.
func (noq *NotificationObjectQuery) Order(o ...notificationobject.OrderOption) *NotificationObjectQuery {
	noq.order = append(noq.order, o...)
	return noq
}

// QueryNotifications chains the current query on the "notifications" edge.
func (noq *NotificationObjectQuery) QueryNotifications() *NotificationQuery {
	query := (&NotificationClient{config: noq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := noq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := noq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationobject.Table, notificationobject.FieldID, selector),
			sqlgraph.To(notification.Table, notification.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notificationobject.NotificationsTable, notificationobject.NotificationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(noq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNotificationChanges chains the current query on the "notificationChanges" edge.
func (noq *NotificationObjectQuery) QueryNotificationChanges() *NotificationChangeQuery {
	query := (&NotificationChangeClient{config: noq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := noq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := noq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationobject.Table, notificationobject.FieldID, selector),
			sqlgraph.To(notificationchange.Table, notificationchange.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, notificationobject.NotificationChangesTable, notificationobject.NotificationChangesColumn),
		)
		fromU = sqlgraph.SetNeighbors(noq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEntityType chains the current query on the "entityType" edge.
func (noq *NotificationObjectQuery) QueryEntityType() *EntityTypeQuery {
	query := (&EntityTypeClient{config: noq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := noq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := noq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notificationobject.Table, notificationobject.FieldID, selector),
			sqlgraph.To(entitytype.Table, entitytype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notificationobject.EntityTypeTable, notificationobject.EntityTypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(noq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NotificationObject entity from the query.
// Returns a *NotFoundError when no NotificationObject was found.
func (noq *NotificationObjectQuery) First(ctx context.Context) (*NotificationObject, error) {
	nodes, err := noq.Limit(1).All(setContextOp(ctx, noq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notificationobject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (noq *NotificationObjectQuery) FirstX(ctx context.Context) *NotificationObject {
	node, err := noq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotificationObject ID from the query.
// Returns a *NotFoundError when no NotificationObject ID was found.
func (noq *NotificationObjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = noq.Limit(1).IDs(setContextOp(ctx, noq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notificationobject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (noq *NotificationObjectQuery) FirstIDX(ctx context.Context) int {
	id, err := noq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotificationObject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotificationObject entity is found.
// Returns a *NotFoundError when no NotificationObject entities are found.
func (noq *NotificationObjectQuery) Only(ctx context.Context) (*NotificationObject, error) {
	nodes, err := noq.Limit(2).All(setContextOp(ctx, noq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notificationobject.Label}
	default:
		return nil, &NotSingularError{notificationobject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (noq *NotificationObjectQuery) OnlyX(ctx context.Context) *NotificationObject {
	node, err := noq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotificationObject ID in the query.
// Returns a *NotSingularError when more than one NotificationObject ID is found.
// Returns a *NotFoundError when no entities are found.
func (noq *NotificationObjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = noq.Limit(2).IDs(setContextOp(ctx, noq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notificationobject.Label}
	default:
		err = &NotSingularError{notificationobject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (noq *NotificationObjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := noq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotificationObjects.
func (noq *NotificationObjectQuery) All(ctx context.Context) ([]*NotificationObject, error) {
	ctx = setContextOp(ctx, noq.ctx, "All")
	if err := noq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NotificationObject, *NotificationObjectQuery]()
	return withInterceptors[[]*NotificationObject](ctx, noq, qr, noq.inters)
}

// AllX is like All, but panics if an error occurs.
func (noq *NotificationObjectQuery) AllX(ctx context.Context) []*NotificationObject {
	nodes, err := noq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotificationObject IDs.
func (noq *NotificationObjectQuery) IDs(ctx context.Context) (ids []int, err error) {
	if noq.ctx.Unique == nil && noq.path != nil {
		noq.Unique(true)
	}
	ctx = setContextOp(ctx, noq.ctx, "IDs")
	if err = noq.Select(notificationobject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (noq *NotificationObjectQuery) IDsX(ctx context.Context) []int {
	ids, err := noq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (noq *NotificationObjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, noq.ctx, "Count")
	if err := noq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, noq, querierCount[*NotificationObjectQuery](), noq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (noq *NotificationObjectQuery) CountX(ctx context.Context) int {
	count, err := noq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (noq *NotificationObjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, noq.ctx, "Exist")
	switch _, err := noq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (noq *NotificationObjectQuery) ExistX(ctx context.Context) bool {
	exist, err := noq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotificationObjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (noq *NotificationObjectQuery) Clone() *NotificationObjectQuery {
	if noq == nil {
		return nil
	}
	return &NotificationObjectQuery{
		config:                  noq.config,
		ctx:                     noq.ctx.Clone(),
		order:                   append([]notificationobject.OrderOption{}, noq.order...),
		inters:                  append([]Interceptor{}, noq.inters...),
		predicates:              append([]predicate.NotificationObject{}, noq.predicates...),
		withNotifications:       noq.withNotifications.Clone(),
		withNotificationChanges: noq.withNotificationChanges.Clone(),
		withEntityType:          noq.withEntityType.Clone(),
		// clone intermediate query.
		sql:  noq.sql.Clone(),
		path: noq.path,
	}
}

// WithNotifications tells the query-builder to eager-load the nodes that are connected to
// the "notifications" edge. The optional arguments are used to configure the query builder of the edge.
func (noq *NotificationObjectQuery) WithNotifications(opts ...func(*NotificationQuery)) *NotificationObjectQuery {
	query := (&NotificationClient{config: noq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	noq.withNotifications = query
	return noq
}

// WithNotificationChanges tells the query-builder to eager-load the nodes that are connected to
// the "notificationChanges" edge. The optional arguments are used to configure the query builder of the edge.
func (noq *NotificationObjectQuery) WithNotificationChanges(opts ...func(*NotificationChangeQuery)) *NotificationObjectQuery {
	query := (&NotificationChangeClient{config: noq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	noq.withNotificationChanges = query
	return noq
}

// WithEntityType tells the query-builder to eager-load the nodes that are connected to
// the "entityType" edge. The optional arguments are used to configure the query builder of the edge.
func (noq *NotificationObjectQuery) WithEntityType(opts ...func(*EntityTypeQuery)) *NotificationObjectQuery {
	query := (&EntityTypeClient{config: noq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	noq.withEntityType = query
	return noq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NotificationObject.Query().
//		GroupBy(notificationobject.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (noq *NotificationObjectQuery) GroupBy(field string, fields ...string) *NotificationObjectGroupBy {
	noq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotificationObjectGroupBy{build: noq}
	grbuild.flds = &noq.ctx.Fields
	grbuild.label = notificationobject.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.NotificationObject.Query().
//		Select(notificationobject.FieldCreatedAt).
//		Scan(ctx, &v)
func (noq *NotificationObjectQuery) Select(fields ...string) *NotificationObjectSelect {
	noq.ctx.Fields = append(noq.ctx.Fields, fields...)
	sbuild := &NotificationObjectSelect{NotificationObjectQuery: noq}
	sbuild.label = notificationobject.Label
	sbuild.flds, sbuild.scan = &noq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotificationObjectSelect configured with the given aggregations.
func (noq *NotificationObjectQuery) Aggregate(fns ...AggregateFunc) *NotificationObjectSelect {
	return noq.Select().Aggregate(fns...)
}

func (noq *NotificationObjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range noq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, noq); err != nil {
				return err
			}
		}
	}
	for _, f := range noq.ctx.Fields {
		if !notificationobject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if noq.path != nil {
		prev, err := noq.path(ctx)
		if err != nil {
			return err
		}
		noq.sql = prev
	}
	return nil
}

func (noq *NotificationObjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotificationObject, error) {
	var (
		nodes       = []*NotificationObject{}
		withFKs     = noq.withFKs
		_spec       = noq.querySpec()
		loadedTypes = [3]bool{
			noq.withNotifications != nil,
			noq.withNotificationChanges != nil,
			noq.withEntityType != nil,
		}
	)
	if noq.withEntityType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, notificationobject.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NotificationObject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NotificationObject{config: noq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, noq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := noq.withNotifications; query != nil {
		if err := noq.loadNotifications(ctx, query, nodes,
			func(n *NotificationObject) { n.Edges.Notifications = []*Notification{} },
			func(n *NotificationObject, e *Notification) { n.Edges.Notifications = append(n.Edges.Notifications, e) }); err != nil {
			return nil, err
		}
	}
	if query := noq.withNotificationChanges; query != nil {
		if err := noq.loadNotificationChanges(ctx, query, nodes,
			func(n *NotificationObject) { n.Edges.NotificationChanges = []*NotificationChange{} },
			func(n *NotificationObject, e *NotificationChange) {
				n.Edges.NotificationChanges = append(n.Edges.NotificationChanges, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := noq.withEntityType; query != nil {
		if err := noq.loadEntityType(ctx, query, nodes, nil,
			func(n *NotificationObject, e *EntityType) { n.Edges.EntityType = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (noq *NotificationObjectQuery) loadNotifications(ctx context.Context, query *NotificationQuery, nodes []*NotificationObject, init func(*NotificationObject), assign func(*NotificationObject, *Notification)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NotificationObject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Notification(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notificationobject.NotificationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.notification_object_notifications
		if fk == nil {
			return fmt.Errorf(`foreign-key "notification_object_notifications" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notification_object_notifications" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (noq *NotificationObjectQuery) loadNotificationChanges(ctx context.Context, query *NotificationChangeQuery, nodes []*NotificationObject, init func(*NotificationObject), assign func(*NotificationObject, *NotificationChange)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NotificationObject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.NotificationChange(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(notificationobject.NotificationChangesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.notification_object_notification_changes
		if fk == nil {
			return fmt.Errorf(`foreign-key "notification_object_notification_changes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "notification_object_notification_changes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (noq *NotificationObjectQuery) loadEntityType(ctx context.Context, query *EntityTypeQuery, nodes []*NotificationObject, init func(*NotificationObject), assign func(*NotificationObject, *EntityType)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*NotificationObject)
	for i := range nodes {
		if nodes[i].entity_type_notification_objects == nil {
			continue
		}
		fk := *nodes[i].entity_type_notification_objects
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(entitytype.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "entity_type_notification_objects" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (noq *NotificationObjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := noq.querySpec()
	_spec.Node.Columns = noq.ctx.Fields
	if len(noq.ctx.Fields) > 0 {
		_spec.Unique = noq.ctx.Unique != nil && *noq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, noq.driver, _spec)
}

func (noq *NotificationObjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notificationobject.Table, notificationobject.Columns, sqlgraph.NewFieldSpec(notificationobject.FieldID, field.TypeInt))
	_spec.From = noq.sql
	if unique := noq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if noq.path != nil {
		_spec.Unique = true
	}
	if fields := noq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notificationobject.FieldID)
		for i := range fields {
			if fields[i] != notificationobject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := noq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := noq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := noq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := noq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (noq *NotificationObjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(noq.driver.Dialect())
	t1 := builder.Table(notificationobject.Table)
	columns := noq.ctx.Fields
	if len(columns) == 0 {
		columns = notificationobject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if noq.sql != nil {
		selector = noq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if noq.ctx.Unique != nil && *noq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range noq.predicates {
		p(selector)
	}
	for _, p := range noq.order {
		p(selector)
	}
	if offset := noq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := noq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotificationObjectGroupBy is the group-by builder for NotificationObject entities.
type NotificationObjectGroupBy struct {
	selector
	build *NotificationObjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nogb *NotificationObjectGroupBy) Aggregate(fns ...AggregateFunc) *NotificationObjectGroupBy {
	nogb.fns = append(nogb.fns, fns...)
	return nogb
}

// Scan applies the selector query and scans the result into the given value.
func (nogb *NotificationObjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nogb.build.ctx, "GroupBy")
	if err := nogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationObjectQuery, *NotificationObjectGroupBy](ctx, nogb.build, nogb, nogb.build.inters, v)
}

func (nogb *NotificationObjectGroupBy) sqlScan(ctx context.Context, root *NotificationObjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nogb.fns))
	for _, fn := range nogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nogb.flds)+len(nogb.fns))
		for _, f := range *nogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotificationObjectSelect is the builder for selecting fields of NotificationObject entities.
type NotificationObjectSelect struct {
	*NotificationObjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nos *NotificationObjectSelect) Aggregate(fns ...AggregateFunc) *NotificationObjectSelect {
	nos.fns = append(nos.fns, fns...)
	return nos
}

// Scan applies the selector query and scans the result into the given value.
func (nos *NotificationObjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nos.ctx, "Select")
	if err := nos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotificationObjectQuery, *NotificationObjectSelect](ctx, nos.NotificationObjectQuery, nos, nos.inters, v)
}

func (nos *NotificationObjectSelect) sqlScan(ctx context.Context, root *NotificationObjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nos.fns))
	for _, fn := range nos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}