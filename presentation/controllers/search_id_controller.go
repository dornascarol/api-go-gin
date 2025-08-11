package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/gin-gonic/gin"
)

// SearchSingerById godoc
// @Summary      Get a singer by ID
// @Description  Retrieves a singer from the database by their ID
// @Tags         singers
// @Produce      json
// @Param        id   path      int  true  "Singer ID"
// @Success      200  {object}  models.Singer "Successful response with the singer data"
// @Failure      404  {object}  map[string]string "Error response if singer not found"
// @Router       /singers/{id} [get]
type SearchSingerByIdController struct {
	Usecase *usecases.SingersUseCase
}

func NewSearchSingerByIdController(uc *usecases.SingersUseCase) *SearchSingerByIdController {
	return &SearchSingerByIdController{
		Usecase: uc,
	}
}

func (cc *SearchSingerByIdController) SearchSingerById(c *gin.Context) {
	id := c.Param("id")

	singer, err := cc.Usecase.GetSingerByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
