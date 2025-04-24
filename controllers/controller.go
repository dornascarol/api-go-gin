package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetSingers(c *gin.Context) {
	c.JSON(200, models.Singers)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Okay, " + name + "?",
	})
}

func CreateNewSinger(c *gin.Context) {
	var singer models.Singer
	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&singer)
	c.JSON(http.StatusOK, singer)
}
