package main

import (
	"github.com/dornascarol/api-go-gin/models"
	"github.com/dornascarol/api-go-gin/routes"
)

func main() {
	models.Singers = []models.Singer{
		{ArtistName: "Péricles", SongName: "Até que durou", MusicalGenre: "Pagode"},
		{ArtistName: "BTS", SongName: "Just One Day", MusicalGenre: "K-Pop"},
	}
	routes.HandleRequests()
}
