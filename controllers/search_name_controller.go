package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// SearchSingerByName godoc
// @Summary      Get a singer by name
// @Description  Retrieves a singer from the database by their artist name
// @Tags         singers
// @Produce      json
// @Param        name  path      string  true  "Artist name"
// @Success      200   {object}  models.Singer "Successful response with the singer data"
// @Failure      404   {object}  map[string]string "Error response if singer not found"
// @Router       /singers/name/{name} [get]
func SearchSingerByName(c *gin.Context) {
	var singer models.Singer
	name := c.Param("name")
	database.DB.Where(&models.Singer{ArtistName: name}).First(&singer)

	if singer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
