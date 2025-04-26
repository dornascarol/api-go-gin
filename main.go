package main

import (
	"github.com/dornascarol/api-go-gin/database"
	"github.com/dornascarol/api-go-gin/routes"
)

func main() {
	database.ConnectToDatabase()
	routes.HandleRequests()
}
