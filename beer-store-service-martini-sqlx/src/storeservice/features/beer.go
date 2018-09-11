package features

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"

	// registrando o driver sql
	_ "github.com/mattn/go-sqlite3"
)

// Beer is our struct to represent database content
type Beer struct {
	gorm.Model
	Idbeer           int        `json:"idbeer,omitempty" gorm:"column:idbeer"`
	Titlebeer        string     `json:"titlebeer,omitempty" gorm:"column:titlebeer"`
	Descriptionbeer  string     `json:"descriptionbeer,omitempty" gorm:"column:descriptionbeer"`
	Creationdatebeer *time.Time `json:"creationdatebeer,omitempty" gorm:"column:creationdatebeer"`
	Idmedia          *int       `json:"idmedia,omitempty" gorm:"column:idmedia"`
}

// HandleBeers installs http handlers on "/beer"
func HandleBeers(r martini.Router) {

	r.Get("/list", func(req *http.Request, res render.Render) {

		db, err := gorm.Open("sqlite3", "beerstore.sqlite3")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&Beer{})

		defer db.Close()
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

		p, _ := strconv.Atoi(page)
		s, _ := strconv.Atoi(pageSize)

		beers := []Beer{}

		db.Where("titlebeer like ?", search).Offset((p - 1) * s).Limit(s).Find(&beers)

		res.JSON(200, beers)

		// foi
	})

	r.Get("/:idbeer", func(params martini.Params, res render.Render) {
		// sql := "select * from beer where idbeer = ?"
		ret := Beer{}

		// err := components.Db.Get(&ret, sql, params["idbeer"])
		// if err != nil {
		// 	res.Error(404)
		// 	fmt.Println("Beer not found")
		// 	return
		// }
		res.JSON(200, ret)
	})
}
