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

const TABLE_HOLIDAYS_JP = "holidays_jp"

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
		Order(goqu.C("holiday_date").Desc()).
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
		log.Fatalln(err)
	}
	if affected, err := result.RowsAffected(); err == nil {
		log.Println("successfull.", affected)
	} else {
		log.Fatalln("Rows Affected err.", err)
	}
}
