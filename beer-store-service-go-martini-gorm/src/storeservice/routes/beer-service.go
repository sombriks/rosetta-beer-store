package routes

import (
	"net/http"
	"strconv"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	components "storeservice/components"

	"storeservice/model"
)

// HandleBeers installs http handlers on "/beer"
func HandleBeers(r martini.Router) {

	r.Get("/list", func(req *http.Request, res render.Render) {
		db := components.GetDb()
		defer db.Close()
		q := req.URL.Query()
		search := q.Get("search")
		if len(search) != 0 {
			search = "%" + search + "%"
		} else {
			search = "%%"
		}

		page := q.Get("page")
		if len(page) == 0 {
			page = "1"
		}
		pageSize := q.Get("pageSize")
		if len(pageSize) == 0 {
			pageSize = "10"
		}
		p, _ := strconv.Atoi(page)
		s, _ := strconv.Atoi(pageSize)
		beers := []model.Beer{}
		db.Table("beer").Where("titlebeer like ?", search).Offset((p - 1) * s).Limit(s).Find(&beers)
		res.JSON(200, beers)
	})

	r.Get("/:idbeer", func(params martini.Params, res render.Render) {
		idbeer, _ := strconv.Atoi(params["idbeer"])
		db := components.GetDb()
		defer db.Close()
		beer := model.Beer{}
		db.First(&beer, idbeer)
		res.JSON(200, beer)
	})
}
