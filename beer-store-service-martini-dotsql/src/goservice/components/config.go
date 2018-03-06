package components

import (
	"database/sql"
	"fmt"

	"github.com/gchaincl/dotsql"
	// Loading sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Db is defined at package level
var Db *sql.DB

// Dot is defined at package leve
var Dot *dotsql.DotSql

func init() {
	// Get a database handle
	db, err := sql.Open("sqlite3", "foo.sqlite3")
	if err != nil {
		panic(err)
	}
	Db = db

	// Loads queries from file
	dot, err := dotsql.LoadFromFile("queries.sql")
	if err != nil {
		panic(err)
	}
	Dot = dot
	res, err := dot.Exec(db, "create-beer-table")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
