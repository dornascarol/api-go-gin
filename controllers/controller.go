package controllers

import (
	"net/http"

	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetSingers(c *gin.Context) {
	var singers []models.Singer

	database.DB.Find(&singers)
	c.JSON(200, singers)
}

func CreateNewSinger(c *gin.Context) {
	var singer models.Singer

	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateSingerData(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&singer)
	c.JSON(http.StatusOK, singer)
}

func SearchSingerById(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")
	database.DB.First(&singer, id)

	if singer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}

func DeleteSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")

	database.DB.Delete(&singer, id)
	c.JSON(http.StatusOK, gin.H{
		"Data": "Singer deleted successfully"})
}

func EditSinger(c *gin.Context) {
	var singer models.Singer
	id := c.Params.ByName("id")
	database.DB.First(&singer, id)

	if err := c.ShouldBindJSON(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateSingerData(&singer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Save(&singer)
	c.JSON(http.StatusOK, singer)
}

func SearchSingerByName(c *gin.Context) {
	var singer models.Singer
	name := c.Param("name")
	database.DB.Where(&models.Singer{ArtistName: name}).First(&singer)

	if singer.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Singer not found"})
		return
	}

	c.JSON(http.StatusOK, singer)
}
