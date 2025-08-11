package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/dornascarol/api-go-gin/infrastructure/persistence/postgres"
	"github.com/gin-gonic/gin"
)
type IndexPageController struct{}

func NewIndexPageController() *IndexPageController {
    return &IndexPageController{}
}

func (ipc *IndexPageController) DisplayIndexPage(c *gin.Context) {
	var singers []entities.Singer
	postgres.DB.Find(&singers)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"singers": singers,
	})
}
