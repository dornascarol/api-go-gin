package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func SearchSingerById(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")
	database.DB.First(&singer, id)

	if singer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
