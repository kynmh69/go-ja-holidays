package controller

import (
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/labstack/echo/v4"
)

type CountStruct struct {
	Count int64 `json:"count"`
}

func CountHolidays(c echo.Context) error {
	if location, err := time.LoadLocation(LOCATION); err != nil {
		return BadRequestJson(c, err.Error())
	} else {
		goqu.SetTimeLocation(location)
	}

	startDay, err := getStartDate(c)
	if err != nil {
		return BadRequestJson(c, err.Error())
	}

	endDay, err := getEndDate(c)
	if err != nil {
		return BadRequestJson(c, err.Error())
	}

	db := database.GetDbConnection()

	dataSet := db.From(TABLE_HOLIDAYS_JP)
	if startDay != nil && endDay != nil {
		dataSet.Where(
			goqu.C(COLUMN_DATE).Gte(startDay),
			goqu.C(COLUMN_DATE).Lte(endDay),
		)
	} else if startDay != nil {
		dataSet.Where(
			goqu.C(COLUMN_DATE).Gte(startDay),
		)
	} else if endDay != nil {
		dataSet = dataSet.Where(
			goqu.C(COLUMN_DATE).Lte(endDay),
		)
	}
	count, err := dataSet.Count()
	if err != nil {
		return BadRequestJson(c, err.Error())
	}
	stru := CountStruct{Count: count}
	return c.JSON(http.StatusOK, stru)
}
