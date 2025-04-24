package models

import "gorm.io/gorm"

type Singer struct {
	gorm.Model
	ArtistName   string `json:"artist_name"`
	SongName     string `json:"song_name"`
	MusicalGenre string `json:"musical_genre"`
}

var Singers []Singer
