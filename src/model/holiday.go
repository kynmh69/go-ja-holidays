package model

import "time"

type HolidayData struct {
	Date time.Time `json:"date" db:"holiday_date"`
	Name string    `json:"name" db:"holiday_name"`
}
