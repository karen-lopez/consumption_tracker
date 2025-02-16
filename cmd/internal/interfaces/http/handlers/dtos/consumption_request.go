package dtos

import (
	"consumption_tracker/cmd/pkg/errors"
	"time"
)

type ConsumptionRequest struct {
	MeterID    string `json:"meters_ids" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	KindPeriod string `json:"kind_period" binding:"required"`
}

func (consumptionRequest *ConsumptionRequest) Validate() error {
	if consumptionRequest.MeterID == "" {
		return errors.ErrInvalidMeterID
	}
	if _, err := time.Parse("2025-01-01", consumptionRequest.StartDate); err != nil {
		return errors.ErrInvalidDate
	}
	if _, err := time.Parse("2025-01-01", consumptionRequest.EndDate); err != nil {
		return errors.ErrInvalidDate
	}
	if consumptionRequest.KindPeriod == "" {
		return errors.ErrInvalidPeriod
	}
	return nil
}
