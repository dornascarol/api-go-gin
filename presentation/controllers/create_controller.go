package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/gin-gonic/gin"
)

// CreateNewSinger godoc
// @Summary      Create a new singer
// @Description  Adds a new singer to the database
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        singer  body      models.Singer  true  "Singer data"
// @Success      200     {object}  models.Singer "Successful response with the created singer data"
// @Failure      400     {object}  map[string]string "Error response with validation message"
// @Router       /singers [post]
type CreateSingerController struct {
	Usecase *usecases.SingersUseCase
}

func NewCreateSingerController(uc *usecases.SingersUseCase) *CreateSingerController {
	return &CreateSingerController{
		Usecase: uc,
	}
}

func (cc *CreateSingerController) CreateNewSinger(c *gin.Context) {
	var singer entities.Singer

	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := entities.ValidateSingerData(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	created, err := cc.Usecase.CreateSinger(c.Request.Context(), &singer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}
