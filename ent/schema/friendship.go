package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Friendship holds the schema definition for the Friendship entity.
type Friendship struct {
	ent.Schema
}

func (Friendship) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return nil
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("senderID", User.Type).
			Ref("friendshipsReceiver").
			Unique(),

		edge.From("receiverID", User.Type).
			Ref("friendshipsSender").
			Unique(),

		edge.From("friendshipStatusID", FriendshipStatus.Type).
			Ref("friendships").
			Unique(),
	}
}
