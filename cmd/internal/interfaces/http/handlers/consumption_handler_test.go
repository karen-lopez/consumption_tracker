package handlers

import (
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockEnergyConsumptionService struct {
	mock.Mock
}

func (m *MockEnergyConsumptionService) GetConsumption(ctx context.Context, meterID int, startDate, endDate, kindPeriod string) (*dtos.ConsumptionResponse, error) {
	args := m.Called(ctx, meterID, startDate, endDate, kindPeriod)
	return args.Get(0).(*dtos.ConsumptionResponse), args.Error(1)
}

func TestGetConsumption(t *testing.T) {
	mockService := new(MockEnergyConsumptionService)

	handler := NewConsumptionHandler(mockService)

	router := gin.Default()
	router.GET("/consumption", handler.GetConsumption)

	mockService.On("GetConsumption", mock.Anything, 1, "2023-01-01", "2023-01-31", "monthly").Return(&dtos.ConsumptionResponse{}, nil)

	req, _ := http.NewRequest("GET", "/consumption?meters_ids=1&start_date=2023-01-01&end_date=2023-01-31&kind_period=monthly", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
