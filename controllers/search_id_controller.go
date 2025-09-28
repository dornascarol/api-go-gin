package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// SearchSingerById godoc
// @Summary      Get a singer by ID
// @Description  Retrieves a singer from the database by their ID
// @Tags         singers
// @Produce      json
// @Param        id   path      string  true  "Singer ID"
// @Success      200  {object}  models.Singer "Successful response with the singer data"
// @Failure      404  {object}  map[string]string "Error response if singer not found"
// @Router       /singers/{id} [get]
func SearchSingerById(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")
	database.DB.First(&singer, id)

	if singer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
