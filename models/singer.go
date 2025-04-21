package models

type Singer struct {
	ArtistName   string `json:"artist_name"`
	SongName     string `json:"song_name"`
	MusicalGenre string `json:"musical_genre"`
}

var Singers []Singer
