package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dornascarol/api-go-gin/controllers"
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
}
