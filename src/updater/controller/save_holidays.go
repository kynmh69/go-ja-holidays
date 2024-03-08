package controller

import (
	"log"
	"time"

	"github.com/kynmh69/go-ja-holidays/database"
)

type HolidayDbData struct {
	Date time.Time `db:"holiday_date"`
	Name string    `db:"holiday_name"`
}

const TABLE_HOLIDAYS_JP = "holidays_jp"

func SaveHolidays(holidays []HolidayDbData) {
	latestHoliday, ok := getLatestHoliday()
	newHoliday := holidays[len(holidays)-1]
	if ok && latestHoliday.Date != newHoliday.Date {
		// 差分を更新
		log.Println("save new holiday")
	} else {
		// 新規登録
		firstInsertHolidays(holidays)
	}

}

func getLatestHoliday() (HolidayDbData, bool) {
	var oldRow HolidayDbData
	db := database.GetDbConnection()
	ok, err := db.From(TABLE_HOLIDAYS_JP).ScanStruct(&oldRow)

	if err != nil {
		log.Fatalln(err)
	}

	if ok {
		log.Println("old row found")
	} else {
		log.Println("old row not found")
	}

	return oldRow, ok
}

func firstInsertHolidays(holidays []HolidayDbData) {
	db := database.GetDbConnection()
	result, err := db.Insert(TABLE_HOLIDAYS_JP).Rows(holidays).Executor().Exec()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)
}
