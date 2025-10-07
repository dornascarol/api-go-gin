package controllers

import (
	"context"
	"net/http"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/gin-gonic/gin"
)

// GetSingers godoc
// @Summary      Get all singers
// @Description  Retrieves a list of all singers from the database
// @Tags         singers
// @Produce      json
// @Success      200  {array}   models.Singer "Successful response with a list of singers"
// @Failure      500  {object}  map[string]string "Error response for internal server error"
// @Router       /singers [get]

type GetAllSingersUsecase interface {
	GetAllSingers(ctx context.Context) ([]entities.Singer, error)
}

type SingerController struct {
	Usecase GetAllSingersUsecase
}

func NewSingerController(uc GetAllSingersUsecase) *SingerController {
	return &SingerController{
		Usecase: uc,
	}
}

func (sc *SingerController) GetSingers(c *gin.Context) {
	singers, err := sc.Usecase.GetAllSingers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching for singers"})
		return
	}

	c.JSON(http.StatusOK, singers)
}
