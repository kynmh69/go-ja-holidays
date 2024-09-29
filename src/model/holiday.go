package model

import (
	"gorm.io/gorm"
	"time"
)

type HolidayData struct {
	gorm.Model
	Date time.Time `json:"date" gorm:"holiday_date;unique;not null"`
	Name string    `json:"name" gorm:"holiday_name;unique;not null"`
}
