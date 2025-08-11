package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/dornascarol/api-go-gin/domain/entities"
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
type DeleteSingerController struct {
	Usecase *usecases.SingersUseCase
}

func NewDeleteSingerController(uc *usecases.SingersUseCase) *DeleteSingerController {
	return &DeleteSingerController{
		Usecase: uc,
	}
}

func (cc *DeleteSingerController) DeleteSinger(c *gin.Context) {
	id := c.Param("id")

	err := cc.Usecase.DeleteSinger(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Singer not found or could not be deleted"})
		return
	}

	c.JSON(http.StatusOK, entities.DeleteResponse{
		Data: "Singer deleted successfully",
	})
}
