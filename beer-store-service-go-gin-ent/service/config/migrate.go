package config

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"

	// registrando o driver sql
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
)

func MigrateUp() {

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	db, err := sql.Open("sqlite3", "./beerstore.sqlite3")
	PanicMaybe(err)

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	PanicMaybe(err)

	fmt.Printf("Applied %d migrations!\n", n)

}
