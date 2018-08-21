package features

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"storeservice/components"
)

// Beer is our struct to represent database content
type Beer struct {
	Idbeer           int        `json:"idbeer,omitempty" db:"idbeer"`
	Titlebeer        string     `json:"titlebeer,omitempty" db:"titlebeer"`
	Descriptionbeer  string     `json:"descriptionbeer,omitempty" db:"descriptionbeer"`
	Creationdatebeer *time.Time `json:"creationdatebeer,omitempty" db:"creationdatebeer"`
	Idmedia          *int       `json:"idmedia,omitempty" db:"idmedia"`
}

// HandleBeers installs http handlers on "/beer"
func HandleBeers(r martini.Router) {

	r.Get("/list", func(req *http.Request, res render.Render) {

		q := req.URL.Query()

		p := q["page"]
		s := q["pageSize"]
		param := q["search"]

		search := ""
		if len(param) != 0 {
			search = param[0]
		}
		search = "%" + search + "%"

		page := "1"
		if len(p) != 0 {
			page = p[0]
		}

		pageSize := "10"
		if len(s) != 0 {
			pageSize = s[0]
		}

		n1, err := strconv.Atoi(pageSize)
		n2, err := strconv.Atoi(page)

		skip := n1 * (n2 - 1)

		ret := []Beer{}
		sql := "select * from beer where titlebeer like ? or descriptionbeer like ? limit ? offset ? "
		err = components.Db.Select(&ret, sql, search, search, pageSize, skip)
		if err != nil {
			log.Fatalln(err)
		}
		res.JSON(200, ret)
		// foi
	})

	r.Get("/:idbeer", func(params martini.Params, res render.Render) {
		sql := "select * from beer where idbeer = ?"
		ret := Beer{}
		log.Println(params["idbeer"])
		err := components.Db.Get(&ret, sql, params["idbeer"])
		if err != nil {
			log.Fatalln(err)
		}
		res.JSON(200, ret)
	})
}