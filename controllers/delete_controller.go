package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

// DeleteSinger godoc
// @Summary      Delete a singer
// @Description  Deletes a singer from the database by ID
// @Tags         singers
// @Produce      json
// @Param        id   path      int  true  "Singer ID"
// @Success      200  {object}  models.DeleteResponse "Successful response with deletion message"
// @Failure      404  {object}  map[string]string "Error response if singer not found"
// @Router       /singers/{id} [delete]
func DeleteSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")

	database.DB.Delete(&singer, id)
	c.JSON(http.StatusOK, models.DeleteResponse{
		Data: "Singer deleted successfully",
	})
}
