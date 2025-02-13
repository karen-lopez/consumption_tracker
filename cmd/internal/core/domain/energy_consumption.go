package domain

import "time"

type EnergyConsumption struct {
	ID                 int
	MeterID            int
	ActiveEnergy       int
	ReactiveEnergy     int
	CapacitiveReactive int
	Solar              int
	Date               time.Time
}
