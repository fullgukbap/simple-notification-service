package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// NotificationChange holds the schema definition for the NotificationChange entity.
type NotificationChange struct {
	ent.Schema
}

func (NotificationChange) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the NotificationChange.
func (NotificationChange) Fields() []ent.Field {
	return nil
}

// Edges of the NotificationChange.
func (NotificationChange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("actor", User.Type).
			Ref("notificationChanges").
			Unique(),

		edge.From("notificationObject", NotificationObject.Type).
			Ref("notificationChanges").
			Unique(),
	}
}
