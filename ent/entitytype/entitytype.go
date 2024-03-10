// Code generated by ent, DO NOT EDIT.

package entitytype

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the entitytype type in the database.
	Label = "entity_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"
	// FieldEntityName holds the string denoting the entityname field in the database.
	FieldEntityName = "entity_name"
	// FieldNotificationDescription holds the string denoting the notificationdescription field in the database.
	FieldNotificationDescription = "notification_description"
	// EdgeNotificationObjectIDs holds the string denoting the notificationobjectids edge name in mutations.
	EdgeNotificationObjectIDs = "notificationObjectIDs"
	// Table holds the table name of the entitytype in the database.
	Table = "entity_types"
	// NotificationObjectIDsTable is the table that holds the notificationObjectIDs relation/edge.
	NotificationObjectIDsTable = "notification_object_ids"
	// NotificationObjectIDsInverseTable is the table name for the NotificationObjectID entity.
	// It exists in this package in order to avoid circular dependency with the "notificationobjectid" package.
	NotificationObjectIDsInverseTable = "notification_object_ids"
	// NotificationObjectIDsColumn is the table column denoting the notificationObjectIDs relation/edge.
	NotificationObjectIDsColumn = "entity_type_notification_object_ids"
)

// Columns holds all SQL columns for entitytype fields.
var Columns = []string{
	FieldID,
	FieldDeleteTime,
	FieldEntityName,
	FieldNotificationDescription,
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

// OrderOption defines the ordering options for the EntityType queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDeleteTime orders the results by the delete_time field.
func ByDeleteTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteTime, opts...).ToFunc()
}

// ByEntityName orders the results by the entityName field.
func ByEntityName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityName, opts...).ToFunc()
}

// ByNotificationDescription orders the results by the notificationDescription field.
func ByNotificationDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotificationDescription, opts...).ToFunc()
}

// ByNotificationObjectIDsCount orders the results by notificationObjectIDs count.
func ByNotificationObjectIDsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationObjectIDsStep(), opts...)
	}
}

// ByNotificationObjectIDs orders the results by notificationObjectIDs terms.
func ByNotificationObjectIDs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationObjectIDsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNotificationObjectIDsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationObjectIDsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotificationObjectIDsTable, NotificationObjectIDsColumn),
	)
}
