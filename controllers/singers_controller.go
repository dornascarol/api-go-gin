package controllers

import (
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// GetSingers godoc
// @Summary      Get all singers
// @Description  Retrieves a list of all singers from the database
// @Tags         singers
// @Produce      json
// @Success      200  {array}   models.Singer "Successful response with a list of singers"
// @Failure      500  {object}  map[string]string "Error response for internal server error"
// @Router       /singers [get]
func GetSingers(c *gin.Context) {
	var singers []models.Singer

	database.DB.Find(&singers)
	c.JSON(200, singers)
}
