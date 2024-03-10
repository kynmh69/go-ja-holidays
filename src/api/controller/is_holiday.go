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
	var isHoliday model.IsHoliday
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

	dataSet := db.From(TABLE_HOLIDAYS_JP).
		Where(goqu.C(TABLE_HOLIDAYS_JP).Eq(dayTime))
	_, err = dataSet.ScanStruct(&isHoliday)

	if err != nil {
		return BadRequestJson(c, err.Error())
	}
	return c.JSON(http.StatusOK, isHoliday)
}
