package controller

import (
	"log"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/kynmh69/go-ja-holidays/database"
)

type HolidayDbData struct {
	Uuid      string    `db:"id" goqu:"skipinsert"`
	Date      time.Time `db:"holiday_date"`
	Name      string    `db:"holiday_name"`
	CreatedAt time.Time `db:"created_at" goqu:"skipinsert"`
	UpdatedAt time.Time `db:"updated_at" goqu:"skipinsert"`
}

const (
	TABLE_HOLIDAYS_JP = "holidays_jp"
	COLUMN_DATE       = "holiday_date"
	LOCATION          = "Asia/Tokyo"
)

func SaveHolidays(holidays []HolidayDbData) {
	latestHoliday, ok := getLatestHoliday()
	if len(holidays) == 0 {
		log.Println("データがないため追加しません")
		return
	}
	newHoliday := holidays[len(holidays)-1]
	if ok && latestHoliday.Date.UTC() != newHoliday.Date.UTC() {
		// 差分を更新
		log.Println("save new holiday", latestHoliday, newHoliday)
		updateData(holidays, latestHoliday)
	} else if latestHoliday.Date.UTC() == newHoliday.Date.UTC() {
		// 差分がない場合
		log.Println("No New data. Did not update", TABLE_HOLIDAYS_JP)
	} else {
		// 新規登録
		log.Println("First submit")
		firstInsertHolidays(holidays)
	}

}

func getLatestHoliday() (HolidayDbData, bool) {
	var oldRow HolidayDbData
	db := database.GetDbConnection()
	ok, err := db.From(TABLE_HOLIDAYS_JP).
		Order(goqu.C(COLUMN_DATE).Desc()).
		ScanStruct(&oldRow)

	if err != nil {
		log.Fatalln(err)
	}

	if ok {
		log.Println("old row found", oldRow)
	} else {
		log.Println("old row not found")
	}

	return oldRow, ok
}

func firstInsertHolidays(holidays []HolidayDbData) {
	db := database.GetDbConnection()

	// タムゾーンをCSV元のものに変更
	timeLocation := holidays[0].Date.Location()
	goqu.SetTimeLocation(timeLocation)

	result, err := db.Insert(TABLE_HOLIDAYS_JP).Rows(holidays).Executor().Exec()
	if err != nil {
		log.Println(err)
	}
	if affected, err := result.RowsAffected(); err == nil {
		log.Println("successfull.", affected)
	} else {
		log.Fatalln("Rows Affected err.", err)
	}
}

func createDiff(newHolidayData []HolidayDbData, oldHolidayData HolidayDbData) []HolidayDbData {
	var updateData []HolidayDbData
	for _, v := range newHolidayData {
		if v.Date.After(oldHolidayData.Date) {
			updateData = append(updateData, v)
		}
	}
	return updateData
}

func updateData(newHolidayData []HolidayDbData, oldData HolidayDbData) {
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

	result, err := db.Insert(TABLE_HOLIDAYS_JP).Rows(updateData).Executor().Exec()

	if err != nil {
		log.Fatalln(err)
	}

	if affected, err := result.RowsAffected(); err == nil {
		log.Println("affected", affected)
	} else {
		log.Fatalln(err)
	}
}
