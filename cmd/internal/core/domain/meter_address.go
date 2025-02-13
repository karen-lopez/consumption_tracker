package domain

type MeterAddress struct {
	ID      int    `json:"id"`
	MeterID int    `json:"meter_id"`
	Address string `json:"address"`
}
