package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// SearchSingerByName godoc
//
// @Summary      Retrieve a singer by name
// @Description  Get details of a singer using their artist name
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "Artist Name"
// @Success      200   {object}  models.Singer
// @Failure      404   {object}  map[string]string
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
