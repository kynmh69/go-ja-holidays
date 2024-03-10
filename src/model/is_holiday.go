package model

type IsHoliday struct {
	IsHoliday bool `json:"is_holiday" db:"-" goqu:"skipinsert"`
	HolidayData
}
