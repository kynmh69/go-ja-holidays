package handler

import (
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
)

type HolidayRequest struct {
	Date time.Time `uri:"date" time_format:"2006-01-02" time_utc:"false"`
}

func IsHoliday(c *gin.Context) {
	var (
		request HolidayRequest
		holiday model.HolidayData
	)
	db := database.GetDbConnection()

	logger := logging.GetLogger()

	if err := c.ShouldBindUri(&request); err != nil {
		logger.Error(err)
		BadRequestJson(c, err.Error())
		return
	}
	logger.Debug(request.Date)

	// Set the time zone to JST.
	loc := request.Date.Location()
	goqu.SetTimeLocation(loc)

	// Get the holiday data for the specified day.
	dataSet := db.From(TableHolidaysJp).
		Where(goqu.C(ColumnDate).Eq(request.Date))
	ok, err := dataSet.ScanStruct(&holiday)

	if err != nil {
		logger.Error(err)
		BadRequestJson(c, err.Error())
		return
	}

	var isHoliday model.IsHoliday
	if ok {
		// If the holiday data exists, return it.
		isHoliday = model.IsHoliday{IsHoliday: ok, HolidayData: holiday}
	} else {
		//	If the holiday data does not exist, return the date.
		isHoliday = model.IsHoliday{IsHoliday: ok, HolidayData: model.HolidayData{Date: request.Date}}
	}
	logger.Debug(isHoliday)
	c.JSON(http.StatusOK, isHoliday)
}
