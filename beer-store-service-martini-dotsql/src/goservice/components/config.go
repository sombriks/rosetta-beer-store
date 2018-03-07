package components

import (
	"database/sql"
	"log"

	"github.com/gchaincl/dotsql"
	// Loading sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Db is defined at package level
var Db *sql.DB

// Dot is defined at package leve
var Dot *dotsql.DotSql

// init initializes the module
func init() {
	log.Println("openning database")
	// Get a database handle
	db, err := sql.Open("sqlite3", "beerstore.sqlite3")
	if err != nil {
		panic(err)
	}
	Db = db

	log.Println("opening scripts")
	// Loads queries from file
	dot, err := dotsql.LoadFromFile("queries.sql")
	if err != nil {
		panic(err)
	}
	Dot = dot

	log.Println("media table")
	// must create the database schema
	_, err = dot.Exec(db, "create-media-table")
	if err != nil {
		panic(err)
	}

	log.Println("beer table")
	_, err = dot.Exec(db, "create-beer-table")
	if err != nil {
		panic(err)
	}

	log.Println("count beers")
	rows, err := dot.Query(db, "count-beers")
	if err != nil {
		panic(err)
	}

	var count int
	if rows.Next() {
		rows.Scan(&count)
		// fmt.Printf("count: %d\n", count)
	}

	// defer rows.Close() would wail untill the surrounding
	// function gets finished...
	rows.Close()

	// ... but we need to seed the database
	if count == 0 {
		log.Println("insert beers")
		_, err := dot.Exec(db, "insert-a-few-beers")
		if err != nil {
			panic(err)
		}
	}
}
