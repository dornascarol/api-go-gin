package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
