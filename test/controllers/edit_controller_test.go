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

type MockEditUsecase struct {
	mock.Mock
}

func (m *MockEditUsecase) UpdateSinger(ctx context.Context, id string, updated *entities.Singer) (*entities.Singer, error) {
	args := m.Called(ctx, id, updated)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Singer), args.Error(1)
}

func TestEditSingerInvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	body := bytes.NewBufferString(`{invalid json}`)
	c.Request, _ = http.NewRequest("PATCH", "/singers/1", body)

	controller := controllers.NewEditSingerController(nil)
	controller.EditSinger(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEditSingerInvalidData(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	singer := entities.Singer{}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("PATCH", "/singers/1", bytes.NewBuffer(jsonData))

	controller := controllers.NewEditSingerController(nil)
	controller.EditSinger(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestEditSingerSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	singer := entities.Singer{
		ArtistName:   "Samba no pé",
		SongName:     "Não fale assim",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("PATCH", "/singers/1", bytes.NewBuffer(jsonData))

	mockUC := new(MockEditUsecase)
	mockUC.On("UpdateSinger", mock.Anything, "1", &singer).Return(&singer, nil)

	controller := controllers.NewEditSingerController(mockUC)
	controller.EditSinger(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Não fale assim")
	mockUC.AssertCalled(t, "UpdateSinger", mock.Anything, "1", &singer)
}

func TestEditSingerNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "999"}}
	singer := entities.Singer{
		ArtistName:   "Pagode na cabeça",
		SongName:     "O tempo de volta",
		MusicalGenre: "Pagode",
	}
	jsonData, _ := json.Marshal(singer)
	c.Request, _ = http.NewRequest("PATCH", "/singers/999", bytes.NewBuffer(jsonData))

	mockUC := new(MockEditUsecase)
	mockUC.On("UpdateSinger", mock.Anything, "999", &singer).Return(nil, assert.AnError)

	controller := controllers.NewEditSingerController(mockUC)
	controller.EditSinger(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Singer not found")
	mockUC.AssertCalled(t, "UpdateSinger", mock.Anything, "999", &singer)
}
