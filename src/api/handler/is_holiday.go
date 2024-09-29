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
	holiday.Date = request.Date
	result := db.First(&holiday)
	err := result.Error
	if err != nil {
		logger.Error(err)
		BadRequestJson(c, err.Error())
		return
	}

	var isHoliday model.IsHoliday
	if result.RowsAffected > 0 {
		// If the holiday data exists, return it.
		isHoliday = model.IsHoliday{IsHoliday: true, HolidayData: holiday}
	} else {
		//	If the holiday data does not exist, return the date.
		isHoliday = model.IsHoliday{IsHoliday: false, HolidayData: holiday}
	}
	logger.Debug(isHoliday)
	c.JSON(http.StatusOK, isHoliday)
}
