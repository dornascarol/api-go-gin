package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dornascarol/api-go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDeleteUsecase struct {
	mock.Mock
}

func (m *MockDeleteUsecase) DeleteSinger(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestDeleteSingerSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	c.Request, _ = http.NewRequest("DELETE", "/singers/1", nil)

	mockUC := new(MockDeleteUsecase)
	mockUC.On("DeleteSinger", mock.Anything, "1").Return(nil)

	controller := controllers.NewDeleteSingerController(mockUC)
	controller.DeleteSinger(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Singer deleted successfully")
	mockUC.AssertCalled(t, "DeleteSinger", mock.Anything, "1")
}

func TestDeleteSingerNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "id", Value: "999"}}
	c.Request, _ = http.NewRequest("DELETE", "/singers/999", nil)

	mockUC := new(MockDeleteUsecase)
	mockUC.On("DeleteSinger", mock.Anything, "999").Return(assert.AnError)

	controller := controllers.NewDeleteSingerController(mockUC)
	controller.DeleteSinger(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Singer not found")
	mockUC.AssertCalled(t, "DeleteSinger", mock.Anything, "999")
}
