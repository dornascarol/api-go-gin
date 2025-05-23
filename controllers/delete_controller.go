package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func DeleteSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")

	database.DB.Delete(&singer, id)
	c.JSON(http.StatusOK, gin.H{
		"Data": "Singer deleted successfully"})
}
