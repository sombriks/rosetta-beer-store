package goservice

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"
import "github.com/martini-contrib/cors"

import "./features"

// Startup opens a server into the given port
func Startup() {
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

	m.Group("/beer", (features.HandleBeers))

	m.Run()
}
