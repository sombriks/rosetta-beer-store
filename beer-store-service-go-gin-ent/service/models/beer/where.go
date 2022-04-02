// Code generated by entc, DO NOT EDIT.

package beer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Creationdatebeer applies equality check predicate on the "creationdatebeer" field. It's identical to CreationdatebeerEQ.
func Creationdatebeer(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationdatebeer), v))
	})
}

// Titlebeer applies equality check predicate on the "titlebeer" field. It's identical to TitlebeerEQ.
func Titlebeer(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitlebeer), v))
	})
}

// Descriptionbeer applies equality check predicate on the "descriptionbeer" field. It's identical to DescriptionbeerEQ.
func Descriptionbeer(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescriptionbeer), v))
	})
}

// Idmedia applies equality check predicate on the "idmedia" field. It's identical to IdmediaEQ.
func Idmedia(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdmedia), v))
	})
}

// CreationdatebeerEQ applies the EQ predicate on the "creationdatebeer" field.
func CreationdatebeerEQ(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationdatebeer), v))
	})
}

// CreationdatebeerNEQ applies the NEQ predicate on the "creationdatebeer" field.
func CreationdatebeerNEQ(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreationdatebeer), v))
	})
}

// CreationdatebeerIn applies the In predicate on the "creationdatebeer" field.
func CreationdatebeerIn(vs ...time.Time) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreationdatebeer), v...))
	})
}

// CreationdatebeerNotIn applies the NotIn predicate on the "creationdatebeer" field.
func CreationdatebeerNotIn(vs ...time.Time) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreationdatebeer), v...))
	})
}

// CreationdatebeerGT applies the GT predicate on the "creationdatebeer" field.
func CreationdatebeerGT(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreationdatebeer), v))
	})
}

// CreationdatebeerGTE applies the GTE predicate on the "creationdatebeer" field.
func CreationdatebeerGTE(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreationdatebeer), v))
	})
}

// CreationdatebeerLT applies the LT predicate on the "creationdatebeer" field.
func CreationdatebeerLT(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreationdatebeer), v))
	})
}

// CreationdatebeerLTE applies the LTE predicate on the "creationdatebeer" field.
func CreationdatebeerLTE(v time.Time) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreationdatebeer), v))
	})
}

// TitlebeerEQ applies the EQ predicate on the "titlebeer" field.
func TitlebeerEQ(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerNEQ applies the NEQ predicate on the "titlebeer" field.
func TitlebeerNEQ(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerIn applies the In predicate on the "titlebeer" field.
func TitlebeerIn(vs ...string) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitlebeer), v...))
	})
}

// TitlebeerNotIn applies the NotIn predicate on the "titlebeer" field.
func TitlebeerNotIn(vs ...string) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitlebeer), v...))
	})
}

// TitlebeerGT applies the GT predicate on the "titlebeer" field.
func TitlebeerGT(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerGTE applies the GTE predicate on the "titlebeer" field.
func TitlebeerGTE(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerLT applies the LT predicate on the "titlebeer" field.
func TitlebeerLT(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerLTE applies the LTE predicate on the "titlebeer" field.
func TitlebeerLTE(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerContains applies the Contains predicate on the "titlebeer" field.
func TitlebeerContains(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerHasPrefix applies the HasPrefix predicate on the "titlebeer" field.
func TitlebeerHasPrefix(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerHasSuffix applies the HasSuffix predicate on the "titlebeer" field.
func TitlebeerHasSuffix(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerEqualFold applies the EqualFold predicate on the "titlebeer" field.
func TitlebeerEqualFold(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitlebeer), v))
	})
}

// TitlebeerContainsFold applies the ContainsFold predicate on the "titlebeer" field.
func TitlebeerContainsFold(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitlebeer), v))
	})
}

// DescriptionbeerEQ applies the EQ predicate on the "descriptionbeer" field.
func DescriptionbeerEQ(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerNEQ applies the NEQ predicate on the "descriptionbeer" field.
func DescriptionbeerNEQ(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerIn applies the In predicate on the "descriptionbeer" field.
func DescriptionbeerIn(vs ...string) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescriptionbeer), v...))
	})
}

// DescriptionbeerNotIn applies the NotIn predicate on the "descriptionbeer" field.
func DescriptionbeerNotIn(vs ...string) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescriptionbeer), v...))
	})
}

// DescriptionbeerGT applies the GT predicate on the "descriptionbeer" field.
func DescriptionbeerGT(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerGTE applies the GTE predicate on the "descriptionbeer" field.
func DescriptionbeerGTE(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerLT applies the LT predicate on the "descriptionbeer" field.
func DescriptionbeerLT(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerLTE applies the LTE predicate on the "descriptionbeer" field.
func DescriptionbeerLTE(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerContains applies the Contains predicate on the "descriptionbeer" field.
func DescriptionbeerContains(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerHasPrefix applies the HasPrefix predicate on the "descriptionbeer" field.
func DescriptionbeerHasPrefix(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerHasSuffix applies the HasSuffix predicate on the "descriptionbeer" field.
func DescriptionbeerHasSuffix(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerEqualFold applies the EqualFold predicate on the "descriptionbeer" field.
func DescriptionbeerEqualFold(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescriptionbeer), v))
	})
}

// DescriptionbeerContainsFold applies the ContainsFold predicate on the "descriptionbeer" field.
func DescriptionbeerContainsFold(v string) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescriptionbeer), v))
	})
}

// IdmediaEQ applies the EQ predicate on the "idmedia" field.
func IdmediaEQ(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdmedia), v))
	})
}

// IdmediaNEQ applies the NEQ predicate on the "idmedia" field.
func IdmediaNEQ(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdmedia), v))
	})
}

// IdmediaIn applies the In predicate on the "idmedia" field.
func IdmediaIn(vs ...int) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdmedia), v...))
	})
}

// IdmediaNotIn applies the NotIn predicate on the "idmedia" field.
func IdmediaNotIn(vs ...int) predicate.Beer {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Beer(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdmedia), v...))
	})
}

// IdmediaGT applies the GT predicate on the "idmedia" field.
func IdmediaGT(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdmedia), v))
	})
}

// IdmediaGTE applies the GTE predicate on the "idmedia" field.
func IdmediaGTE(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdmedia), v))
	})
}

// IdmediaLT applies the LT predicate on the "idmedia" field.
func IdmediaLT(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdmedia), v))
	})
}

// IdmediaLTE applies the LTE predicate on the "idmedia" field.
func IdmediaLTE(v int) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdmedia), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Beer) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Beer) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
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
func Not(p predicate.Beer) predicate.Beer {
	return predicate.Beer(func(s *sql.Selector) {
		p(s.Not())
	})
}