package postgresql

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/core/ports"
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) ports.ConsumptionRepository {
	return &Repository{db: db}
}

func (r *Repository) GetConsumption(ctx context.Context, meterID, startDate, endDate string) ([]domain.EnergyConsumption, error) {
	return nil, nil
}
