// Code generated by ent, DO NOT EDIT.

package notificationobject

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the notificationobject type in the database.
	Label = "notification_object"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgeNotificationChanges holds the string denoting the notificationchanges edge name in mutations.
	EdgeNotificationChanges = "notificationChanges"
	// EdgeEntityType holds the string denoting the entitytype edge name in mutations.
	EdgeEntityType = "entityType"
	// Table holds the table name of the notificationobject in the database.
	Table = "notification_objects"
	// NotificationsTable is the table that holds the notifications relation/edge.
	NotificationsTable = "notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// NotificationsColumn is the table column denoting the notifications relation/edge.
	NotificationsColumn = "notification_object_notifications"
	// NotificationChangesTable is the table that holds the notificationChanges relation/edge.
	NotificationChangesTable = "notification_changes"
	// NotificationChangesInverseTable is the table name for the NotificationChange entity.
	// It exists in this package in order to avoid circular dependency with the "notificationchange" package.
	NotificationChangesInverseTable = "notification_changes"
	// NotificationChangesColumn is the table column denoting the notificationChanges relation/edge.
	NotificationChangesColumn = "notification_object_notification_changes"
	// EntityTypeTable is the table that holds the entityType relation/edge.
	EntityTypeTable = "notification_objects"
	// EntityTypeInverseTable is the table name for the EntityType entity.
	// It exists in this package in order to avoid circular dependency with the "entitytype" package.
	EntityTypeInverseTable = "entity_types"
	// EntityTypeColumn is the table column denoting the entityType relation/edge.
	EntityTypeColumn = "entity_type_notification_objects"
)

// Columns holds all SQL columns for notificationobject fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntityID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notification_objects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"entity_type_notification_objects",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	EntityIDValidator func(int) error
)

// OrderOption defines the ordering options for the NotificationObject queries.
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

// ByEntityTypeField orders the results by entityType field.
func ByEntityTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEntityTypeStep(), sql.OrderByField(field, opts...))
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
func newEntityTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EntityTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EntityTypeTable, EntityTypeColumn),
	)
}