package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FriendshipStatus holds the schema definition for the FriendshipStatus entity.
type FriendshipStatus struct {
	ent.Schema
}

func (FriendshipStatus) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the FriendshipStatus.
func (FriendshipStatus) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			Unique(),
	}
}

// Edges of the FriendshipStatus.
func (FriendshipStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friendships", Friendship.Type),
	}
}
