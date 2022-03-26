package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/beer"

	ent "github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models"
)

func List(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	search := c.DefaultQuery("search", "")

	log.Printf("page: %d pageSize: %d search: %s", page, pageSize, search)

	client, _ := ent.Open("sqlite3", "./beerstore.sqlite3")
	defer client.Close()

	beers, _ := client.Beer.Query().Where(
		beer.Or(
			beer.TitlebeerContainsFold(search),
			beer.DescriptionbeerContainsFold(search),
		),
	).Offset((page - 1) * pageSize).Limit(pageSize).All(context.TODO())

	c.JSON(http.StatusOK, beers)
}

func Find(c *gin.Context) {

	idbeer, _ := strconv.Atoi(c.Param("idbeer"))

	client, _ := ent.Open("sqlite3", "./beerstore.sqlite3")
	defer client.Close()

	oneBeer, _ := client.Beer.Query().Where(beer.IDEQ(idbeer)).First(context.TODO())

	c.JSON(http.StatusOK, oneBeer)
}
