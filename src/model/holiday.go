package model

import "time"

type HolidayData struct {
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}
