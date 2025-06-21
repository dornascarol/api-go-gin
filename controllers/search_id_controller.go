package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// SearchSingerById godoc
//
// @Summary		Search singer by Id
// @Description	Get details of a singer using their unique ID
// @Tags			singers
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Singer ID"
// @Success		200	{object}	models.Singer
// @Failure		404	{object}	map[string]string
// @Router			/singers/{id} [get]
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
