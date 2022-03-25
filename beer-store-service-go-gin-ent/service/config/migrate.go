package config

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"

	// registrando o driver sql
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

func panicMaybe(err error) {
	if err != nil {
		panic(err)
	}
}

func MigrateUp() {

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	db, err := sql.Open("sqlite3", "./beerstore.sqlite3")
	panicMaybe(err)

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	panicMaybe(err)

	fmt.Printf("Applied %d migrations!\n", n)

}
