package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EntityType holds the schema definition for the EntityType entity.
type EntityType struct {
	ent.Schema
}

func (EntityType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the EntityType.
func (EntityType) Fields() []ent.Field {
	return []ent.Field{
		field.String("entityName"),
		field.String("notificationDescription"),
	}
}

// Edges of the EntityType.
func (EntityType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notificationObjectIDs", NotificationObjectID.Type),
	}
}
