package domain

import (
	"consumption_tracker/cmd/pkg/errors"
)

type MeterAddress struct {
	ID      int
	MeterID int
	Address string
}

func (m *MeterAddress) Validate() error {
	if m.MeterID <= 0 {
		return errors.ErrInvalidMeterID
	}
	if m.Address == "" {
		return errors.ErrInvalidAddress
	}
	return nil
}
