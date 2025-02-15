package postgresql

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/infrastructure/dtos"
	"consumption_tracker/cmd/pkg/errors"
	"consumption_tracker/cmd/pkg/utils"
	"log"

	"context"
	"database/sql"
)

const energyConsumptionQuery = `SELECT meter_id, active_energy, reactive_energy, capacitive_reactive, solar, consumption_date 
		FROM energy_consumption
        WHERE meter_id = $1 AND consumption_date BETWEEN $2 AND $3
        ORDER BY consumption_date ASC`

type Repository struct {
	db *sql.DB
}

func NewPostgresqlRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetConsumption(ctx context.Context, meterID int, startDate, endDate string) ([]domain.EnergyConsumption, error) {
	start, end, errParseDate := parseDates(startDate, endDate)
	if errParseDate != nil {
		log.Print("Error parsing dates " + errParseDate.Error())
		return nil, errParseDate
	}

	rows, errQuery := r.db.QueryContext(ctx, energyConsumptionQuery, meterID, start, end)
	if errQuery != nil {
		return nil, errors.ErrSearchingData
	}
	defer rows.Close()

	var consumptionsDB []dtos.EnergyConsumptionDB
	for rows.Next() {
		var ec dtos.EnergyConsumptionDB
		if err := rows.Scan(&ec.MeterID, &ec.ActiveEnergy, &ec.ReactiveEnergy, &ec.CapacitiveReactive, &ec.Solar, &ec.ConsumptionDate); err != nil {
			return nil, errors.ErrScanningData
		}
		consumptionsDB = append(consumptionsDB, ec)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.ErrIteratingData
	}
	return dtoToDomainEntity(consumptionsDB)
}

func dtoToDomainEntity(consumptionsDB []dtos.EnergyConsumptionDB) ([]domain.EnergyConsumption, error) {
	var consumptions []domain.EnergyConsumption
	for index, consumption := range consumptionsDB {
		consumptions = append(consumptions, domain.EnergyConsumption{
			MeterID:            consumption.MeterID,
			ActiveEnergy:       consumption.ActiveEnergy,
			ReactiveEnergy:     consumption.ReactiveEnergy,
			CapacitiveReactive: consumption.CapacitiveReactive,
			Solar:              consumption.Solar,
			Date:               consumption.ConsumptionDate,
		})
		if consumptions[index].Validate() != nil {
			return nil, errors.ErrParsingData
		}
	}
	return consumptions, nil
}

func parseDates(startDate string, endDate string) (string, string, error) {
	start, errStartDate := utils.ParseDateToTime(startDate)
	end, errEndDate := utils.ParseDateToTime(endDate)
	if errStartDate != nil || errEndDate != nil {
		return "", "", errors.ErrParsingDate
	}
	formatStartDate := utils.ParseToString(start)
	formatEndDate := utils.ParseToString(end)
	log.Println("Start ConsumptionDate: ", formatStartDate)
	return formatStartDate, formatEndDate, nil
}
