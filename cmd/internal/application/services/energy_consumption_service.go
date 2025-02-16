package services

import (
	"consumption_tracker/cmd/internal/core/domain"
	"consumption_tracker/cmd/internal/core/ports"
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"consumption_tracker/cmd/pkg/utils"
	"context"
	"fmt"
	"log"
	"time"
)

const (
	MonthFormat = "Jan 2006"
	WeekFormat  = "Jan 02"
	DayFormat   = "Jan 2"
	Month       = "monthly"
	Week        = "weekly"
	Day         = "daily"
	DaysInWeek  = 7
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
		return getConsumptionsByRange(consumptions, address, DayFormat, startDate), nil
	case Month:
		return getConsumptionsByRange(consumptions, address, MonthFormat, startDate), nil
	case Week:
		return getConsumptionsByRange(consumptions, address, WeekFormat, startDate), nil
	}
	return nil, nil
}

func getConsumptionsByRange(consumptions []domain.EnergyConsumption, address, periodFormat, startDate string) *dtos.ConsumptionResponse {
	var groupedData *dtos.MeterData
	var currentPeriod time.Time
	periodIndex := 0
	isWeek := periodFormat == WeekFormat
	var periods []string

	for index, c := range consumptions {
		date := c.Date.Format(periodFormat)
		if index == 0 {
			log.Print("First consumption: ", c)
			groupedData = initializeGroupedData(c, address)
			currentPeriod, periods = initializePeriods(isWeek, startDate, periodFormat, periods)
		}
		if isNextRangePeriod(currentPeriod, c.Date, isWeek, periodFormat) {
			currentPeriod, periodIndex, periods = nextPeriod(currentPeriod, periodIndex, groupedData, date, periods, isWeek)
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

func initializePeriods(isWeek bool, startDate string, periodFormat string, periods []string) (time.Time, []string) {
	currentPeriod, err := utils.ParseDateToTime(startDate)
	if err != nil {
		return time.Time{}, nil
	}
	log.Print("Current Period: ", currentPeriod)
	if isWeek {
		endWeek := currentPeriod.AddDate(0, 0, DaysInWeek-1)
		log.Println("Current Period : ", currentPeriod)
		weekPeriod := fmt.Sprintf("%s - %s", currentPeriod.Format(periodFormat), endWeek.Format(periodFormat))
		periods = append(periods, weekPeriod)
		log.Println("endWeek Period: ", endWeek)
		log.Println("weekPeriod: ", weekPeriod)
	} else {
		periods = append(periods, currentPeriod.Format(periodFormat))
	}
	return currentPeriod, periods
}

func nextPeriod(currentPeriod time.Time, periodIndex int, groupedData *dtos.MeterData, date string, periods []string, isWeek bool) (time.Time, int, []string) {
	periodIndex++
	period := date
	groupedData.Active = append(groupedData.Active, 0)
	groupedData.ReactiveInductive = append(groupedData.ReactiveInductive, 0)
	groupedData.ReactiveCapacitive = append(groupedData.ReactiveCapacitive, 0)
	groupedData.Exported = append(groupedData.Exported, 0)
	if isWeek {
		startNextWeek := currentPeriod.AddDate(0, 0, DaysInWeek)
		endNextWeek := startNextWeek.AddDate(0, 0, DaysInWeek)
		period = fmt.Sprintf("%s - %s", startNextWeek.Format(WeekFormat), endNextWeek.Format(WeekFormat))
		currentPeriod = startNextWeek
	}
	periods = append(periods, period)
	return currentPeriod, periodIndex, periods
}

func isNextRangePeriod(currentPeriod time.Time, date time.Time, isWeek bool, periodFormat string) bool {
	if isWeek {
		nextPeriodDate := currentPeriod.AddDate(0, 0, DaysInWeek-1)
		return date.After(nextPeriodDate)
	}
	return date.Format(periodFormat) != currentPeriod.Format(periodFormat)
}

func initializeGroupedData(c domain.EnergyConsumption, address string) *dtos.MeterData {
	return &dtos.MeterData{
		MeterID:            c.MeterID,
		Address:            address,
		Active:             []int{c.ActiveEnergy},
		ReactiveInductive:  []int{c.ReactiveEnergy},
		ReactiveCapacitive: []int{c.CapacitiveReactive},
		Exported:           []int{c.Solar},
	}
}

func sumConsumptionToGroupedData(groupedData *dtos.MeterData, periodIndex int, c domain.EnergyConsumption) {
	groupedData.Active[periodIndex] += c.ActiveEnergy
	groupedData.ReactiveInductive[periodIndex] += c.ReactiveEnergy
	groupedData.ReactiveCapacitive[periodIndex] += c.CapacitiveReactive
	groupedData.Exported[periodIndex] += c.Solar
}
