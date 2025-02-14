package domain

import "time"

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
		return ErrInvalidInput
	}
	if e.MeterAddress == "" {
		return ErrInvalidInput
	}
	if e.ActiveEnergy < 0 || e.ReactiveEnergy < 0 || e.CapacitiveReactive < 0 || e.Solar < 0 {
		return ErrInvalidInput
	}
	if e.Date.IsZero() {
		return ErrInvalidInput
	}
	return nil
}
