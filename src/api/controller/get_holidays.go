package controller

import (
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/labstack/echo/v4"
)

const (
	TABLE_HOLIDAYS_JP = "holidays_jp"
	COLUMN_DATE       = "holiday_date"
	LOCATION          = "Asia/Tokyo"
)

type HolidaysQ struct {
	StartDate time.Time `query:"start-date"`
	EndDate   time.Time `query:"end-date"`
}

func GetHolidays(c echo.Context) error {
	logger := c.Logger()
	h := new(HolidaysQ)
	var holidays []model.HolidayData
	db := database.GetDbConnection()

	if err := c.Bind(h); err != nil {
		logger.Error("Bind err", err)
		return BadRequestJson(c, "cannot bind query param")
	}
	logger.Debug(h)
	dataSet := db.From(TABLE_HOLIDAYS_JP).Order(goqu.C(COLUMN_DATE).Asc())
	dataSet.ScanStructs(&holidays)

	return c.JSON(http.StatusOK, holidays)
}
