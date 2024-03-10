package controller

import (
	"net/http"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func IsHoliday(c echo.Context) error {
	var holiday model.HolidayData
	db := database.GetDbConnection()

	logger := c.Logger()

	day := c.Param("day")
	logger.Debug(day)
	if day == "" {
		msg := "cannot get day."
		logger.Warn(msg)
		return BadRequestJson(c, msg)
	}
	dayTime, err := util.ParseDay(day)
	if err != nil {
		return BadRequestJson(c, err.Error())
	}

	loc := dayTime.Location()
	goqu.SetTimeLocation(loc)

	dataSet := db.From(TABLE_HOLIDAYS_JP).
		Where(goqu.C(COLUMN_DATE).Eq(dayTime))
	ok, err := dataSet.ScanStruct(&holiday)

	if err != nil {
		logger.Error(err)
		return BadRequestJson(c, err.Error())
	}

	var isHoliday model.IsHoliday
	if ok {
		isHoliday = model.IsHoliday{IsHoliday: ok, HolidayData: holiday}
	} else {
		isHoliday = model.IsHoliday{IsHoliday: ok, HolidayData: model.HolidayData{Date: *dayTime}}
	}
	logger.Debug(isHoliday)
	return c.JSON(http.StatusOK, isHoliday)
}
