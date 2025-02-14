package utils

import "time"

func ParseDateToTime(date string) (time.Time, error) {
	return time.Parse("2025-01-01", date)

}
