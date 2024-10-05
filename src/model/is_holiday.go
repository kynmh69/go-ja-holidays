package model

type IsHoliday struct {
	IsHoliday bool `json:"is_holiday" gorm:"-:migration;-:all"`
	HolidayData
}
