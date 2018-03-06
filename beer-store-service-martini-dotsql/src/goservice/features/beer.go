package features

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// Beer is our struct to represent database content
type Beer struct {
	Idbeer    int    `json:"idbeer,omitempty"`
	Titlebeer string `json:"titlebeer,omitempty"`
}

// HandleBeers installs http handlers on "/beer"
func HandleBeers(r martini.Router) {

	r.Get("/list", func(req *http.Request, res render.Render) {
		var ret [2]Beer
		ret[0] = Beer{Idbeer: 9, Titlebeer: "skol"}
		ret[1] = Beer{Idbeer: 1, Titlebeer: "Brahma"}
		res.JSON(200, ret)
	})
}
