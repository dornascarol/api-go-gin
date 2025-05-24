package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func EditSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")
	database.DB.First(&singer, id)

	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateSingerData(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Save(&singer)
	c.JSON(http.StatusOK, singer)
}
