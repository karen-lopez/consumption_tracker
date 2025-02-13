package dtos

import "errors"

var (
	ErrInvalidMeterID = errors.New("invalid meter ID")
	ErrInvalidDate    = errors.New("invalid date")
	ErrInvalidPeriod  = errors.New("invalid period")
)
