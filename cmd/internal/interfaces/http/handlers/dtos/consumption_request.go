package dtos

import "time"

type ConsumptionRequest struct {
	MeterID    string `json:"meters_ids" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	KindPeriod string `json:"kind_period" binding:"required"`
}

func (consumptionRequest *ConsumptionRequest) Validate() error {
	if consumptionRequest.MeterID == "" {
		return ErrInvalidMeterID
	}
	if _, err := time.Parse("2025-01-01", consumptionRequest.StartDate); err != nil {
		return ErrInvalidDate
	}
	if _, err := time.Parse("2025-01-01", consumptionRequest.EndDate); err != nil {
		return ErrInvalidDate
	}
	if consumptionRequest.KindPeriod == "" {
		return ErrInvalidPeriod
	}
	return nil
}
