package controllers

import (
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
