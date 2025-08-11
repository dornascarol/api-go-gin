package controllers

import (
	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/gin-gonic/gin"
)

// Greeting godoc
// @Summary      Get personalized greeting
// @Description  Returns a personalized greeting message
// @Tags         name
// @Produce      json
// @Param        name  path      string  true  "Name for the greeting"
// @Success      200   {object}  models.GreetingResponse "Successful response with greeting"
// @Failure      400   {object}  map[string]string "Error response for invalid parameters"
// @Failure      500   {object}  map[string]string "Internal server error"
// @Router       /{name} [get]
type GreetingController struct{}

func NewGreetingController() *GreetingController {
    return &GreetingController{}
}

func (gc *GreetingController) Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, entities.GreetingResponse{
		Message: "Okay, " + name + "?",
	})
}
