package components

import (
	"database/sql"

	"github.com/gchaincl/dotsql"
	_ "rsc.io/sqlite"
)

func init() {
	// Get a database handle
	Db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		panic(err)
	}

	// Loads queries from file
	Dot, err := dotsql.LoadFromFile("queries.sql")
}
