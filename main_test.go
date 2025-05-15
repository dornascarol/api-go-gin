package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dornascarol/api-go-gin/controllers"
	"github.com/dornascarol/api-go-gin/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
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
	r := SetupTestRoutes()
	r.GET("/singers", controllers.GetSingers)
	req, _ := http.NewRequest("GET", "/singers", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
