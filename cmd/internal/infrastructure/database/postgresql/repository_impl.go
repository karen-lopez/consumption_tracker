package postgresql

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/core/ports"
	"consumption_tracker/cmd/pkg/errors"
	"consumption_tracker/cmd/pkg/utils"
	"time"

	"context"
	"database/sql"
)

const energyConsumptionQuery = `SELECT meter_id, active_energy, reactive_energy, capacitive_reactive, solar, date
        FROM energy_consumption
        WHERE date BETWEEN $2 AND $3
        ORDER BY date ASC`

type Repository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) ports.ConsumptionRepository {
	return &Repository{db: db}
}

func (r *Repository) GetConsumption(ctx context.Context, meterID int, startDate, endDate, kindPeriod string) ([]domain.EnergyConsumption, error) {

	start, end, errParseDate := parseDates(startDate, endDate)
	if errParseDate != nil {
		return nil, errParseDate
	}

	rows, errQuery := r.db.QueryContext(ctx, energyConsumptionQuery, meterID, start, end)
	if errQuery != nil {
		return nil, errors.ErrSearchingData
	}
	defer rows.Close()

	var consumptions []domain.EnergyConsumption
	for rows.Next() {
		var ec domain.EnergyConsumption
		if err := rows.Scan(&ec.MeterID, &ec.ActiveEnergy, &ec.ReactiveEnergy, &ec.CapacitiveReactive, &ec.Solar, &ec.Date); err != nil {
			return nil, errors.ErrScanningData
		}
		consumptions = append(consumptions, ec)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.ErrIteratingData
	}

	return consumptions, nil
}

func parseDates(startDate string, endDate string) (time.Time, time.Time, error) {
	start, errStartDate := utils.ParseDateToTime(startDate)
	end, errEndDate := utils.ParseDateToTime(endDate)
	if errStartDate != nil || errEndDate != nil {
		return start, end, errors.ErrParsingDate
	}
	return start, end, nil
}
