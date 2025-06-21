package controllers

import (
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// GetSingers godoc
//
// @Summary      Retrieve all singers
// @Description  Get a list of all singers from the database
// @Tags         singers
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Singer
// @Router       /singers [get]
func GetSingers(c *gin.Context) {
	var singers []models.Singer

	database.DB.Find(&singers)
	c.JSON(200, singers)
}
