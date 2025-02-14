package ports

import (
	"consumption_tracker/cmd/internal/core/domain"
	"context"
)

type ConsumptionRepository interface {
	GetConsumption(ctx context.Context, meterID int, startDate, endDate string) ([]domain.EnergyConsumption, error)
}
