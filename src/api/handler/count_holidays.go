package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
	"net/http"
)

type CountStruct struct {
	Count int64 `json:"count"`
}

func CountHolidays(c *gin.Context) {
	logger := logging.GetLogger()
	var request HolidaysRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error(err)
		BadRequestJson(c, err.Error())
		return
	}
	logger.Debug("request:", request)
	db := database.GetDbConnection()

	dataSet := db.Model(&model.HolidayData{})
	if !request.StartDay.IsZero() && !request.EndDay.IsZero() {
		dataSet = dataSet.Where("date BETWEEN ? AND ?", request.StartDay, request.EndDay)
	} else if !request.StartDay.IsZero() {
		dataSet = dataSet.Where("date >= ?", request.StartDay)
	} else if !request.EndDay.IsZero() {
		dataSet = dataSet.Where("date <= ?", request.EndDay)
	}
	var count int64
	err := dataSet.Count(&count).Error
	if err != nil {
		BadRequestJson(c, err.Error())
		return
	}
	cs := CountStruct{Count: count}
	logger.Debug("count:", cs)
	c.JSON(http.StatusOK, cs)
}
