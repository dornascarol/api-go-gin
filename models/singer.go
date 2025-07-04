package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Singer struct {
	gorm.Model
	ArtistName   string `json:"artist_name" validate:"nonzero, min=1, max=50"`
	SongName     string `json:"song_name" validate:"nonzero, min=1, max=100"`
	MusicalGenre string `json:"musical_genre" validate:"nonzero, len=6, regexp=^[a-zA-Z]*$"`
}

type GreetingResponse struct {
	Message string `json:"message"`
}

type DeleteResponse struct {
	Data string `json:"data"`
}

func ValidateSingerData(singer *Singer) error {
	if err := validator.Validate(singer); err != nil {
		return err
	}
	return nil
}
