// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package car

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/start/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// RegisteredAt applies equality check predicate on the "registered_at" field. It's identical to RegisteredAtEQ.
func RegisteredAt(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisteredAt), v))
	})
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldModel), v))
	})
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldModel), v))
	})
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldModel), v...))
	})
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldModel), v...))
	})
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldModel), v))
	})
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldModel), v))
	})
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldModel), v))
	})
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldModel), v))
	})
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldModel), v))
	})
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldModel), v))
	})
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldModel), v))
	})
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldModel), v))
	})
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldModel), v))
	})
}

// RegisteredAtEQ applies the EQ predicate on the "registered_at" field.
func RegisteredAtEQ(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRegisteredAt), v))
	})
}

// RegisteredAtNEQ applies the NEQ predicate on the "registered_at" field.
func RegisteredAtNEQ(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRegisteredAt), v))
	})
}

// RegisteredAtIn applies the In predicate on the "registered_at" field.
func RegisteredAtIn(vs ...time.Time) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRegisteredAt), v...))
	})
}

// RegisteredAtNotIn applies the NotIn predicate on the "registered_at" field.
func RegisteredAtNotIn(vs ...time.Time) predicate.Car {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Car(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRegisteredAt), v...))
	})
}

// RegisteredAtGT applies the GT predicate on the "registered_at" field.
func RegisteredAtGT(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRegisteredAt), v))
	})
}

// RegisteredAtGTE applies the GTE predicate on the "registered_at" field.
func RegisteredAtGTE(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRegisteredAt), v))
	})
}

// RegisteredAtLT applies the LT predicate on the "registered_at" field.
func RegisteredAtLT(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRegisteredAt), v))
	})
}

// RegisteredAtLTE applies the LTE predicate on the "registered_at" field.
func RegisteredAtLTE(v time.Time) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRegisteredAt), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
