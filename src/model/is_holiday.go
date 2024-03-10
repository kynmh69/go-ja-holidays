package model

import "time"

type IsHoliday struct {
	IsHoliday bool      `json:"is_holiday" db:"-" goqq:"skipinsert"`
	Date      time.Time `json:"date" db:"holiday_date"`
	Name      string    `json:"name" db:"holiday_name"`
}
