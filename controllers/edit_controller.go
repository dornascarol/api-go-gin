package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// EditSinger godoc
// @Summary      Edit a singer
// @Description  Updates an existing singer in the database by ID
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "Singer ID"
// @Param        singer body      models.Singer true  "Singer data to update"
// @Success      200    {object}  models.Singer "Successful response with the updated singer data"
// @Failure      400    {object}  map[string]string "Error response with validation message"
// @Failure      404    {object}  map[string]string "Error response if singer not found"
// @Router       /singers/{id} [patch]
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
