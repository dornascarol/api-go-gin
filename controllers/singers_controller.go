package controllers

import (
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetSingers(c *gin.Context) {
	var singers []models.Singer

	database.DB.Find(&singers)
	c.JSON(200, singers)
}
