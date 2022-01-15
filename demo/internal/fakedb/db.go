package fakedb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/dchest/uniuri"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	running     bool
	runningLock sync.Mutex
	ip, port    string
)

// New spawns the PostgreSQL container if not already running, creates a new
// database, and returns a connection to it.
//
// The name of the dabase is prefix_rand, where prefix is the value given as an
// argument and rand is a random alphanumeric string of 8 characters.
//
// It is the user's responsibility to close the connection.
func New(prefix string) (*sql.DB, error) {
	runningLock.Lock()
	defer runningLock.Unlock()

	if !running {
		// Instantiate container
		err := instantiateContainer()
		if err != nil {
			return nil, err
		}

		running = true
	}

	// Connect to PostgreSQL instance
	url := fmt.Sprintf(
		"postgresql://postgres:postgres@%s:%s?sslmode=disable",
		ip,
		port,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// Create database
	name := prefix + "_" + uniuri.NewLen(8)

	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, name))
	if err != nil {
		return nil, err
	}

	// Connect to database
	url = fmt.Sprintf(
		"postgresql://postgres:postgres@%s:%s/%s?sslmode=disable",
		ip, port, name,
	)

	log.Printf("Database connection: %s", url)

	err = db.Close()
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func instantiateContainer() error {
	name := "temporary_test_db_" + uniuri.NewLen(4)

	req := tc.ContainerRequest{
		Name:  name,
		Image: "postgres:12.3",
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432"),
		AutoRemove:   true,
	}

	ctx := context.Background()

	container, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	ip, err = container.Host(ctx)
	if err != nil {
		return err
	}

	natPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return err
	}

	port = natPort.Port()

	return nil
}
