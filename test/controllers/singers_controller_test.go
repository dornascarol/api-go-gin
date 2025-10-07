package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/dornascarol/api-go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetAllUsecase struct {
	mock.Mock
}

func (m *MockGetAllUsecase) GetAllSingers(ctx context.Context) ([]entities.Singer, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Singer), args.Error(1)
}

func TestGetSingersSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/singers", nil)

	expected := []entities.Singer{
		{ArtistName: "Sambba", SongName: "Anos luz", MusicalGenre: "Pagode"},
		{ArtistName: "Pagogo", SongName: "Valeu a pena", MusicalGenre: "Pagode"},
	}

	mockUC := new(MockGetAllUsecase)
	mockUC.On("GetAllSingers", mock.Anything).Return(expected, nil)

	controller := controllers.NewSingerController(mockUC)
	controller.GetSingers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Anos luz")
	assert.Contains(t, w.Body.String(), "Pagogo")
	mockUC.AssertCalled(t, "GetAllSingers", mock.Anything)
}

func TestGetSingersInternalError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/singers", nil)

	mockUC := new(MockGetAllUsecase)
	mockUC.On("GetAllSingers", mock.Anything).Return(nil, assert.AnError)

	controller := controllers.NewSingerController(mockUC)
	controller.GetSingers(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error searching for singers")
	mockUC.AssertCalled(t, "GetAllSingers", mock.Anything)
}
