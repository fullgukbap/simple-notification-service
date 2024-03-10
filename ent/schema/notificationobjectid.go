package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotificationObjectID holds the schema definition for the NotificationObjectID entity.
type NotificationObjectID struct {
	ent.Schema
}

func (NotificationObjectID) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the NotificationObjectID.
func (NotificationObjectID) Fields() []ent.Field {
	return []ent.Field{
		field.Int("entity_id").
			Positive(),
	}
}

// Edges of the NotificationObjectID.
func (NotificationObjectID) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notifications", Notification.Type),
		edge.To("notificationChanges", NotificationChange.Type),

		edge.From("entityTypeID", EntityType.Type).
			Ref("notificationObjectIDs").
			Unique(),
	}
}
