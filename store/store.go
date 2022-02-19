package store

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
)

func New() *Store {
	s := &Store{}

	return s
}

// Store ...
type Store struct {
	DB *sql.DB

	Types map[string]Type
}

func (s *Store) Builder() *Query {
	return &Query{}
}

func (s *Store) Get(ctx context.Context, r Query) (interface{}, error) {
	return nil, nil
}

func (s *Store) Set(ctx context.Context, r Query) (interface{}, error) {
	return nil, nil
}

func InjectStore(s *Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = WithContext(ctx, s)

			*r = *r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// WithContext ...
func WithContext(ctx context.Context, s *Store) context.Context {
	return context.WithValue(ctx, keyStore{}, s)
}

// FromContext ...
func FromContext(ctx context.Context) *Store {
	if s, ok := ctx.Value(keyStore{}).(*Store); ok {
		return s
	}

	return &Store{}
}

// With ...
func With(ctx context.Context) *Query {
	s := FromContext(ctx)

	fields := graphql.CollectAllFields(ctx)

	cols := []string{}

	table := "unknown"

	fCtx := graphql.GetFieldContext(ctx)

	fmt.Printf("fCtx = %+v\n", fCtx)
	fmt.Printf("fCtx.Field.Name = %s\n", fCtx.Field.Name)
	fmt.Printf("fCtx.Field.Definition.Type = %s\n", fCtx.Field.Definition.Type)

	parentResolver := ""

	if fCtx.Parent != nil && fCtx.Parent.Field.Field != nil {
		parentResolver = fCtx.Parent.Field.Name
		fmt.Printf("fCtx.Parent.Field = %+v\n", fCtx.Parent.Field)
		fmt.Printf("fCtx.Parent.Field.Name = %s\n", fCtx.Parent.Field.Name)
	}

	if fCtx.Field.Name == "user" {
		table = "users"
	} else if fCtx.Field.Name == "articles" {
		table = "articles"
	}

	fmt.Printf("table is %s\n", table)

	for _, f := range fields {
		if table == "users" && s.Types["user"].HasScalar(f) {
			cols = append(cols, f)
		}

		if table == "articles" && s.Types["article"].HasScalar(f) {
			cols = append(cols, f)
		}
	}

	for i, c := range cols {
		switch c {
		case "createdAt":
			cols[i] = "created_at"
		}
	}

	qry := &Query{
		Table:          table,
		Columns:        cols,
		currResolver:   fCtx.Field.Name,
		parentResolver: parentResolver,
		db:             s.DB,
	}

	return qry
}

type keyStore struct{}
