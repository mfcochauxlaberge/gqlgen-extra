package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mfcochauxlaberge/gqlgen-extra/demo"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/internal/fakedb"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

func main() {
	db, err := fakedb.New("fakedb")
	if err != nil {
		exitWithError(err)
		os.Exit(1)
	}

	store := &store.Store{
		DB: db,
	}

	server := demo.Server{
		DB:    db,
		Store: store,
		Port:  8080,
	}

	done := make(chan bool)

	go func() {
		err := server.Run()
		if err != nil && err != http.ErrServerClosed {
			exitWithError(err)
		}

		log.Printf("Server closed")
		done <- true
	}()

	log.Printf("Server now listening of port %d...", server.Port)

	<-done
	log.Printf("Exiting...")
	os.Exit(0)
}

func exitWithError(err error) {
	log.Fatalf("error: %s\n", err)
}
