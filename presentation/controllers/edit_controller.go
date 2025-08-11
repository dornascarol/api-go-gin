package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/gin-gonic/gin"
)

// EditSinger godoc
// @Summary      Edit a singer
// @Description  Updates an existing singer in the database by ID
// @Tags         singers
// @Accept       json
// @Produce      json
// @Param        id     path      int           true  "Singer ID"
// @Param        singer body      models.Singer true  "Singer data to update"
// @Success      200    {object}  models.Singer "Successful response with the updated singer data"
// @Failure      400    {object}  map[string]string "Error response with validation message"
// @Failure      404    {object}  map[string]string "Error response if singer not found"
// @Router       /singers/{id} [patch]
type EditSingerController struct {
	Usecase *usecases.SingersUseCase
}

func NewEditSingerController(uc *usecases.SingersUseCase) *EditSingerController {
	return &EditSingerController{
		Usecase: uc,
	}
}

func (cc *EditSingerController) EditSinger(c *gin.Context) {
	id := c.Param("id")

	var updated entities.Singer
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entities.ValidateSingerData(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	singer, err := cc.Usecase.UpdateSinger(c.Request.Context(), id, &updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
