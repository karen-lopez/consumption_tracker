package services

import (
	"consumption_tracker/cmd/internal/core/domain"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockConsumptionRepository struct {
	mock.Mock
}

func (m *MockConsumptionRepository) GetConsumption(ctx context.Context, meterID int, startDate, endDate string) ([]domain.EnergyConsumption, error) {
	args := m.Called(ctx, meterID, startDate, endDate)
	return args.Get(0).([]domain.EnergyConsumption), args.Error(1)
}

type MockAddressService struct {
	mock.Mock
}

func (m *MockAddressService) GetAddressByMeterID(ctx context.Context, meterID int) (string, error) {
	args := m.Called(ctx, meterID)
	return args.String(0), args.Error(1)
}

func TestGetConsumption(t *testing.T) {
	mockRepo := new(MockConsumptionRepository)
	mockAddressService := new(MockAddressService)

	service := NewEnergyConsumptionService(mockRepo, mockAddressService)

	consumptions := []domain.EnergyConsumption{
		{MeterID: 1, ActiveEnergy: 1000, ReactiveEnergy: 200, CapacitiveReactive: 50, Solar: 300, Date: time.Now()},
	}
	address := "address mock 123"

	mockRepo.On("GetConsumption", mock.Anything, 1, "2023-01-01", "2023-01-31").Return(consumptions, nil)
	mockAddressService.On("GetAddressByMeterID", mock.Anything, 1).Return(address, nil)

	response, err := service.GetConsumption(context.Background(), 1, "2023-01-01", "2023-01-31", "monthly")

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, address, response.DataGraph[0].Address)
	mockRepo.AssertExpectations(t)
	mockAddressService.AssertExpectations(t)
}
