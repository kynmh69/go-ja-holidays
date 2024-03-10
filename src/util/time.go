package util

import "time"

func ParseDay(day string) (*time.Time, error) {
	layout := "2006-01-02"
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	dayTime, err := time.ParseInLocation(layout, day, loc)
	if err != nil {
		return nil, err
	}
	return &dayTime, nil
}
