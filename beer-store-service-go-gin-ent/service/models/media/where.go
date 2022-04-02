// Code generated by entc, DO NOT EDIT.

package media

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Creationdatemedia applies equality check predicate on the "creationdatemedia" field. It's identical to CreationdatemediaEQ.
func Creationdatemedia(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationdatemedia), v))
	})
}

// Datamedia applies equality check predicate on the "datamedia" field. It's identical to DatamediaEQ.
func Datamedia(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatamedia), v))
	})
}

// Nomemedia applies equality check predicate on the "nomemedia" field. It's identical to NomemediaEQ.
func Nomemedia(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNomemedia), v))
	})
}

// Mimemedia applies equality check predicate on the "mimemedia" field. It's identical to MimemediaEQ.
func Mimemedia(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMimemedia), v))
	})
}

// CreationdatemediaEQ applies the EQ predicate on the "creationdatemedia" field.
func CreationdatemediaEQ(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationdatemedia), v))
	})
}

// CreationdatemediaNEQ applies the NEQ predicate on the "creationdatemedia" field.
func CreationdatemediaNEQ(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreationdatemedia), v))
	})
}

// CreationdatemediaIn applies the In predicate on the "creationdatemedia" field.
func CreationdatemediaIn(vs ...time.Time) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreationdatemedia), v...))
	})
}

// CreationdatemediaNotIn applies the NotIn predicate on the "creationdatemedia" field.
func CreationdatemediaNotIn(vs ...time.Time) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreationdatemedia), v...))
	})
}

// CreationdatemediaGT applies the GT predicate on the "creationdatemedia" field.
func CreationdatemediaGT(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreationdatemedia), v))
	})
}

// CreationdatemediaGTE applies the GTE predicate on the "creationdatemedia" field.
func CreationdatemediaGTE(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreationdatemedia), v))
	})
}

// CreationdatemediaLT applies the LT predicate on the "creationdatemedia" field.
func CreationdatemediaLT(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreationdatemedia), v))
	})
}

// CreationdatemediaLTE applies the LTE predicate on the "creationdatemedia" field.
func CreationdatemediaLTE(v time.Time) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreationdatemedia), v))
	})
}

// DatamediaEQ applies the EQ predicate on the "datamedia" field.
func DatamediaEQ(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatamedia), v))
	})
}

// DatamediaNEQ applies the NEQ predicate on the "datamedia" field.
func DatamediaNEQ(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatamedia), v))
	})
}

// DatamediaIn applies the In predicate on the "datamedia" field.
func DatamediaIn(vs ...[]byte) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDatamedia), v...))
	})
}

// DatamediaNotIn applies the NotIn predicate on the "datamedia" field.
func DatamediaNotIn(vs ...[]byte) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDatamedia), v...))
	})
}

// DatamediaGT applies the GT predicate on the "datamedia" field.
func DatamediaGT(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatamedia), v))
	})
}

// DatamediaGTE applies the GTE predicate on the "datamedia" field.
func DatamediaGTE(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatamedia), v))
	})
}

// DatamediaLT applies the LT predicate on the "datamedia" field.
func DatamediaLT(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatamedia), v))
	})
}

// DatamediaLTE applies the LTE predicate on the "datamedia" field.
func DatamediaLTE(v []byte) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatamedia), v))
	})
}

// NomemediaEQ applies the EQ predicate on the "nomemedia" field.
func NomemediaEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNomemedia), v))
	})
}

// NomemediaNEQ applies the NEQ predicate on the "nomemedia" field.
func NomemediaNEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNomemedia), v))
	})
}

// NomemediaIn applies the In predicate on the "nomemedia" field.
func NomemediaIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNomemedia), v...))
	})
}

// NomemediaNotIn applies the NotIn predicate on the "nomemedia" field.
func NomemediaNotIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNomemedia), v...))
	})
}

// NomemediaGT applies the GT predicate on the "nomemedia" field.
func NomemediaGT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNomemedia), v))
	})
}

// NomemediaGTE applies the GTE predicate on the "nomemedia" field.
func NomemediaGTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNomemedia), v))
	})
}

// NomemediaLT applies the LT predicate on the "nomemedia" field.
func NomemediaLT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNomemedia), v))
	})
}

// NomemediaLTE applies the LTE predicate on the "nomemedia" field.
func NomemediaLTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNomemedia), v))
	})
}

// NomemediaContains applies the Contains predicate on the "nomemedia" field.
func NomemediaContains(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNomemedia), v))
	})
}

// NomemediaHasPrefix applies the HasPrefix predicate on the "nomemedia" field.
func NomemediaHasPrefix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNomemedia), v))
	})
}

// NomemediaHasSuffix applies the HasSuffix predicate on the "nomemedia" field.
func NomemediaHasSuffix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNomemedia), v))
	})
}

// NomemediaEqualFold applies the EqualFold predicate on the "nomemedia" field.
func NomemediaEqualFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNomemedia), v))
	})
}

// NomemediaContainsFold applies the ContainsFold predicate on the "nomemedia" field.
func NomemediaContainsFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNomemedia), v))
	})
}

// MimemediaEQ applies the EQ predicate on the "mimemedia" field.
func MimemediaEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMimemedia), v))
	})
}

// MimemediaNEQ applies the NEQ predicate on the "mimemedia" field.
func MimemediaNEQ(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMimemedia), v))
	})
}

// MimemediaIn applies the In predicate on the "mimemedia" field.
func MimemediaIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMimemedia), v...))
	})
}

// MimemediaNotIn applies the NotIn predicate on the "mimemedia" field.
func MimemediaNotIn(vs ...string) predicate.Media {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Media(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMimemedia), v...))
	})
}

// MimemediaGT applies the GT predicate on the "mimemedia" field.
func MimemediaGT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMimemedia), v))
	})
}

// MimemediaGTE applies the GTE predicate on the "mimemedia" field.
func MimemediaGTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMimemedia), v))
	})
}

// MimemediaLT applies the LT predicate on the "mimemedia" field.
func MimemediaLT(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMimemedia), v))
	})
}

// MimemediaLTE applies the LTE predicate on the "mimemedia" field.
func MimemediaLTE(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMimemedia), v))
	})
}

// MimemediaContains applies the Contains predicate on the "mimemedia" field.
func MimemediaContains(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMimemedia), v))
	})
}

// MimemediaHasPrefix applies the HasPrefix predicate on the "mimemedia" field.
func MimemediaHasPrefix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMimemedia), v))
	})
}

// MimemediaHasSuffix applies the HasSuffix predicate on the "mimemedia" field.
func MimemediaHasSuffix(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMimemedia), v))
	})
}

// MimemediaEqualFold applies the EqualFold predicate on the "mimemedia" field.
func MimemediaEqualFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMimemedia), v))
	})
}

// MimemediaContainsFold applies the ContainsFold predicate on the "mimemedia" field.
func MimemediaContainsFold(v string) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMimemedia), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
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
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		p(s.Not())
	})
}