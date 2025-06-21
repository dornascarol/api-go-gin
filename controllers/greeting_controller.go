package controllers

import (
	"github.com/gin-gonic/gin"
)

// Greeting godoc
//
// @Summary		Greet a name
// @Description	Returns Returns a greeting message for the user with the provided name
// @Tags			name
// @Accept			json
// @Produce		json
// @Param			name	path		string	true	"Name"
// @Success		200	{object}	map[string]string
// @Router			/{name} [get]
func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Okay, " + name + "?",
	})
}
