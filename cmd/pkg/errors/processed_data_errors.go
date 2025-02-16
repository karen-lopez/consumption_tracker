package errors

import "errors"

var (
	ErrSearchingData = errors.New("error searching data")
	ErrScanningData  = errors.New("error scanning data")
	ErrIteratingData = errors.New("error iterating data")
	ErrParsingDate   = errors.New("error parsing date")
	ErrParsingData   = errors.New("error parsing data")
)
