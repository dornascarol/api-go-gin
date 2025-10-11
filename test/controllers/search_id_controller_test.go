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

type MockSearchByIDUsecase struct {
	mock.Mock
}

func (m *MockSearchByIDUsecase) GetSingerByID(ctx context.Context, id string) (*entities.Singer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func TestSearchSingerByIdSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request, _ = http.NewRequest("GET", "/singers/1", nil)

	expected := &entities.Singer{
		ArtistName:   "Alegria do tan tan",
		SongName:     "Lua de mel",
		MusicalGenre: "Pagode",
	}

	mockUC := new(MockSearchByIDUsecase)
	mockUC.On("GetSingerByID", mock.Anything, "1").Return(expected, nil)

	controller := controllers.NewSearchSingerByIdController(mockUC)
	controller.SearchSingerById(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Lua de mel")
	mockUC.AssertCalled(t, "GetSingerByID", mock.Anything, "1")
}

func TestSearchSingerByIdNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "999"}}
	c.Request, _ = http.NewRequest("GET", "/singers/999", nil)

	mockUC := new(MockSearchByIDUsecase)
	mockUC.On("GetSingerByID", mock.Anything, "999").Return(nil, assert.AnError)

	controller := controllers.NewSearchSingerByIdController(mockUC)
	controller.SearchSingerById(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Singer not found")
	mockUC.AssertCalled(t, "GetSingerByID", mock.Anything, "999")
}
