package controllers

import (
	"github.com/gin-gonic/gin"
)

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Okay, " + name + "?",
	})
}
