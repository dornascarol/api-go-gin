package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dornascarol/api-go-gin/domain/entities"
	"github.com/dornascarol/api-go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) CreateSinger(ctx context.Context, singer *entities.Singer) (*entities.Singer, error) {
	args := m.Called(ctx, singer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func TestCreateNewSingerInvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{invalid json}`)
	c.Request, _ = http.NewRequest("POST", "/singers", body)

	controller := controllers.NewCreateSingerController(nil)
	controller.CreateNewSinger(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateNewSingerInvalidData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	singer := entities.Singer{}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("POST", "/singers", bytes.NewBuffer(jsonData))

	controller := controllers.NewCreateSingerController(nil)
	controller.CreateNewSinger(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateNewSingerSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	singer := entities.Singer{
		ArtistName:   "Samba Raiz",
		SongName:     "Jangada",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("POST", "/singers", bytes.NewBuffer(jsonData))

	mockUC := new(MockUsecase)
	mockUC.On("CreateSinger", mock.Anything, &singer).Return(&singer, nil)

	controller := controllers.NewCreateSingerController(mockUC)
	controller.CreateNewSinger(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUC.AssertCalled(t, "CreateSinger", mock.Anything, &singer)
}

func TestCreateNewSingerInternalError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	singer := entities.Singer{
		ArtistName:   "Samba Nutela",
		SongName:     "A ilha",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("POST", "/singers", bytes.NewBuffer(jsonData))

	mockUC := new(MockUsecase)
	mockUC.On("CreateSinger", mock.Anything, &singer).Return(nil, assert.AnError)

	controller := controllers.NewCreateSingerController(mockUC)
	controller.CreateNewSinger(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUC.AssertCalled(t, "CreateSinger", mock.Anything, &singer)
}
