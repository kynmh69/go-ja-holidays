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

	dataSet := db.From(TABLE_HOLIDAYS_JP).Order(goqu.C(COLUMN_DATE).Asc())
	if startDay != nil && endDay != nil {
		dataSet = dataSet.Where(
			goqu.C(COLUMN_DATE).Gte(startDay),
			goqu.C(COLUMN_DATE).Lte(endDay),
		)
	} else if startDay != nil {
		dataSet = dataSet.Where(
			goqu.C(COLUMN_DATE).Gte(startDay),
		)
	} else if endDay != nil {
		dataSet = dataSet.Where(
			goqu.C(COLUMN_DATE).Lte(endDay),
		)
	}
	dataSet.ScanStructs(&holidays)

	return c.JSON(http.StatusOK, holidays)
}

func getStartDate(c echo.Context) (*time.Time, error) {
	logger := c.Logger()
	startDateStr := c.QueryParam("start-day")
	logger.Debug("start-day: ", startDateStr)
	if startDateStr == "" {
		return nil, nil
	}
	startDate, err := createTime(c, startDateStr)
	if err != nil {
		return nil, err
	}
	if startDate == nil {
		return nil, errors.New("cannot parse start day")
	}

	return startDate, nil
}

func getEndDate(c echo.Context) (*time.Time, error) {
	logger := c.Logger()
	endDateStr := c.QueryParam("end-day")
	logger.Debug("end-day: ", endDateStr)
	if endDateStr == "" {
		return nil, nil
	}
	endDate, err := createTime(c, endDateStr)
	if err != nil {
		return nil, err
	}
	if endDate == nil {
		return nil, errors.New("cannot parse end day")
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

func createTime(c echo.Context, dateStr string) (*time.Time, error) {
	logger := c.Logger()
	layout := "2006-01-02"
	location, err := time.LoadLocation(LOCATION)
	if err != nil {
		log.Fatalln(err)
	}
	if !isValidDate(dateStr) {
		return nil, errors.New("the date format is invalid. like " + layout)
	}

	parseTime, err := time.ParseInLocation(layout, dateStr, location)
	logger.Debug(parseTime)

	if err != nil {
		return nil, err
	}

	return &parseTime, nil
}
