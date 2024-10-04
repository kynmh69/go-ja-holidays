package model

import (
	"gorm.io/gorm"
	"time"
)

type HolidayData struct {
	gorm.Model
	Date time.Time `json:"date" gorm:"unique;not null"`
	Name string    `json:"name" gorm:"not null"`
}
