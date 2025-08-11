package repositories

import (
	"context"

	"github.com/dornascarol/api-go-gin/domain/entities"
)

type SingerRepository interface {
	FindAll(ctx context.Context) ([]entities.Singer, error)
	Save(ctx context.Context, singer *entities.Singer) error
	FindByID(ctx context.Context, id string) (*entities.Singer, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string, updated *entities.Singer) (*entities.Singer, error)
	FindByName(ctx context.Context, name string) (*entities.Singer, error)
}
