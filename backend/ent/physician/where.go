// Code generated by entc, DO NOT EDIT.

package physician

import (
	"github.com/Pichais/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// PhysicianName applies equality check predicate on the "PhysicianName" field. It's identical to PhysicianNameEQ.
func PhysicianName(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhysicianName), v))
	})
}

// PhysicianEmail applies equality check predicate on the "PhysicianEmail" field. It's identical to PhysicianEmailEQ.
func PhysicianEmail(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianNameEQ applies the EQ predicate on the "PhysicianName" field.
func PhysicianNameEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameNEQ applies the NEQ predicate on the "PhysicianName" field.
func PhysicianNameNEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameIn applies the In predicate on the "PhysicianName" field.
func PhysicianNameIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhysicianName), v...))
	})
}

// PhysicianNameNotIn applies the NotIn predicate on the "PhysicianName" field.
func PhysicianNameNotIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhysicianName), v...))
	})
}

// PhysicianNameGT applies the GT predicate on the "PhysicianName" field.
func PhysicianNameGT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameGTE applies the GTE predicate on the "PhysicianName" field.
func PhysicianNameGTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameLT applies the LT predicate on the "PhysicianName" field.
func PhysicianNameLT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameLTE applies the LTE predicate on the "PhysicianName" field.
func PhysicianNameLTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameContains applies the Contains predicate on the "PhysicianName" field.
func PhysicianNameContains(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameHasPrefix applies the HasPrefix predicate on the "PhysicianName" field.
func PhysicianNameHasPrefix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameHasSuffix applies the HasSuffix predicate on the "PhysicianName" field.
func PhysicianNameHasSuffix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameEqualFold applies the EqualFold predicate on the "PhysicianName" field.
func PhysicianNameEqualFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhysicianName), v))
	})
}

// PhysicianNameContainsFold applies the ContainsFold predicate on the "PhysicianName" field.
func PhysicianNameContainsFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhysicianName), v))
	})
}

// PhysicianEmailEQ applies the EQ predicate on the "PhysicianEmail" field.
func PhysicianEmailEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailNEQ applies the NEQ predicate on the "PhysicianEmail" field.
func PhysicianEmailNEQ(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailIn applies the In predicate on the "PhysicianEmail" field.
func PhysicianEmailIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhysicianEmail), v...))
	})
}

// PhysicianEmailNotIn applies the NotIn predicate on the "PhysicianEmail" field.
func PhysicianEmailNotIn(vs ...string) predicate.Physician {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Physician(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhysicianEmail), v...))
	})
}

// PhysicianEmailGT applies the GT predicate on the "PhysicianEmail" field.
func PhysicianEmailGT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailGTE applies the GTE predicate on the "PhysicianEmail" field.
func PhysicianEmailGTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailLT applies the LT predicate on the "PhysicianEmail" field.
func PhysicianEmailLT(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailLTE applies the LTE predicate on the "PhysicianEmail" field.
func PhysicianEmailLTE(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailContains applies the Contains predicate on the "PhysicianEmail" field.
func PhysicianEmailContains(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailHasPrefix applies the HasPrefix predicate on the "PhysicianEmail" field.
func PhysicianEmailHasPrefix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailHasSuffix applies the HasSuffix predicate on the "PhysicianEmail" field.
func PhysicianEmailHasSuffix(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailEqualFold applies the EqualFold predicate on the "PhysicianEmail" field.
func PhysicianEmailEqualFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhysicianEmail), v))
	})
}

// PhysicianEmailContainsFold applies the ContainsFold predicate on the "PhysicianEmail" field.
func PhysicianEmailContainsFold(v string) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhysicianEmail), v))
	})
}

// HasForuser applies the HasEdge predicate on the "foruser" edge.
func HasForuser() predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ForuserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ForuserTable, ForuserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasForuserWith applies the HasEdge predicate on the "foruser" edge with a given conditions (other predicates).
func HasForuserWith(preds ...predicate.Spaciality) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ForuserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ForuserTable, ForuserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
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
func Not(p predicate.Physician) predicate.Physician {
	return predicate.Physician(func(s *sql.Selector) {
		p(s.Not())
	})
}
