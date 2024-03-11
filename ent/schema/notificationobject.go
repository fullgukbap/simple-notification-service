package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// NotificationObject holds the schema definition for the NotificationObject entity.
type NotificationObject struct {
	ent.Schema
}

func (NotificationObject) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the NotificationObject.
func (NotificationObject) Fields() []ent.Field {
	return []ent.Field{
		field.Int("entity_id").
			Positive(),
	}
}

// Edges of the NotificationObject.
func (NotificationObject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("notifications", Notification.Type),
		edge.To("notificationChanges", NotificationChange.Type),

		edge.From("entityType", EntityType.Type).
			Ref("notificationObjects").
			Unique(),
	}
}
