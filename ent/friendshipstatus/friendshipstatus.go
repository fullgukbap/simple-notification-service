// Code generated by ent, DO NOT EDIT.

package friendshipstatus

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the friendshipstatus type in the database.
	Label = "friendship_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeFriendships holds the string denoting the friendships edge name in mutations.
	EdgeFriendships = "friendships"
	// Table holds the table name of the friendshipstatus in the database.
	Table = "friendship_status"
	// FriendshipsTable is the table that holds the friendships relation/edge.
	FriendshipsTable = "friendships"
	// FriendshipsInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsInverseTable = "friendships"
	// FriendshipsColumn is the table column denoting the friendships relation/edge.
	FriendshipsColumn = "friendship_status_friendships"
)

// Columns holds all SQL columns for friendshipstatus fields.
var Columns = []string{
	FieldID,
	FieldDeleteTime,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "notification-service/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
)

// OrderOption defines the ordering options for the FriendshipStatus queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByFriendshipsCount orders the results by friendships count.
func ByFriendshipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendshipsStep(), opts...)
	}
}

// ByFriendships orders the results by friendships terms.
func ByFriendships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendshipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFriendshipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendshipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FriendshipsTable, FriendshipsColumn),
	)
}
