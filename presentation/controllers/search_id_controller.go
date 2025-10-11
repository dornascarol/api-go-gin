package controllers

import (
	"context"
	"net/http"

	"github.com/dornascarol/api-go-gin/domain/entities"
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

type SearchSingerByIDUsecase interface {
	GetSingerByID(ctx context.Context, id string) (*entities.Singer, error)
}

type SearchSingerByIdController struct {
	Usecase SearchSingerByIDUsecase
}

func NewSearchSingerByIdController(uc SearchSingerByIDUsecase) *SearchSingerByIdController {
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
