package store

import (
	"context"
	"database/sql"
	"net/http"
)

func New() *Store {
	s := &Store{}

	return s
}

// Store ...
type Store struct {
	Conn *sql.DB
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

// Builder ...
func Builder(ctx context.Context) *Query {
	// if s, ok := ctx.Value(keyStore{}).(*Query); ok {
	// 	return s
	// }

	return &Query{}
}

type keyStore struct{}
