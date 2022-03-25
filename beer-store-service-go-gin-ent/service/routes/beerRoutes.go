package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	c.String(http.StatusOK, "AAA")
}

func Find(c *gin.Context) {
	c.String(http.StatusOK, "AAA")
}
