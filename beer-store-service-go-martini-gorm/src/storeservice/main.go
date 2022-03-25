package storeservice

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"

	"storeservice/components"
	"storeservice/routes"
)

// Startup opens a server into the given port
func Startup() {

	components.Setup()

	m := martini.Classic()

	m.Use(render.Renderer())

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "x-api-key", "x-secret-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	m.Use(martini.Static("static"))

	m.Get("/status", func() (int, string) { return 200, "ONLINE" })

	m.Group("/beer", (routes.HandleBeers))

	m.Run()
}
