package components

import (
	"log"

	"github.com/jinzhu/gorm"
)

func iferr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// GetDb returns the database instance
func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "beerstore.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
