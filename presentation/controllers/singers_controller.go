package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
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
type SingerController struct {
	Usecase *usecases.SingersUseCase
}

func NewSingerController(uc *usecases.SingersUseCase) *SingerController {
	return &SingerController{
		Usecase: uc,
	}
}

func (sc *SingerController) GetSingers(c *gin.Context) {
	singers, err := sc.Usecase.GetAllSingers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cantores"})
		return
	}

	c.JSON(http.StatusOK, singers)
}
