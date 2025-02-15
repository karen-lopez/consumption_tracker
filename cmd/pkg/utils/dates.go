package utils

import "time"

func ParseDateToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

func ParseToString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05+00")
}
