package demo_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mfcochauxlaberge/gold"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/internal/fakedb"
	"github.com/mfcochauxlaberge/gqlgen-extra/demo/internal/scenarios"
	"github.com/mfcochauxlaberge/gqlgen-extra/gqltest"
	"github.com/mfcochauxlaberge/gqlgen-extra/store"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

func TestDemo(t *testing.T) {
	runner := gold.NewRunner("testdata")
	runner.Update = true
	runner.Filters = []gold.Filter{
		// // Replace `"id": "some_id"`
		// gold.CustomFilter(
		// 	regexp.MustCompile(`"id": "[a-zA-Z0-9_]{1,}"`), `"id": "__ID__"`,
		// ),
		// // Replace `id: "some_id"`
		// gold.CustomFilter(
		// 	regexp.MustCompile(`id: "[a-zA-Z0-9_]{1,}"`), `id: "__ID__"`,
		// ),
		// // Replace `Id: "some_id"`
		// gold.CustomFilter(
		// 	regexp.MustCompile(`Id: "[a-zA-Z0-9_]{1,}"`), `Id: "__ID__"`,
		// ),
	}

	for _, test := range scenarios.Scenarios {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)

			server := startServer()

			endpoint := fmt.Sprintf("http://localhost:%d/query", server.Port)

			env := &gqltest.Env{
				Endpoint: endpoint,
				Rec: &gqltest.Recorder{
					Addr: endpoint,
				},
			}

			test.Play(env)

			// Golden file
			filename := strings.ReplaceAll(test.Name, " ", "_") + ".txt"
			path := filepath.Join("scenarios", filename)

			err := runner.Test(path, env.Rec.Summary())
			if _, ok := err.(gold.ComparisonError); ok {
				assert.Fail("file is different", test.Name)
			} else if err != nil {
				panic(err)
			}

			// err = server.Conn.Close()
			// assert.NoError(err)
		})
	}
}

func startServer() *demo.Server {
	// Get connection to a temporary database.
	db, err := fakedb.New("fakedb")
	if err != nil {
		panic(err)
	}

	// Find free port
	port := uint(freeport.GetPort())

	server := demo.Server{
		Port: port,
		DB:   db,
		Store: &store.Store{
			DB: db,
		},
	}

	// Run server
	go func() {
		err := server.Run()
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(200 * time.Millisecond)

	return &server
}
