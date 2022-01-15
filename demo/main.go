package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/graph/gen"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/internal/fakedb"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func main() {
	db, err := fakedb.New("fakedb")
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	log.Printf("Creating demo schema...")

	_, err = db.Exec(`
		CREATE TABLE "users" (
			"id" TEXT PRIMARY KEY,
			"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
			"username" TEXT NOT NULL
		);

		CREATE TABLE "articles" (
			"id" TEXT PRIMARY KEY,
			"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
			"title" TEXT NOT NULL,
			"content" TEXT NOT NULL,
			"author" TEXT NOT NULL,
			FOREIGN KEY ("author") REFERENCES "users" ("id")
		);
	
		CREATE TABLE "comments" (
			"id" TEXT PRIMARY KEY,
			"created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
			"content" TEXT NOT NULL,
			"article" TEXT NOT NULL,
			"author" TEXT NOT NULL,
			FOREIGN KEY ("article") REFERENCES "articles" ("id"),
			FOREIGN KEY ("author") REFERENCES "users" ("id")
		);

		CREATE TABLE "articles_tags" (
			"label" TEXT NOT NULL,
			"article" TEXT NOT NULL,
			PRIMARY KEY ("label", "article"),
			FOREIGN KEY ("article") REFERENCES "articles" ("id")
		);

		CREATE TABLE "likes" (
			"user" TEXT NOT NULL,
			"article" TEXT NOT NULL,
			PRIMARY KEY ("user", "article"),
			FOREIGN KEY ("user") REFERENCES "users" ("id"),
			FOREIGN KEY ("article") REFERENCES "articles" ("id")
		);
	`)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	log.Printf("Inserting seed data...")

	_, err = db.Exec(`
		-- Users
		INSERT INTO "users" ("id", "created_at", "username")
			VALUES ('u1', NOW(), 'user1');
		INSERT INTO "users" ("id", "created_at", "username")
			VALUES ('u2', NOW(), 'user2');
		INSERT INTO "users" ("id", "created_at", "username")
			VALUES ('u3', NOW(), 'user3');

		-- Articles
		INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
			VALUES ('a1', NOW(), 'Article 1', 'This is the content of article 1.', 'u1');
		INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
			VALUES ('a2', NOW(), 'Article 2', 'This is the content of article 2.', 'u1');
		INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
			VALUES ('a3', NOW(), 'Article 3', 'This is the content of article 3.', 'u2');
		INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
			VALUES ('a4', NOW(), 'Article 4', 'This is the content of article 4.', 'u2');
		INSERT INTO "articles" ("id", "created_at", "title", "content", "author")
			VALUES ('a5', NOW(), 'Article 5', 'This is the content of article 5.', 'u3');

		-- Comments
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c1', NOW(), 'This is comment 1.', 'a1', 'u1');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c2', NOW(), 'This is comment 2.', 'a1', 'u1');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c3', NOW(), 'This is comment 3.', 'a2', 'u2');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c4', NOW(), 'This is comment 4.', 'a2', 'u2');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c5', NOW(), 'This is comment 5.', 'a3', 'u1');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c6', NOW(), 'This is comment 5.', 'a4', 'u3');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c7', NOW(), 'This is comment 5.', 'a4', 'u2');
		INSERT INTO "comments" ("id", "created_at", "content", "article", "author")
			VALUES ('c8', NOW(), 'This is comment 5.', 'a4', 'u1');

		-- Tags
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('tech', 'a1');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('tech', 'a2');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('sports', 'a1');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('sports', 'a3');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('tech', 'a4');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('health', 'a4');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('politics', 'a5');
		INSERT INTO "articles_tags" ("label", "article")
			VALUES ('health', 'a5');

		-- Likes
		INSERT INTO "likes" ("user", "article")
			VALUES ('u1', 'a1');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u2', 'a2');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u1', 'a4');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u2', 'a1');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u3', 'a4');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u3', 'a5');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u2', 'a3');
		INSERT INTO "likes" ("user", "article")
			VALUES ('u1', 'a5');
	`)
	if err != nil {
		printError(err)
		os.Exit(1)
	}

	s := &store.Store{
		Conn: db,
	}

	r := chi.NewRouter()

	r.Use(store.InjectStore(s))

	r.Handle("/query", handler.NewDefaultServer(
		gen.NewExecutableSchema(
			gen.Config{Resolvers: &graph.Resolver{}},
		),
	))

	r.Handle("/", playground.Handler("Playground", "/query"))

	log.Printf("Server is read to accept requests")

	err = http.ListenAndServe(":8181", r)
	if err != nil && err == http.ErrServerClosed {
		os.Exit(0)
	} else if err != nil {
		printError(err)
		os.Exit(1)
	}
}

func printError(err error) {
	fmt.Printf("error: %s\n", err)
}
