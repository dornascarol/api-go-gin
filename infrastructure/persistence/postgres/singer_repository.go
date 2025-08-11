package postgres

import (
	"context"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"gorm.io/gorm"
)

type PostgresSingerRepository struct {
	DB *gorm.DB
}

func NewPostgresSingerRepository(db *gorm.DB) *PostgresSingerRepository {
	return &PostgresSingerRepository{DB: db}
}

func (r *PostgresSingerRepository) FindAll(ctx context.Context) ([]entities.Singer, error) {
	var singers []entities.Singer
	err := r.DB.WithContext(ctx).Find(&singers).Error
	return singers, err
}

func (r *PostgresSingerRepository) Save(ctx context.Context, singer *entities.Singer) error {
	return r.DB.WithContext(ctx).Create(singer).Error
}

func (r *PostgresSingerRepository) FindByID(ctx context.Context, id string) (*entities.Singer, error) {
	var singer entities.Singer
	err := r.DB.WithContext(ctx).First(&singer, id).Error
	if err != nil {
		return nil, err
	}
	return &singer, nil
}

func (r *PostgresSingerRepository) DeleteByID(ctx context.Context, id string) error {
	return r.DB.WithContext(ctx).Delete(&entities.Singer{}, id).Error
}

func (r *PostgresSingerRepository) UpdateByID(ctx context.Context, id string, updated *entities.Singer) (*entities.Singer, error) {
	var singer entities.Singer
	if err := r.DB.WithContext(ctx).First(&singer, id).Error; err != nil {
		return nil, err
	}

	singer.ArtistName = updated.ArtistName
	singer.SongName = updated.SongName
	singer.MusicalGenre = updated.MusicalGenre

	if err := r.DB.WithContext(ctx).Save(&singer).Error; err != nil {
		return nil, err
	}

	return &singer, nil
}

func (r *PostgresSingerRepository) FindByName(ctx context.Context, name string) (*entities.Singer, error) {
	var singer entities.Singer
	err := r.DB.WithContext(ctx).Where("artist_name = ?", name).First(&singer).Error
	if err != nil {
		return nil, err
	}
	return &singer, nil
}
