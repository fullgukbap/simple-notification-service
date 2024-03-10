// Code generated by ent, DO NOT EDIT.

package friendship

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the friendship type in the database.
	Label = "friendship"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// EdgeSenderID holds the string denoting the senderid edge name in mutations.
	EdgeSenderID = "senderID"
	// EdgeReceiverID holds the string denoting the receiverid edge name in mutations.
	EdgeReceiverID = "receiverID"
	// EdgeFriendshipStatusID holds the string denoting the friendshipstatusid edge name in mutations.
	EdgeFriendshipStatusID = "friendshipStatusID"
	// Table holds the table name of the friendship in the database.
	Table = "friendships"
	// SenderIDTable is the table that holds the senderID relation/edge.
	SenderIDTable = "friendships"
	// SenderIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SenderIDInverseTable = "users"
	// SenderIDColumn is the table column denoting the senderID relation/edge.
	SenderIDColumn = "user_friendships_receiver"
	// ReceiverIDTable is the table that holds the receiverID relation/edge.
	ReceiverIDTable = "friendships"
	// ReceiverIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReceiverIDInverseTable = "users"
	// ReceiverIDColumn is the table column denoting the receiverID relation/edge.
	ReceiverIDColumn = "user_friendships_sender"
	// FriendshipStatusIDTable is the table that holds the friendshipStatusID relation/edge.
	FriendshipStatusIDTable = "friendships"
	// FriendshipStatusIDInverseTable is the table name for the FriendshipStatus entity.
	// It exists in this package in order to avoid circular dependency with the "friendshipstatus" package.
	FriendshipStatusIDInverseTable = "friendship_status"
	// FriendshipStatusIDColumn is the table column denoting the friendshipStatusID relation/edge.
	FriendshipStatusIDColumn = "friendship_status_friendships"
)

// Columns holds all SQL columns for friendship fields.
var Columns = []string{
	FieldID,
	FieldDeleteTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "friendships"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"friendship_status_friendships",
	"user_friendships_receiver",
	"user_friendships_sender",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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

// OrderOption defines the ordering options for the Friendship queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// BySenderIDField orders the results by senderID field.
func BySenderIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderIDStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiverIDField orders the results by receiverID field.
func ByReceiverIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiverIDStep(), sql.OrderByField(field, opts...))
	}
}

// ByFriendshipStatusIDField orders the results by friendshipStatusID field.
func ByFriendshipStatusIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendshipStatusIDStep(), sql.OrderByField(field, opts...))
	}
}
func newSenderIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenderIDTable, SenderIDColumn),
	)
}
func newReceiverIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiverIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReceiverIDTable, ReceiverIDColumn),
	)
}
func newFriendshipStatusIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendshipStatusIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FriendshipStatusIDTable, FriendshipStatusIDColumn),
	)
}