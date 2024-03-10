// Code generated by ent, DO NOT EDIT.

package notificationchange

import (
	"notification-service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldLTE(FieldID, id))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.NotificationChange {
	return predicate.NotificationChange(sql.FieldNotNull(FieldDeleteTime))
}

// HasUserID applies the HasEdge predicate on the "userID" edge.
func HasUserID() predicate.NotificationChange {
	return predicate.NotificationChange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "userID" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.User) predicate.NotificationChange {
	return predicate.NotificationChange(func(s *sql.Selector) {
		step := newUserIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNotificationObjectID applies the HasEdge predicate on the "notificationObjectID" edge.
func HasNotificationObjectID() predicate.NotificationChange {
	return predicate.NotificationChange(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NotificationObjectIDTable, NotificationObjectIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotificationObjectIDWith applies the HasEdge predicate on the "notificationObjectID" edge with a given conditions (other predicates).
func HasNotificationObjectIDWith(preds ...predicate.NotificationObjectID) predicate.NotificationChange {
	return predicate.NotificationChange(func(s *sql.Selector) {
		step := newNotificationObjectIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NotificationChange) predicate.NotificationChange {
	return predicate.NotificationChange(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NotificationChange) predicate.NotificationChange {
	return predicate.NotificationChange(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NotificationChange) predicate.NotificationChange {
	return predicate.NotificationChange(sql.NotPredicates(p))
}