package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/gin-gonic/gin"
)

// SearchSingerByName godoc
// @Summary      Get a singer by name
// @Description  Retrieves a singer from the database by their artist name
// @Tags         singers
// @Produce      json
// @Param        name  path      string  true  "Artist name"
// @Success      200   {object}  models.Singer "Successful response with the singer data"
// @Failure      404   {object}  map[string]string "Error response if singer not found"
// @Router       /singers/name/{name} [get]
type SearchSingerByNameController struct {
	Usecase *usecases.SingersUseCase
}

func NewSearchSingerByNameController(uc *usecases.SingersUseCase) *SearchSingerByNameController {
	return &SearchSingerByNameController{
		Usecase: uc,
	}
}

func (cc *SearchSingerByNameController) SearchSingerByName(c *gin.Context) {
	name := c.Param("name")

	singer, err := cc.Usecase.GetSingerByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
