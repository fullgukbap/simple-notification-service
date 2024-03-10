package schema

import (
	"notification-service/ent/schema/mymixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// TokTok holds the schema definition for the TokTok entity.
type TokTok struct {
	ent.Schema
}

func (TokTok) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mymixin.Time{},
	}
}

// Fields of the TokTok.
func (TokTok) Fields() []ent.Field {
	return nil
}

// Edges of the TokTok.
func (TokTok) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("receiverID", User.Type).
			Ref("toktoks_receiver").
			Unique(),

		edge.From("senderID", User.Type).
			Ref("toktoks_sender").
			Unique(),
	}
}
