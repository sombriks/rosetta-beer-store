package components

import (

	// registrando o driver sql
	_ "github.com/mattn/go-sqlite3"
	"github.com/DavidHuie/gomigrate"
	"github.com/jinzhu/gorm"
	"database/sql"
)

type result struct {
	created int
}

func Setup() {
	db1,_ := sql.Open("sqlite3","beerstore.sqlite3")
	migrator, _ := gomigrate.NewMigrator(db1, gomigrate.Sqlite3{}, "./migrations")
	migrator.Migrate()
	db1.Close()
}

// GetDb returns the database instance
func GetDb() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("sqlite3", "beerstore.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
