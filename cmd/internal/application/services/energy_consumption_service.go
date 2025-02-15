package services

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/core/ports"
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"context"
	"log"
	"time"
)

const (
	MonthFormat = "Jan 2006"
	WeekFormat  = "Jan 01 - Jan 07"
	DayFormat   = "Jan 02"
	Month       = "monthly"
	Week        = "weekly"
	Day         = "daily"
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
	if len(consumptions) == 0 {
		return nil, nil
	}
	address, addressErr := s.addressService.GetAddressByMeterID(ctx, meterId)
	if addressErr != nil {
		return nil, addressErr
	}
	switch kindPeriod {
	case Day:
		return getConsumptionsByRange(consumptions, address, DayFormat), nil
	case Month:
		return getConsumptionsByRange(consumptions, address, MonthFormat), nil
	case Week:
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
			groupedData.Active = append(groupedData.Active, 0)
			groupedData.ReactiveInductive = append(groupedData.ReactiveInductive, 0)
			groupedData.ReactiveCapacitive = append(groupedData.ReactiveCapacitive, 0)
			groupedData.Exported = append(groupedData.Exported, 0)
			currentPeriod = period
			periods = append(periods, period)
		}
		if index == 0 {
			log.Print("First consumption: ", c)
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
		log.Print("Grouped Data : ", groupedData)
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
