package components

import (
	"log"
)

func iferr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// init initializes the module
func init() {

}
