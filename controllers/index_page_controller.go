package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func DisplayIndexPage(c *gin.Context) {
	var singers []models.Singer
	database.DB.Find(&singers)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"singers": singers,
	})
}
