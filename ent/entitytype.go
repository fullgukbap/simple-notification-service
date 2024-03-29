// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"notification-service/ent/entitytype"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EntityType is the model entity for the EntityType schema.
type EntityType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// EntityName holds the value of the "entityName" field.
	EntityName string `json:"entityName,omitempty"`
	// NotificationDescription holds the value of the "notificationDescription" field.
	NotificationDescription string `json:"notificationDescription,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntityTypeQuery when eager-loading is set.
	Edges        EntityTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EntityTypeEdges holds the relations/edges for other nodes in the graph.
type EntityTypeEdges struct {
	// NotificationObjects holds the value of the notificationObjects edge.
	NotificationObjects []*NotificationObject `json:"notificationObjects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// NotificationObjectsOrErr returns the NotificationObjects value or an error if the edge
// was not loaded in eager-loading.
func (e EntityTypeEdges) NotificationObjectsOrErr() ([]*NotificationObject, error) {
	if e.loadedTypes[0] {
		return e.NotificationObjects, nil
	}
	return nil, &NotLoadedError{edge: "notificationObjects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntityType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entitytype.FieldID:
			values[i] = new(sql.NullInt64)
		case entitytype.FieldEntityName, entitytype.FieldNotificationDescription:
			values[i] = new(sql.NullString)
		case entitytype.FieldCreatedAt, entitytype.FieldUpdatedAt, entitytype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntityType fields.
func (et *EntityType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entitytype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			et.ID = int(value.Int64)
		case entitytype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				et.CreatedAt = value.Time
			}
		case entitytype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				et.UpdatedAt = value.Time
			}
		case entitytype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				et.DeletedAt = value.Time
			}
		case entitytype.FieldEntityName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entityName", values[i])
			} else if value.Valid {
				et.EntityName = value.String
			}
		case entitytype.FieldNotificationDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notificationDescription", values[i])
			} else if value.Valid {
				et.NotificationDescription = value.String
			}
		default:
			et.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EntityType.
// This includes values selected through modifiers, order, etc.
func (et *EntityType) Value(name string) (ent.Value, error) {
	return et.selectValues.Get(name)
}

// QueryNotificationObjects queries the "notificationObjects" edge of the EntityType entity.
func (et *EntityType) QueryNotificationObjects() *NotificationObjectQuery {
	return NewEntityTypeClient(et.config).QueryNotificationObjects(et)
}

// Update returns a builder for updating this EntityType.
// Note that you need to call EntityType.Unwrap() before calling this method if this EntityType
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *EntityType) Update() *EntityTypeUpdateOne {
	return NewEntityTypeClient(et.config).UpdateOne(et)
}

// Unwrap unwraps the EntityType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (et *EntityType) Unwrap() *EntityType {
	_tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntityType is not a transactional entity")
	}
	et.config.driver = _tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *EntityType) String() string {
	var builder strings.Builder
	builder.WriteString("EntityType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", et.ID))
	builder.WriteString("created_at=")
	builder.WriteString(et.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(et.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(et.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("entityName=")
	builder.WriteString(et.EntityName)
	builder.WriteString(", ")
	builder.WriteString("notificationDescription=")
	builder.WriteString(et.NotificationDescription)
	builder.WriteByte(')')
	return builder.String()
}

// EntityTypes is a parsable slice of EntityType.
type EntityTypes []*EntityType
