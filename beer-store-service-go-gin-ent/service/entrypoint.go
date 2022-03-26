package service

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/config"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/routes"
)

func Start() {

	// run migrations
	config.MigrateUp()

	// configure middleware
	r := gin.Default()
	r.GET("/beer/list", routes.List)
	r.GET("/beer/:idbeer", routes.Find)

	// cors
	r.Use(cors.Default())

	// start service
	r.Run(":3000")
}
