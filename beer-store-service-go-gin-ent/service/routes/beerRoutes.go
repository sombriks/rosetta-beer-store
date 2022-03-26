package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/config"

	ent "github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models"
)

func List(c *gin.Context) {
	client, err := ent.Open("sqlite3", "file:./beerstore.sqlite3")
	config.PanicMaybe(err)
	client.Beer.Query().Where()
	c.String(http.StatusOK, "AAA")
}

func Find(c *gin.Context) {
	c.String(http.StatusOK, "AAA")
}
