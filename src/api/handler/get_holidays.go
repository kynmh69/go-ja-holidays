package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
)

const (
	TableHolidaysJp = "holidays_jp"
	ColumnDate      = "holiday_date"
	LOCATION        = "Asia/Tokyo"
)

type HolidaysRequest struct {
	StartDay string `form:"start-day" binding:"datetime"`
	EndDay   string `form:"end-day" binding:"datetime"`
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
	dataSet := db.From(TableHolidaysJp).Order(goqu.C(ColumnDate).Asc())
	if reqParams.StartDay != "" && reqParams.EndDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Gte(reqParams.StartDay),
			goqu.C(ColumnDate).Lte(reqParams.EndDay),
		)
	} else if reqParams.StartDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Gte(reqParams.StartDay),
		)
	} else if reqParams.EndDay != "" {
		dataSet = dataSet.Where(
			goqu.C(ColumnDate).Lte(reqParams.EndDay),
		)
	}
	// データ取得
	if err := dataSet.ScanStructs(&holidays); err != nil {
		// エラー時はBad Requestを返却
		BadRequestJson(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, holidays)
}
