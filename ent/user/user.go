// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// EdgeFriendshipsReceiver holds the string denoting the friendshipsreceiver edge name in mutations.
	EdgeFriendshipsReceiver = "friendshipsReceiver"
	// EdgeFriendshipsSender holds the string denoting the friendshipssender edge name in mutations.
	EdgeFriendshipsSender = "friendshipsSender"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgeNotificationChanges holds the string denoting the notificationchanges edge name in mutations.
	EdgeNotificationChanges = "notificationChanges"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendshipsReceiverTable is the table that holds the friendshipsReceiver relation/edge.
	FriendshipsReceiverTable = "friendships"
	// FriendshipsReceiverInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsReceiverInverseTable = "friendships"
	// FriendshipsReceiverColumn is the table column denoting the friendshipsReceiver relation/edge.
	FriendshipsReceiverColumn = "user_friendships_receiver"
	// FriendshipsSenderTable is the table that holds the friendshipsSender relation/edge.
	FriendshipsSenderTable = "friendships"
	// FriendshipsSenderInverseTable is the table name for the Friendship entity.
	// It exists in this package in order to avoid circular dependency with the "friendship" package.
	FriendshipsSenderInverseTable = "friendships"
	// FriendshipsSenderColumn is the table column denoting the friendshipsSender relation/edge.
	FriendshipsSenderColumn = "user_friendships_sender"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "user_notifications"
	// NotificationChangesTable is the table that holds the notificationChanges relation/edge.
	NotificationChangesTable = "notification_changes"
	// NotificationChangesInverseTable is the table name for the NotificationChange entity.
	// It exists in this package in order to avoid circular dependency with the "notificationchange" package.
	NotificationChangesInverseTable = "notification_changes"
	// NotificationChangesColumn is the table column denoting the notificationChanges relation/edge.
	NotificationChangesColumn = "user_notification_changes"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUsername,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByFriendshipsReceiverCount orders the results by friendshipsReceiver count.
func ByFriendshipsReceiverCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendshipsReceiverStep(), opts...)
	}
}

// ByFriendshipsReceiver orders the results by friendshipsReceiver terms.
func ByFriendshipsReceiver(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendshipsReceiverStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFriendshipsSenderCount orders the results by friendshipsSender count.
func ByFriendshipsSenderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendshipsSenderStep(), opts...)
	}
}

// ByFriendshipsSender orders the results by friendshipsSender terms.
func ByFriendshipsSender(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendshipsSenderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotificationChangesCount orders the results by notificationChanges count.
func ByNotificationChangesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationChangesStep(), opts...)
	}
}

// ByNotificationChanges orders the results by notificationChanges terms.
func ByNotificationChanges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationChangesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFriendshipsReceiverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendshipsReceiverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FriendshipsReceiverTable, FriendshipsReceiverColumn),
	)
}
func newFriendshipsSenderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FriendshipsSenderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FriendshipsSenderTable, FriendshipsSenderColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationsTable, NotificationsColumn),
	)
}
func newNotificationChangesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationChangesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationChangesTable, NotificationChangesColumn),
	)
}
