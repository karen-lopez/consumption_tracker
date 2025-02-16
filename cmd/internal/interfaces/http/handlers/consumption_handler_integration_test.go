package handlers

import (
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetConsumptionIntegration(t *testing.T) {
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
