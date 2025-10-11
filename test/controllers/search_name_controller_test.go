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

type MockSearchByNameUsecase struct {
	mock.Mock
}

func (m *MockSearchByNameUsecase) GetSingerByName(ctx context.Context, name string) (*entities.Singer, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func TestSearchSingerByNameSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Vany"}}
	c.Request, _ = http.NewRequest("GET", "/singers/name/Vany", nil)

	expected := &entities.Singer{
		ArtistName:   "Vany",
		SongName:     "Sua estante",
		MusicalGenre: "Pagode",
	}

	mockUC := new(MockSearchByNameUsecase)
	mockUC.On("GetSingerByName", mock.Anything, "Vany").Return(expected, nil)

	controller := controllers.NewSearchSingerByNameController(mockUC)
	controller.SearchSingerByName(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Sua estante")
	mockUC.AssertCalled(t, "GetSingerByName", mock.Anything, "Vany")
}

func TestSearchSingerByNameNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "name", Value: "Luks"}}
	c.Request, _ = http.NewRequest("GET", "/singers/name/Luks", nil)

	mockUC := new(MockSearchByNameUsecase)
	mockUC.On("GetSingerByName", mock.Anything, "Luks").Return(nil, assert.AnError)

	controller := controllers.NewSearchSingerByNameController(mockUC)
	controller.SearchSingerByName(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Singer not found")
	mockUC.AssertCalled(t, "GetSingerByName", mock.Anything, "Luks")
}
