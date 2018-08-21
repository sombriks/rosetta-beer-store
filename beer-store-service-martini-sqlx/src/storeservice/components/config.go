package components

import (
	"log"

	"io/ioutil"

	"github.com/jmoiron/sqlx"
	// Loading sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Db visible db connection
var Db *sqlx.DB

// init initializes the module
func init() {
	log.Println("openning database")
	// Get a database handle
	db, err := sqlx.Connect("sqlite3", "beerstore.sqlite3")
	if err != nil {
		log.Fatalln(err)
	}
	q, err := ioutil.ReadFile("./queries/create-table-media.sql")
	if err != nil {
		panic(err)

	}
	db.MustExec(string(q))
	q, err = ioutil.ReadFile("./queries/create-table-beer.sql")
	db.MustExec(string(q))
	q, err = ioutil.ReadFile("./queries/count-beers.sql")

	r, err := db.Queryx(string(q))
	count := 0
	if r.Next() {
		r.Scan(&count)
	}
	r.Close()
	if count == 0 {
		q, err = ioutil.ReadFile("./queries/insert-a-few-beers.sql")
		db.MustExec(string(q))
	}
	Db = db
}
