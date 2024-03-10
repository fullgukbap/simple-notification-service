// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"
	"notification-service/ent"
	"notification-service/ent/entitytype"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"notification-service/ent/notification"
	"notification-service/ent/notificationchange"
	"notification-service/ent/notificationobjectid"
	"notification-service/ent/predicate"
	"notification-service/ent/user"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The EntityTypeFunc type is an adapter to allow the use of ordinary function as a Querier.
type EntityTypeFunc func(context.Context, *ent.EntityTypeQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f EntityTypeFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.EntityTypeQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.EntityTypeQuery", q)
}

// The TraverseEntityType type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEntityType func(context.Context, *ent.EntityTypeQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEntityType) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEntityType) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.EntityTypeQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.EntityTypeQuery", q)
}

// The FriendshipFunc type is an adapter to allow the use of ordinary function as a Querier.
type FriendshipFunc func(context.Context, *ent.FriendshipQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f FriendshipFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.FriendshipQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.FriendshipQuery", q)
}

// The TraverseFriendship type is an adapter to allow the use of ordinary function as Traverser.
type TraverseFriendship func(context.Context, *ent.FriendshipQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFriendship) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFriendship) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FriendshipQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.FriendshipQuery", q)
}

// The FriendshipStatusFunc type is an adapter to allow the use of ordinary function as a Querier.
type FriendshipStatusFunc func(context.Context, *ent.FriendshipStatusQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f FriendshipStatusFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.FriendshipStatusQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.FriendshipStatusQuery", q)
}

// The TraverseFriendshipStatus type is an adapter to allow the use of ordinary function as Traverser.
type TraverseFriendshipStatus func(context.Context, *ent.FriendshipStatusQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFriendshipStatus) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFriendshipStatus) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FriendshipStatusQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.FriendshipStatusQuery", q)
}

// The NotificationFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotificationFunc func(context.Context, *ent.NotificationQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotificationFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotificationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotificationQuery", q)
}

// The TraverseNotification type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotification func(context.Context, *ent.NotificationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotification) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotification) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotificationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotificationQuery", q)
}

// The NotificationChangeFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotificationChangeFunc func(context.Context, *ent.NotificationChangeQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotificationChangeFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotificationChangeQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotificationChangeQuery", q)
}

// The TraverseNotificationChange type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotificationChange func(context.Context, *ent.NotificationChangeQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotificationChange) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotificationChange) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotificationChangeQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotificationChangeQuery", q)
}

// The NotificationObjectIDFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotificationObjectIDFunc func(context.Context, *ent.NotificationObjectIDQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotificationObjectIDFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotificationObjectIDQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotificationObjectIDQuery", q)
}

// The TraverseNotificationObjectID type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotificationObjectID func(context.Context, *ent.NotificationObjectIDQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotificationObjectID) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotificationObjectID) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotificationObjectIDQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotificationObjectIDQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.EntityTypeQuery:
		return &query[*ent.EntityTypeQuery, predicate.EntityType, entitytype.OrderOption]{typ: ent.TypeEntityType, tq: q}, nil
	case *ent.FriendshipQuery:
		return &query[*ent.FriendshipQuery, predicate.Friendship, friendship.OrderOption]{typ: ent.TypeFriendship, tq: q}, nil
	case *ent.FriendshipStatusQuery:
		return &query[*ent.FriendshipStatusQuery, predicate.FriendshipStatus, friendshipstatus.OrderOption]{typ: ent.TypeFriendshipStatus, tq: q}, nil
	case *ent.NotificationQuery:
		return &query[*ent.NotificationQuery, predicate.Notification, notification.OrderOption]{typ: ent.TypeNotification, tq: q}, nil
	case *ent.NotificationChangeQuery:
		return &query[*ent.NotificationChangeQuery, predicate.NotificationChange, notificationchange.OrderOption]{typ: ent.TypeNotificationChange, tq: q}, nil
	case *ent.NotificationObjectIDQuery:
		return &query[*ent.NotificationObjectIDQuery, predicate.NotificationObjectID, notificationobjectid.OrderOption]{typ: ent.TypeNotificationObjectID, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User, user.OrderOption]{typ: ent.TypeUser, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
