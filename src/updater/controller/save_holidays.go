package controller

import (
	"log"

	"github.com/doug-martin/goqu"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/model"
)

type HolidayData struct {
	Date string `db:"date"`
	Name string `db:"name"`
}

func SaveHolidays(holidays []model.HolidayData) {
	var oldRow HolidayData
	db := database.GetDbConnection()
	found, err := db.From("ja").Order(goqu.I("date").Desc()).Limit(1).ScanStruct(&oldRow)

	if err != nil {
		log.Fatalln(err)
	}

	if found {
		log.Println(oldRow)
	}

}
