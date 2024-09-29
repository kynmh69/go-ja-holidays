package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
)

const (
	ColumnDate = "holiday_date"
	LOCATION   = "Asia/Tokyo"
)

type HolidaysRequest struct {
	StartDay time.Time `form:"start_day" time_format:"2006-01-02"`
	EndDay   time.Time `form:"end_day" time_format:"2006-01-02"`
}

func (receiver HolidaysRequest) String() string {
	return fmt.Sprintf("StartDay: \"%s\", EndDay: \"%s\"", receiver.StartDay, receiver.EndDay)
}

func GetHolidays(c *gin.Context) {
	var (
		reqParams HolidaysRequest
		holidays  []model.HolidayData
	)
	// DB接続
	db := database.GetDbConnection()
	// タイムゾーン設定
	if location, err := time.LoadLocation(LOCATION); err != nil {
		BadRequestJson(c, err.Error())
		return
	} else {
		// クエリビルダーにタイムゾーンを設定
		goqu.SetTimeLocation(location)
	}
	// リクエストパラメータを取得
	if err := c.BindQuery(&reqParams); err != nil {
		BadRequestJson(c, err.Error())
		return
	}
	// リクエストパラメータから開始日と終了日を取得
	dataSet := db.Model(&model.HolidayData{}).Order(goqu.C(ColumnDate).Asc())
	if !reqParams.StartDay.IsZero() && !reqParams.EndDay.IsZero() {
		dataSet = dataSet.Where("date between ? and ?", reqParams.StartDay, reqParams.EndDay)
	} else if !reqParams.StartDay.IsZero() {
		dataSet = dataSet.Where("date >= ?", reqParams.StartDay)
	} else if !reqParams.EndDay.IsZero() {
		dataSet = dataSet.Where("date <= ?", reqParams.EndDay)
	}
	// データ取得
	if err := dataSet.Find(&holidays).Error; err != nil {
		// エラー時はBad Requestを返却
		BadRequestJson(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, holidays)
}
