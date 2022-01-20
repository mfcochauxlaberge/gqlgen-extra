package store

import (
	"log"

	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/model"
)

// OptionFunc ...
type OptionFunc func(q *Query) (*Query, error)

// Collection ...
func Collection(t string) OptionFunc {
	return func(q *Query) (*Query, error) {
		q.Table = t
		return q, nil
	}
}

// Attributes ...
func Attributes(cols ...string) OptionFunc {
	return func(q *Query) (*Query, error) {
		q.Columns = cols
		return q, nil
	}
}

// Filter ...
func Filter(c string, comp Condition, v interface{}) OptionFunc {
	return func(q *Query) (*Query, error) {
		f := FilterRule{
			Column:    c,
			Condition: comp,
			Value:     v,
		}

		q.Filters = append(q.Filters, f)

		return q, nil
	}
}

// OrderBy ...
func OrderBy(c string, comp Condition, v interface{}) OptionFunc {
	return func(q *Query) (*Query, error) {
		f := FilterRule{
			Column:    c,
			Condition: comp,
			Value:     v,
		}

		q.Filters = append(q.Filters, f)

		return q, nil
	}
}

// Into ...
func Into(v interface{}) OptionFunc {
	return func(q *Query) (*Query, error) {
		err := q.Into(v)
		if err != nil {
			return nil, err
		}

		return q, nil
	}
}

// ByID ...
func ByID(id string) OptionFunc {
	return Filter("id", Eq, id)
}

// FromRelationship ...
func FromRelationship(v interface{}) OptionFunc {
	id := ""

	switch t := v.(type) {
	case *model.User:
		id = t.ID
	case *model.Article:
		id = t.ID
	}

	log.Printf("v: %+v\n", v)

	return func(q *Query) (*Query, error) {
		f := FilterRule{
			Column:    "author",
			Condition: Eq,
			Value:     id,
		}

		q.Filters = append(q.Filters, f)

		return q, nil
	}
}
