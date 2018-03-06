package goservice

import "github.com/go-martini/martini"
import "github.com/martini-contrib/render"

import "./features"

// Startup opens a server into the given port
func Startup() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/status", func() (int, string) { return 200, "ONLINE" })
	m.Group("/beer", (features.HandleBeers))
	m.Run()
}
