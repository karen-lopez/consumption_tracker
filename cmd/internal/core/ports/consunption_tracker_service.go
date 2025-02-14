package ports

import (
	"consumption_tracker/cmd/internal/core/domain"
	"context"
)

type ConsumptionTrackerService interface {
	GetConsumption(ctx context.Context, meterId, startDate, endDate string) ([]domain.EnergyConsumption, error)
	GetConsumptionWithAddress(ctx context.Context, meterId, startDate, endDate string) ([]domain.EnergyConsumption, error)
}
