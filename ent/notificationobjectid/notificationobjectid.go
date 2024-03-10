// Code generated by ent, DO NOT EDIT.

package notificationobjectid

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notificationobjectid type in the database.
	Label = "notification_object_id"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgeNotificationChanges holds the string denoting the notificationchanges edge name in mutations.
	EdgeNotificationChanges = "notificationChanges"
	// EdgeEntityTypeID holds the string denoting the entitytypeid edge name in mutations.
	EdgeEntityTypeID = "entityTypeID"
	// Table holds the table name of the notificationobjectid in the database.
	Table = "notification_object_ids"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "notification_object_id_notifications"
	// NotificationChangesTable is the table that holds the notificationChanges relation/edge.
	NotificationChangesTable = "notification_changes"
	// NotificationChangesInverseTable is the table name for the NotificationChange entity.
	// It exists in this package in order to avoid circular dependency with the "notificationchange" package.
	NotificationChangesInverseTable = "notification_changes"
	// NotificationChangesColumn is the table column denoting the notificationChanges relation/edge.
	NotificationChangesColumn = "notification_object_id_notification_changes"
	// EntityTypeIDTable is the table that holds the entityTypeID relation/edge.
	EntityTypeIDTable = "notification_object_ids"
	// EntityTypeIDInverseTable is the table name for the EntityType entity.
	// It exists in this package in order to avoid circular dependency with the "entitytype" package.
	EntityTypeIDInverseTable = "entity_types"
	// EntityTypeIDColumn is the table column denoting the entityTypeID relation/edge.
	EntityTypeIDColumn = "entity_type_notification_object_ids"
)

// Columns holds all SQL columns for notificationobjectid fields.
var Columns = []string{
	FieldID,
	FieldDeleteTime,
	FieldEntityID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notification_object_ids"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"entity_type_notification_object_ids",
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
	// EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	EntityIDValidator func(int) error
)

// OrderOption defines the ordering options for the NotificationObjectID queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByEntityID orders the results by the entity_id field.
func ByEntityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityID, opts...).ToFunc()
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

// ByEntityTypeIDField orders the results by entityTypeID field.
func ByEntityTypeIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEntityTypeIDStep(), sql.OrderByField(field, opts...))
	}
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
func newEntityTypeIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EntityTypeIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EntityTypeIDTable, EntityTypeIDColumn),
	)
}
