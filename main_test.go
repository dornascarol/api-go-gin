package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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

	fmt.Println(string(responseBody))
	fmt.Println(mockResponse)
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
