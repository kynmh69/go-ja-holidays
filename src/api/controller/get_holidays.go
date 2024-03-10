package controller

import (
	"errors"
	"log"
	"net/http"
	"regexp"
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

func GetHolidays(c echo.Context) error {
	var holidays []model.HolidayData
	db := database.GetDbConnection()

	startDay, err := getStartDate(c)
	if err != nil {
		return BadRequestJson(c, err.Error())
	}

	endDay, err := getEndDate(c)
	if err != nil {
		return BadRequestJson(c, err.Error())
	}

	dataSet := db.From(TABLE_HOLIDAYS_JP).Order(goqu.C(COLUMN_DATE).Asc())
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
		dataSet.Where(
			goqu.C(COLUMN_DATE).Lte(endDay),
		)
	}
	dataSet.ScanStructs(&holidays)

	return c.JSON(http.StatusOK, holidays)
}

func getStartDate(c echo.Context) (*time.Time, error) {
	startDateStr := c.QueryParam("start-day")
	if startDateStr == "" {
		return nil, nil
	}
	startDate := createTime(startDateStr)
	if startDate == nil {
		return nil, errors.New("cannot parse start day")
	}

	return startDate, nil
}

func getEndDate(c echo.Context) (*time.Time, error) {
	endDateStr := c.QueryParam("end-day")
	if endDateStr == "" {
		return nil, nil
	}
	endDate := createTime(endDateStr)
	if endDate == nil {
		return nil, errors.New("cannot parse start day")
	}

	return endDate, nil
}

func isValidDate(dateStr string) bool {
	// 正規表現パターンを作成
	pattern := `^\d{4}-\d{2}-\d{2}$`
	matched, err := regexp.MatchString(pattern, dateStr)
	if err != nil {
		return false
	}
	return matched
}

func createTime(dateStr string) *time.Time {
	layout := "2006-01-02"
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	if !isValidDate(dateStr) {
		return nil
	}

	parseTime, err := time.ParseInLocation(layout, dateStr, location)

	if err != nil {
		log.Fatalln(err)
	}

	return &parseTime
}
