package store

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/model"
)

func NewQuery() *Query {
	qry := &Query{}

	return qry
}

// Query ...
type Query struct {
	Table   string
	Columns []string
	Filters []FilterRule
	Order   []Order

	currResolver   string
	parentResolver string

	db   *sql.DB
	into interface{}
}

func (q *Query) From(t string) *Query {
	q.Table = t
	return q
}

func (q *Query) Select(cols ...string) *Query {
	q.Columns = cols
	return q
}

func (q *Query) Where(f string) *Query {
	filter := FilterRule{
		Column:    "author",
		Condition: Eq,
		Value:     "u1",
	}

	q.Filters = append(q.Filters, filter)

	return q
}

func (q *Query) OrderBy(f string) *Query {
	order := Order{
		Column:    "created_at",
		Direction: Descending,
	}

	q.Order = append(q.Order, order)

	return q
}

func (q *Query) GetMany(opts ...OptionFunc) error {
	var err error

	for _, opt := range opts {
		q, err = opt(q)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *Query) GetOne(opts ...OptionFunc) error {
	var err error

	for _, opt := range opts {
		q, err = opt(q)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *Query) SQL() string {
	table := `"` + q.Table + `"`

	cols := make([]string, 0, len(q.Columns))
	for _, c := range q.Columns {
		cols = append(cols, `"`+c+`"`)
	}

	stmt := `SELECT ` + strings.Join(cols, ", ")
	stmt += ` FROM ` + table

	if len(q.Filters) > 0 {
		wheres := []string{}

		for _, o := range q.Filters {
			col := `"` + o.Column + `"`
			wheres = append(wheres, fmt.Sprintf(
				"%s %s '%s'",
				col, o.Condition, o.Value,
			))
		}

		stmt += " WHERE "
		stmt += strings.Join(wheres, " AND ")
	}

	return stmt
}

func (q *Query) Into(dst interface{}) error {
	val := reflect.ValueOf(dst)
	eval := val.Elem()

	fmt.Printf("kind: %s\n", eval.Kind())

	if eval.Kind() == reflect.Slice {
		stmt := q.SQL()

		fmt.Printf("stmt: %s\n", stmt)

		rows, err := q.db.Query(stmt)
		if err != nil {
			return err
		}

		t := eval.Type().Elem()
		// el := reflect.New(t).Elem()
		els := reflect.New(reflect.SliceOf(t)).Elem()

		fmt.Printf("els kind: %s\n", els.Kind())

		for rows.Next() {
			vals := (model.Article{}).ValsFromCols(nil)

			err := rows.Scan(vals...)
			if err != nil {
				return err
			}

			a := (model.Article{}).ScanInto(vals)

			// for i:=0; i<val.NumField(); i++ {
			// 	f:= val.FieldByIndex(i)
			// 	f.

			// 	if val.
			// }

			// f := el.FieldByName("ID")
			// if f.CanSet() {
			// 	f.SetString(vals[0].(string))
			// }

			// f = el.FieldByName("CreatedAt")
			// if f.CanSet() {
			// 	f.Set(reflect.ValueOf(vals[1].(time.Time)))
			// }

			// f = el.FieldByName("Title")
			// if f.CanSet() {
			// 	f.SetString(vals[2].(string))
			// }

			// f = el.FieldByName("Content")
			// if f.CanSet() {
			// 	f.SetString(vals[3].(string))
			// }

			el := reflect.ValueOf(a)
			els = reflect.Append(els, el)
		}

		fmt.Printf("can set val: %t\n", val.CanSet())
		fmt.Printf("can set eval: %t\n", eval.CanSet())

		eval.Set(els)

		return nil
	}

	stmt := q.SQL()

	fmt.Printf("stmt: %s\n", stmt)

	row := q.db.QueryRow(stmt)

	vals := []interface{}{
		"",
		"",
	}

	err := row.Scan(
		&vals[0],
		&vals[1],
	)
	if err != nil {
		return err
	}

	// for i:=0; i<val.NumField(); i++ {
	// 	f:= val.FieldByIndex(i)
	// 	f.

	// 	if val.
	// }

	f := eval.FieldByName("ID")
	if f.CanSet() {
		f.SetString(vals[0].(string))
	}

	f = eval.FieldByName("Username")
	if f.CanSet() {
		f.SetString(vals[1].(string))
	}

	return nil
}

type Condition string

const (
	Eq Condition = "="
)

type FilterRule struct {
	Column    string
	Condition Condition
	Value     interface{}
}

type Direction int

const (
	Ascending Direction = iota
	Descending
)

type Order struct {
	Column    string
	Direction Direction
}
