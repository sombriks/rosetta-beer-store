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

func iferr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// init initializes the module
func init() {
	log.Println("openning database")
	// var Db *sqlx.DB
	var err error
	var q []byte
	Db, err = sqlx.Connect("sqlite3", "beerstore.sqlite3")
	iferr(err)
	q, err = ioutil.ReadFile("./queries/create-table-media.sql")
	iferr(err)
	Db.MustExec(string(q))
	q, err = ioutil.ReadFile("./queries/create-table-beer.sql")
	Db.MustExec(string(q))
	q, err = ioutil.ReadFile("./queries/count-beers.sql")

	r, err := Db.Queryx(string(q))
	count := 0
	if r.Next() {
		r.Scan(&count)
	}
	r.Close()
	if count == 0 {

		q, err = ioutil.ReadFile("./queries/insert-a-few-beers.sql")
		Db.MustExec(string(q))
	}
}
