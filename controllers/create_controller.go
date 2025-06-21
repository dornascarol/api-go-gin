package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// CreateNewSinger godoc
//
// @Summary      Create a new singer
// @Description  Adds a new singer to the database
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        singer  body      models.Singer  true  "Singer data"
// @Success      200     {object}  models.Singer
// @Failure      400     {object}  map[string]string
// @Router       /singers [post]
func CreateNewSinger(c *gin.Context) {
	var singer models.Singer

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

	database.DB.Create(&singer)
	c.JSON(http.StatusOK, singer)
}
