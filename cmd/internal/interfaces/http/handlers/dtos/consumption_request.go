package dtos

type ConsumptionRequest struct {
	MeterID    string `json:"meters_ids" binding:"required"`
	StartDate  string `json:"start_date" binding:"required"`
	EndDate    string `json:"end_date" binding:"required"`
	KindPeriod string `json:"kind_period" binding:"required"`
}
