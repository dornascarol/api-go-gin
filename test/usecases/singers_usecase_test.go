package usecases_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCacheService struct {
	mock.Mock
}

func (m *MockCacheService) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockCacheService) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	args := m.Called(ctx, key, value, ttl)
	return args.Error(0)
}

func (m *MockCacheService) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

type MockSingerRepository struct {
	mock.Mock
}

func (m *MockSingerRepository) FindByID(ctx context.Context, id string) (*entities.Singer, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func (m *MockSingerRepository) DeleteByID(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockSingerRepository) FindAll(ctx context.Context) ([]entities.Singer, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entities.Singer), args.Error(1)
}

func (m *MockSingerRepository) Save(ctx context.Context, singer *entities.Singer) error {
	args := m.Called(ctx, singer)
	return args.Error(0)
}

func (m *MockSingerRepository) UpdateByID(ctx context.Context, id string, singer *entities.Singer) (*entities.Singer, error) {
	args := m.Called(ctx, id, singer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func (m *MockSingerRepository) FindByName(ctx context.Context, name string) (*entities.Singer, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func TestGetSingerByIDCacheHit(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := &entities.Singer{ArtistName: "Sambaí", SongName: "Falando com os astros", MusicalGenre: "Pagode"}
	jsonData, _ := json.Marshal(expected)

	mockCache.On("Get", mock.Anything, "singer:1").Return(string(jsonData), nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetSingerByID(context.Background(), "1")

	assert.NoError(t, err)
	assert.Equal(t, expected.ArtistName, result.ArtistName)
	mockCache.AssertCalled(t, "Get", mock.Anything, "singer:1")
}

func TestGetSingerByIDCacheMiss(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := &entities.Singer{ArtistName: "XPTO Samba", SongName: "Vem pra ser meu refrão", MusicalGenre: "Pagode"}

	mockCache.On("Get", mock.Anything, "singer:2").Return("", assert.AnError)
	mockRepo.On("FindByID", mock.Anything, "2").Return(expected, nil)
	mockCache.On("Set", mock.Anything, "singer:2", mock.Anything, time.Minute*5).Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetSingerByID(context.Background(), "2")

	assert.NoError(t, err)
	assert.Equal(t, expected.MusicalGenre, result.MusicalGenre)
	mockRepo.AssertCalled(t, "FindByID", mock.Anything, "2")
	mockCache.AssertCalled(t, "Set", mock.Anything, "singer:2", mock.Anything, time.Minute*5)
}

func TestGetAllSingersCacheHit(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := []entities.Singer{
		{ArtistName: "Samba de Esquina", SongName: "Desejo me chama", MusicalGenre: "Pagode"},
		{ArtistName: "Os Sambistas", SongName: "Pensa bem", MusicalGenre: "Pagode"},
	}
	jsonData, _ := json.Marshal(expected)

	mockCache.On("Get", mock.Anything, "singers:all").Return(string(jsonData), nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetAllSingers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expected[0].ArtistName, result[0].ArtistName)
	mockCache.AssertCalled(t, "Get", mock.Anything, "singers:all")
}

func TestGetAllSingersCacheMiss(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := []entities.Singer{
		{ArtistName: "Cavaquinho do Som", SongName: "Chega Aí", MusicalGenre: "Pagode"},
	}

	mockCache.On("Get", mock.Anything, "singers:all").Return("", assert.AnError)
	mockRepo.On("FindAll", mock.Anything).Return(expected, nil)
	mockCache.On("Set", mock.Anything, "singers:all", mock.Anything, time.Minute*5).Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetAllSingers(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, expected[0].SongName, result[0].SongName)
	mockRepo.AssertCalled(t, "FindAll", mock.Anything)
	mockCache.AssertCalled(t, "Set", mock.Anything, "singers:all", mock.Anything, time.Minute*5)
}

func TestCreateSingerSuccess(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	newSinger := &entities.Singer{
		ArtistName:   "Todos juntos na canção",
		SongName:     "Dom de amar",
		MusicalGenre: "Pagode",
	}

	mockRepo.On("Save", mock.Anything, newSinger).Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.CreateSinger(context.Background(), newSinger)

	assert.NoError(t, err)
	assert.Equal(t, newSinger.ArtistName, result.ArtistName)
	mockRepo.AssertCalled(t, "Save", mock.Anything, newSinger)
}

func TestCreateSingerError(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	newSinger := &entities.Singer{
		ArtistName:   "Harmonia do Pagode",
		SongName:     "Céu e estrelas",
		MusicalGenre: "Pagode",
	}

	mockRepo.On("Save", mock.Anything, newSinger).Return(assert.AnError)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.CreateSinger(context.Background(), newSinger)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertCalled(t, "Save", mock.Anything, newSinger)
}

func TestUpdateSingerSuccess(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	updatedSinger := &entities.Singer{
		ArtistName:   "Carol Bela",
		SongName:     "Minha canção favorita",
		MusicalGenre: "Pagode",
	}

	mockRepo.On("UpdateByID", mock.Anything, "1", updatedSinger).Return(updatedSinger, nil)
	mockCache.On("Delete", mock.Anything, "singer:1").Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.UpdateSinger(context.Background(), "1", updatedSinger)

	assert.NoError(t, err)
	assert.Equal(t, updatedSinger.SongName, result.SongName)
	mockRepo.AssertCalled(t, "UpdateByID", mock.Anything, "1", updatedSinger)
	mockCache.AssertCalled(t, "Delete", mock.Anything, "singer:1")
}

func TestUpdateSingerError(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	updatedSinger := &entities.Singer{
		ArtistName:   "Violão Verde",
		SongName:     "Escuta aqui",
		MusicalGenre: "Rock",
	}

	mockRepo.On("UpdateByID", mock.Anything, "2", updatedSinger).Return(nil, assert.AnError)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.UpdateSinger(context.Background(), "2", updatedSinger)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertCalled(t, "UpdateByID", mock.Anything, "2", updatedSinger)
	mockCache.AssertNotCalled(t, "Delete", mock.Anything, "singer:2")
}

func TestGetSingerByNameCacheHit(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := &entities.Singer{
		ArtistName:   "Quarteto do Samba",
		SongName:     "Jogo de sedução",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(expected)

	mockCache.On("Get", mock.Anything, "singer:name:Quarteto do Samba").Return(string(jsonData), nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetSingerByName(context.Background(), "Quarteto do Samba")

	assert.NoError(t, err)
	assert.Equal(t, expected.SongName, result.SongName)
	mockCache.AssertCalled(t, "Get", mock.Anything, "singer:name:Quarteto do Samba")
	mockRepo.AssertNotCalled(t, "FindByName", mock.Anything, "Quarteto so Samba")
}

func TestGetSingerByNameCacheMiss(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	expected := &entities.Singer{
		ArtistName:   "Admira Samba",
		SongName:     "Doce feitiço",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(expected)

	mockCache.On("Get", mock.Anything, "singer:name:Admira Samba").Return("", assert.AnError)
	mockRepo.On("FindByName", mock.Anything, "Admira Samba").Return(expected, nil)
	mockCache.On("Set", mock.Anything, "singer:name:Admira Samba", string(jsonData), time.Minute*2).Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	result, err := usecase.GetSingerByName(context.Background(), "Admira Samba")

	assert.NoError(t, err)
	assert.Equal(t, expected.MusicalGenre, result.MusicalGenre)
	mockRepo.AssertCalled(t, "FindByName", mock.Anything, "Admira Samba")
	mockCache.AssertCalled(t, "Set", mock.Anything, "singer:name:Admira Samba", string(jsonData), time.Minute*2)
}

func TestDeleteSingerSuccess(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	mockRepo.On("DeleteByID", mock.Anything, "1").Return(nil)
	mockCache.On("Delete", mock.Anything, "singer:1").Return(nil)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	err := usecase.DeleteSinger(context.Background(), "1")

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteByID", mock.Anything, "1")
	mockCache.AssertCalled(t, "Delete", mock.Anything, "singer:1")
}

func TestDeleteSingerError(t *testing.T) {
	mockRepo := new(MockSingerRepository)
	mockCache := new(MockCacheService)

	mockRepo.On("DeleteByID", mock.Anything, "2").Return(assert.AnError)

	usecase := usecases.NewSingersUseCase(mockRepo, mockCache)
	err := usecase.DeleteSinger(context.Background(), "2")

	assert.Error(t, err)
	mockRepo.AssertCalled(t, "DeleteByID", mock.Anything, "2")
	mockCache.AssertNotCalled(t, "Delete", mock.Anything, "singer:2")
}
