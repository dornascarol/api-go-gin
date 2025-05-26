package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Welcome",
	})
}
