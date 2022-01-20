package demo

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/gen"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"

	// embed package
	_ "embed"
)

//go:embed seed.sql
var seed string

// Server ...
type Server struct {
	Port  uint
	DB    *sql.DB
	Store *store.Store
}

func (s *Server) Run() error {
	r := chi.NewRouter()

	// Initialize store
	s.Store.Types = map[string]store.Type{
		"article": {
			Scalars: []string{"id", "createdAt", "title", "content"},
		},
		"user": {
			Scalars: []string{"id", "username"},
		},
	}

	// Seed
	_, err := s.DB.Exec(seed)
	if err != nil {
		return err
	}

	r.Use(store.InjectStore(s.Store))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("heyho!\n")
			next.ServeHTTP(w, r)
		})
	})

	r.Handle("/query", handler.NewDefaultServer(
		gen.NewExecutableSchema(
			gen.Config{Resolvers: &graph.Resolver{}},
		),
	))

	r.Handle("/", playground.Handler("Playground", "/query"))

	port := s.Port
	if s.Port == 0 {
		port = 8080
	}

	err = http.ListenAndServe(":"+strconv.FormatUint(uint64(port), 10), r)
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
