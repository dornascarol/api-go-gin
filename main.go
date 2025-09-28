package main

import (
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/routes"
)

// @title API Go Gin - Swagger Documentation
// @version 1.0
// @description This is a sample API built with Go and Gin framework.

// @host localhost:8080
// @BasePath /
func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
