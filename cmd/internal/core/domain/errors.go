package domain

import "errors"

var (
	ErrInvalidMeterID = errors.New("invalid meter ID")
	ErrInvalidAddress = errors.New("invalid address")
	ErrInvalidInput   = errors.New("input is invalid")
)
