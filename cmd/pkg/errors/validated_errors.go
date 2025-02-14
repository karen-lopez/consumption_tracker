package errors

import "errors"

var (
	ErrInvalidMeterID = errors.New("invalid meter ID")
	ErrInvalidAddress = errors.New("invalid address")
	ErrInvalidInput   = errors.New("input is invalid")
	ErrInvalidDate    = errors.New("invalid date")
	ErrInvalidPeriod  = errors.New("invalid period")
)
