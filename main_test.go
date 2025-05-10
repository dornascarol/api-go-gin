package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestFailed(t *testing.T) {
	t.Fatalf("Test failed intentionally!")
}
