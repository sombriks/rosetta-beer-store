package features

import (
	"fmt"
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

		search := q.Get("search")
		if len(search) != 0 {
			search = "%" + search + "%"
		}

		page := q.Get("page")
		if len(page) == 0 {
			page = "1"
		}

		pageSize := q.Get("pageSize")
		if len(pageSize) == 0 {
			pageSize = "10"
		}

		n1, err := strconv.Atoi(pageSize)
		if err != nil {
			res.Text(500, pageSize+" is not a number")
			return
		}
		n2, err := strconv.Atoi(page)
		if err != nil {
			res.Text(500, page+" is not a number")
			return
		}

		skip := n1 * (n2 - 1)

		ret := []Beer{}
		sql := "select * from beer where titlebeer like ? or descriptionbeer like ? limit ? offset ? "
		err = components.Db.Select(&ret, sql, search, search, pageSize, skip)
		if err != nil {
			res.Text(500, "unable to select: "+err.Error())
			return
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
			res.Error(404)
			fmt.Println("Beer not found")
			return
		}
		res.JSON(200, ret)
	})
}
