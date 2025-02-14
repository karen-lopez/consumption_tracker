package domain

import (
	"consumption_tracker/cmd/pkg/errors"
	"time"
)

type EnergyConsumption struct {
	ID                 int
	MeterID            int
	MeterAddress       string
	ActiveEnergy       int
	ReactiveEnergy     int
	CapacitiveReactive int
	Solar              int
	Date               time.Time
}

func (e *EnergyConsumption) Validate() error {
	if e.MeterID <= 0 {
		return errors.ErrInvalidMeterID
	}
	if e.MeterAddress == "" {
		return errors.ErrInvalidInput
	}
	if e.ActiveEnergy < 0 || e.ReactiveEnergy < 0 || e.CapacitiveReactive < 0 || e.Solar < 0 {
		return errors.ErrInvalidInput
	}
	if e.Date.IsZero() {
		return errors.ErrInvalidInput
	}
	return nil
}
