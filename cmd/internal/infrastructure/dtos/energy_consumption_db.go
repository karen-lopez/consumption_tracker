package dtos

import "time"

type EnergyConsumptionDB struct {
	ID                 int       `db:"id"`
	MeterID            int       `db:"meter_id"`
	ActiveEnergy       int       `db:"active_energy"`
	ReactiveEnergy     int       `db:"reactive_energy"`
	CapacitiveReactive int       `db:"capacitive_reactive"`
	Solar              int       `db:"solar"`
	ConsumptionDate    time.Time `db:"consumption_date"`
}
