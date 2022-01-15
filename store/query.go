package store

func NewQuery() *Query {
	qry := &Query{}

	return qry
}

// Query ...
type Query struct {
	Table   string
	Columns []string
	Filters []Filter
	Order   []Order
}

func (q *Query) From(t string) *Query {
	q.Table = t
	return q
}

func (q *Query) Select(cols ...string) *Query {
	q.Columns = cols
	return q
}

func (q *Query) SQL() string {
	return ""
}

func (q *Query) Into(dst interface{}) error {
	return nil
}

type Condition string

const (
	Eq Condition = "="
)

type Filter struct {
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
