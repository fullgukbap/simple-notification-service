// Code generated by ent, DO NOT EDIT.

package friendship

import (
	"notification-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldID, id))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.Friendship {
	return predicate.Friendship(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.Friendship {
	return predicate.Friendship(sql.FieldNotNull(FieldDeleteTime))
}

// HasSenderID applies the HasEdge predicate on the "senderID" edge.
func HasSenderID() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SenderIDTable, SenderIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSenderIDWith applies the HasEdge predicate on the "senderID" edge with a given conditions (other predicates).
func HasSenderIDWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := newSenderIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceiverID applies the HasEdge predicate on the "receiverID" edge.
func HasReceiverID() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReceiverIDTable, ReceiverIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiverIDWith applies the HasEdge predicate on the "receiverID" edge with a given conditions (other predicates).
func HasReceiverIDWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := newReceiverIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriendshipStatusID applies the HasEdge predicate on the "friendshipStatusID" edge.
func HasFriendshipStatusID() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FriendshipStatusIDTable, FriendshipStatusIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendshipStatusIDWith applies the HasEdge predicate on the "friendshipStatusID" edge with a given conditions (other predicates).
func HasFriendshipStatusIDWith(preds ...predicate.FriendshipStatus) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := newFriendshipStatusIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(sql.NotPredicates(p))
}
