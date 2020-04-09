package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePageHandlder(c *gin.Context) {
	c.String(http.StatusOK, "Skrawki Puszczy discord bot")
}
