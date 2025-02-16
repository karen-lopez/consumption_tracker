package ports

import (
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"context"
)

type ConsumptionTrackerService interface {
	GetConsumption(ctx context.Context, meterId int, startDate, endDate, kindPeriod string) (*dtos.ConsumptionResponse, error)
}
