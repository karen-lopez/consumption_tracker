package domain

type MeterAddress struct {
	ID      int
	MeterID int
	Address string
}

func (m *MeterAddress) Validate() error {
	if m.MeterID <= 0 {
		return ErrInvalidMeterID
	}
	if m.Address == "" {
		return ErrInvalidAddress
	}
	return nil
}
