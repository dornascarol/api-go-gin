package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/dornascarol/api-go-gin/controllers"
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateSingerMock() {
	singer := models.Singer{ArtistName: "Test Singer", SongName: "Test Music", MusicalGenre: "Test Pagode"}
	database.DB.Create(&singer)
	ID = int(singer.ID)
}

func DeleteSingerMock() {
	var singer models.Singer
	database.DB.Delete(&singer, ID)
}

func TestFailed(t *testing.T) {
	t.Fatalf("Test failed intentionally!")
}

func TestGreetingStatusCode(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/Dornas", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Status should be equal")

	mockResponse := `{"API says:":"Okay, Dornas?"}`
	responseBody, _ := io.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}

func TestAllSingersHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateSingerMock()
	defer DeleteSingerMock()

	r := SetupTestRoutes()
	r.GET("/singers", controllers.GetSingers)
	req, _ := http.NewRequest("GET", "/singers", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchSingerByNameHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateSingerMock()
	defer DeleteSingerMock()

	r := SetupTestRoutes()
	r.GET("/singers/name/:name", controllers.SearchSingerByName)
	req, _ := http.NewRequest("GET", "/singers/name/Exaltasamba", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchSingerByIdHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateSingerMock()
	defer DeleteSingerMock()

	r := SetupTestRoutes()
	r.GET("/singers/:id", controllers.SearchSingerById)
	pathOfSearch := "/singers/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var mockSingerName models.Singer
	json.Unmarshal(response.Body.Bytes(), &mockSingerName)

	assert.Equal(t, "Test Singer", mockSingerName.ArtistName, "Singer names should be equal")
	assert.Equal(t, "Test Music", mockSingerName.SongName)
	assert.Equal(t, "Test Pagode", mockSingerName.MusicalGenre)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteSingerHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateSingerMock()

	r := SetupTestRoutes()
	r.DELETE("/singers/:id", controllers.DeleteSinger)
	pathOfSearch := "/singers/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathOfSearch, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
