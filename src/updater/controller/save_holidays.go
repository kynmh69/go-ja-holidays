package controller

import (
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
)

const (
	LOCATION = "Asia/Tokyo"
)

func SaveHolidays(holidays []*model.HolidayData) {
	logger := logging.GetLogger()
	latestHoliday := getLatestHoliday()
	if len(holidays) == 0 {
		logger.Warnln("データがないため追加しません")
		return
	}
	newHoliday := holidays[len(holidays)-1]
	if latestHoliday.Date.UTC() != newHoliday.Date.UTC() {
		// 差分を更新
		logger.Infoln("save new holiday", latestHoliday, newHoliday)
		updateData(holidays, latestHoliday)
	} else if latestHoliday.Date.UTC() == newHoliday.Date.UTC() {
		// 差分がない場合
		logger.Warnln("No New data. Did not update")
	} else {
		// 新規登録
		logger.Warnln("First submit")
		firstInsertHolidays(holidays)
	}

}

func getLatestHoliday() *model.HolidayData {
	logger := logging.GetLogger()
	var oldRow model.HolidayData
	db := database.GetDbConnection()
	err := db.First(&oldRow).Order("created_at desc").Error

	if err != nil {
		logger.Panicln(err)
	}
	logger.Infoln("old row found", oldRow)

	return &oldRow
}

func firstInsertHolidays(holidays []*model.HolidayData) {
	db := database.GetDbConnection()
	logger := logging.GetLogger()
	// タムゾーンをCSV元のものに変更
	timeLocation := holidays[0].Date.Location()
	goqu.SetTimeLocation(timeLocation)

	result := db.Create(holidays)
	if result.Error != nil {
		logger.Errorln(result.Error)
	}
	if affected := result.RowsAffected; affected > 0 {
		logger.Infoln("successfull.", affected)
	} else {
		logger.Panicln("Rows Affected err.", affected)
	}
}

func createDiff(newHolidayData []*model.HolidayData, oldHolidayData *model.HolidayData) []*model.HolidayData {
	var updateData []*model.HolidayData
	for _, v := range newHolidayData {
		if v.Date.After(oldHolidayData.Date) {
			updateData = append(updateData, v)
		}
	}
	return updateData
}

func updateData(newHolidayData []*model.HolidayData, oldData *model.HolidayData) {
	logger := logging.GetLogger()
	updateData := createDiff(newHolidayData, oldData)
	var location *time.Location
	var err error
	if len(updateData) > 0 {
		location = updateData[0].Date.Location()
	} else {
		if location, err = time.LoadLocation(LOCATION); err != nil {
			log.Fatalln(err)
		}
	}

	db := database.GetDbConnection()
	goqu.SetTimeLocation(location)

	result := db.Create(updateData)
	if result.Error != nil {
		logger.Panicln(err)
	}

	if affected := result.RowsAffected; affected > 0 {
		logger.Infoln("affected", affected)
	} else {
		logger.Panicln("Rows Affected err.", affected)
	}
}
