package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePageHandler(c *gin.Context) {
	c.String(http.StatusOK, "Skrawki Puszczy discord bot")
}
