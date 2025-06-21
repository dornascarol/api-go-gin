package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// DeleteSinger godoc
//
// @Summary      Delete a singer by ID
// @Description  Removes a singer from the database using their unique ID
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Singer ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /singers/{id} [delete]
func DeleteSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")

	database.DB.Delete(&singer, id)
	c.JSON(http.StatusOK, gin.H{
		"Data": "Singer deleted successfully"})
}
