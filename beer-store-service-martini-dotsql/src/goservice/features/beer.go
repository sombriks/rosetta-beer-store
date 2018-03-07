package features

import (
	"fmt"
	"net/http"

	"../components"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// Beer is our struct to represent database content
type Beer struct {
	Idbeer          int    `json:"idbeer,omitempty"`
	Titlebeer       string `json:"titlebeer,omitempty"`
	Descriptionbeer string `json:"descriptionbeer,omitempty"`
}

// HandleBeers installs http handlers on "/beer"
func HandleBeers(r martini.Router) {

	r.Get("/list", func(req *http.Request, res render.Render) {

		fmt.Println("GET params were:", req.URL.Query()["a"])

		q := req.URL.Query()

		p := q["page"]
		s := q["pageSize"]
		param := q["textobusca"]

		textobusca := ""
		if len(param) != 0 {
			textobusca = param[0]
		}
		textobusca = "%" + textobusca + "%"

		page := "1"
		if len(p) != 0 {
			page = p[0]
		}

		pageSize := "10"
		if len(s) != 0 {
			pageSize = s[0]
		}

		// var ret [2]Beer
		// ret[0] = Beer{Idbeer: 9, Titlebeer: "skol"}
		// ret[1] = Beer{Idbeer: 1, Titlebeer: "Brahma"}
		rows, err := components.Dot.Query(components.Db, "find-beers", textobusca, textobusca, page, pageSize)
		if err != nil {
			panic(err)
		}

		ret := make([]Beer, 0)
		for rows.Next() {
			b := Beer{}
			rows.Scan(&b)
			ret = append(ret, b)
		}

		res.JSON(200, ret)
	})
}
