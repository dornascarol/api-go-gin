package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Singer struct {
	gorm.Model
	ArtistName   string `json:"artist_name" validate:"nonzero"`
	SongName     string `json:"song_name" validate:"nonzero"`
	MusicalGenre string `json:"musical_genre" validate:"len=6"`
}

func ValidateSingerData(singer *Singer) error {
	if err := validator.Validate(singer); err != nil {
		return err
	}
	return nil
}
