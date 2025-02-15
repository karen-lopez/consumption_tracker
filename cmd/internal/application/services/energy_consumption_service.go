package services

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/core/ports"
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"context"
	"time"
)

const (
	MonthFormat = "Jan 2024"
	WeekFormat  = "Jan 01 - Jan 07"
	DayFormat   = "Jan 01"
)

type EnergyConsumptionService struct {
	repository     ports.ConsumptionRepository
	addressService ports.AddressService
}

func NewEnergyConsumptionService(repository ports.ConsumptionRepository, addressService ports.AddressService) *EnergyConsumptionService {
	return &EnergyConsumptionService{repository: repository, addressService: addressService}
}

func (s *EnergyConsumptionService) GetConsumption(ctx context.Context, meterId int, startDate, endDate, kindPeriod string) (*dtos.ConsumptionResponse, error) {
	consumptions, err := s.repository.GetConsumption(ctx, meterId, startDate, endDate)
	if err != nil {
		return nil, err
	}
	address, addressErr := s.addressService.GetAddressByMeterID(ctx, meterId)
	if addressErr != nil {
		return nil, addressErr
	}
	switch kindPeriod {
	case "day":
		return getConsumptionsByRange(consumptions, address, DayFormat), nil
	case "month":
		return getConsumptionsByRange(consumptions, address, MonthFormat), nil
	case "week":
		return getConsumptionsByRange(consumptions, address, WeekFormat), nil
	}
	return nil, nil
}

func getConsumptionsByRange(consumptions []domain.EnergyConsumption, address, periodFormat string) *dtos.ConsumptionResponse {
	var groupedData *dtos.MeterData
	periodIndex := 0
	isWeek := periodFormat == WeekFormat
	if isWeek {
		periodFormat = DayFormat
	}
	currentPeriod := consumptions[0].Date.Format(periodFormat)
	periods := []string{currentPeriod}

	for index, c := range consumptions {
		period := c.Date.Format(periodFormat)
		if isNextRangePeriod(c.Date, currentPeriod, period, isWeek) {
			periodIndex++
			currentPeriod = period
			periods = append(periods, period)
		}
		if index == 0 {
			groupedData = &dtos.MeterData{
				MeterID:            c.MeterID,
				Address:            address,
				Active:             []int{c.ActiveEnergy},
				ReactiveInductive:  []int{c.ReactiveEnergy},
				ReactiveCapacitive: []int{c.CapacitiveReactive},
				Exported:           []int{c.Solar},
			}
		}
		sumConsumptionToGroupedData(groupedData, periodIndex, c)
	}

	response := &dtos.ConsumptionResponse{
		Period:    periods,
		DataGraph: []*dtos.MeterData{groupedData},
	}
	return response
}

func isNextRangePeriod(date time.Time, currentPeriod string, period string, isWeek bool) bool {
	if isWeek {
		nextDate := date.AddDate(0, 0, 6)
		return nextDate.Format(DayFormat) != currentPeriod
	}
	return period != currentPeriod
}

func sumConsumptionToGroupedData(groupedData *dtos.MeterData, periodIndex int, c domain.EnergyConsumption) {
	groupedData.Active[periodIndex] += c.ActiveEnergy
	groupedData.ReactiveInductive[periodIndex] += c.ReactiveEnergy
	groupedData.ReactiveCapacitive[periodIndex] += c.CapacitiveReactive
	groupedData.Exported[periodIndex] += c.Solar
}
