// Code generated by entc, DO NOT EDIT.

package timestamp

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/zackartz/ttd/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
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
func IDNotIn(ids ...uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
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
func IDGT(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// Project applies equality check predicate on the "project" field. It's identical to ProjectEQ.
func Project(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProject), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActive), v))
	})
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActive), v))
	})
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComment), v))
	})
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComment), v))
	})
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldComment), v...))
	})
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldComment), v...))
	})
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComment), v))
	})
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComment), v))
	})
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComment), v))
	})
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComment), v))
	})
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComment), v))
	})
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComment), v))
	})
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComment), v))
	})
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldComment)))
	})
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldComment)))
	})
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComment), v))
	})
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComment), v))
	})
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategory), v))
	})
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCategory), v...))
	})
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCategory), v...))
	})
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCategory), v))
	})
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCategory), v))
	})
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCategory), v))
	})
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCategory), v))
	})
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCategory), v))
	})
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCategory), v))
	})
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCategory), v))
	})
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCategory), v))
	})
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCategory), v))
	})
}

// ProjectEQ applies the EQ predicate on the "project" field.
func ProjectEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProject), v))
	})
}

// ProjectNEQ applies the NEQ predicate on the "project" field.
func ProjectNEQ(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProject), v))
	})
}

// ProjectIn applies the In predicate on the "project" field.
func ProjectIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProject), v...))
	})
}

// ProjectNotIn applies the NotIn predicate on the "project" field.
func ProjectNotIn(vs ...string) predicate.Timestamp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Timestamp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProject), v...))
	})
}

// ProjectGT applies the GT predicate on the "project" field.
func ProjectGT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProject), v))
	})
}

// ProjectGTE applies the GTE predicate on the "project" field.
func ProjectGTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProject), v))
	})
}

// ProjectLT applies the LT predicate on the "project" field.
func ProjectLT(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProject), v))
	})
}

// ProjectLTE applies the LTE predicate on the "project" field.
func ProjectLTE(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProject), v))
	})
}

// ProjectContains applies the Contains predicate on the "project" field.
func ProjectContains(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProject), v))
	})
}

// ProjectHasPrefix applies the HasPrefix predicate on the "project" field.
func ProjectHasPrefix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProject), v))
	})
}

// ProjectHasSuffix applies the HasSuffix predicate on the "project" field.
func ProjectHasSuffix(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProject), v))
	})
}

// ProjectIsNil applies the IsNil predicate on the "project" field.
func ProjectIsNil() predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProject)))
	})
}

// ProjectNotNil applies the NotNil predicate on the "project" field.
func ProjectNotNil() predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProject)))
	})
}

// ProjectEqualFold applies the EqualFold predicate on the "project" field.
func ProjectEqualFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProject), v))
	})
}

// ProjectContainsFold applies the ContainsFold predicate on the "project" field.
func ProjectContainsFold(v string) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProject), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
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
func Not(p predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		p(s.Not())
	})
}