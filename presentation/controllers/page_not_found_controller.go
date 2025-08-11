package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotFoundController struct{}

func NewNotFoundController() *NotFoundController {
	return &NotFoundController{}
}

func (nfc *NotFoundController) RouteNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
