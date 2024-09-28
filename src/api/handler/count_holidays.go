package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
)

type CountStruct struct {
	Count int64 `json:"count"`
}

func CountHolidays(c *gin.Context) {
	logger := logging.GetLogger()
	var request HolidaysRequest
	if location, err := time.LoadLocation(LOCATION); err != nil {
		BadRequestJson(c, err.Error())
		return
	} else {
		goqu.SetTimeLocation(location)
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error(err)
		BadRequestJson(c, err.Error())
		return
	}

	db := database.GetDbConnection()

	dataSet := db.From(TableHolidaysJp)
	if request.StartDay != "" && request.EndDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Gte(request.StartDay),
			goqu.C(ColumnDate).Lte(request.EndDay),
		)
	} else if request.StartDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Gte(request.StartDay),
		)
	} else if request.EndDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Lte(request.EndDay),
		)
	}
	count, err := dataSet.Count()
	if err != nil {
		BadRequestJson(c, err.Error())
		return
	}
	cs := CountStruct{Count: count}
	logger.Debug("count:", cs)
	c.JSON(http.StatusOK, cs)
}
