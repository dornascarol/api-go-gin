package usecases

import (
	"context"
	"time"

	"encoding/json"
	"log"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/dornascarol/api-go-gin/domain/repositories"
	"github.com/dornascarol/api-go-gin/infrastructure/cache"
)

type SingersUseCase struct {
	SingerRepo   repositories.SingerRepository
	CacheService cache.CacheServiceInterface
}

func NewSingersUseCase(repo repositories.SingerRepository, cache cache.CacheServiceInterface) *SingersUseCase {
	return &SingersUseCase{
		SingerRepo:   repo,
		CacheService: cache,
	}
}

func (uc *SingersUseCase) GetAllSingers(ctx context.Context) ([]entities.Singer, error) {
	cacheKey := "singers:all"

	cachedData, err := uc.CacheService.Get(ctx, cacheKey)
	if err == nil {
		var cachedSingers []entities.Singer
		if err := json.Unmarshal([]byte(cachedData), &cachedSingers); err == nil {
			return cachedSingers, nil
		}
	}

	singers, err := uc.SingerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(singers)
	if err == nil {
		err = uc.CacheService.Set(ctx, cacheKey, string(jsonData), time.Minute*5)
		if err != nil {
			log.Println("Error saving to cache:", err)
		}
	}

	return singers, nil
}

func (uc *SingersUseCase) CreateSinger(ctx context.Context, singer *entities.Singer) (*entities.Singer, error) {
	if err := uc.SingerRepo.Save(ctx, singer); err != nil {
		return nil, err
	}
	return singer, nil
}

func (uc *SingersUseCase) GetSingerByID(ctx context.Context, id string) (*entities.Singer, error) {
	cacheKey := "singer:" + id

	cachedData, err := uc.CacheService.Get(ctx, cacheKey)
	if err == nil {
		var cachedSinger entities.Singer
		if err := json.Unmarshal([]byte(cachedData), &cachedSinger); err == nil {
			return &cachedSinger, nil
		}
	}

	singer, err := uc.SingerRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(singer)
	if err == nil {
		_ = uc.CacheService.Set(ctx, cacheKey, string(jsonData), time.Minute*5)
	}

	return singer, nil
}

func (uc *SingersUseCase) DeleteSinger(ctx context.Context, id string) error {
	err := uc.SingerRepo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	cacheKey := "singer:" + id
	_ = uc.CacheService.Delete(ctx, cacheKey)

	return nil
}

func (uc *SingersUseCase) UpdateSinger(ctx context.Context, id string, updated *entities.Singer) (*entities.Singer, error) {
	singer, err := uc.SingerRepo.UpdateByID(ctx, id, updated)
	if err != nil {
		return nil, err
	}

	cacheKey := "singer:" + id
	_ = uc.CacheService.Delete(ctx, cacheKey)

	return singer, nil
}

func (uc *SingersUseCase) GetSingerByName(ctx context.Context, name string) (*entities.Singer, error) {
	cacheKey := "singer:name:" + name

	cachedData, err := uc.CacheService.Get(ctx, cacheKey)
	if err == nil {
		var cachedSinger entities.Singer
		if err := json.Unmarshal([]byte(cachedData), &cachedSinger); err == nil {
			return &cachedSinger, nil
		}
	}

	singer, err := uc.SingerRepo.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(singer)
	if err == nil {
		_ = uc.CacheService.Set(ctx, cacheKey, string(jsonData), time.Minute*2)
	}

	return singer, nil
}
