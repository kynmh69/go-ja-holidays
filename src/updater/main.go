package main

import (
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/updater/controller"
	"github.com/kynmh69/go-ja-holidays/util"
)

func init() {
	util.LoggerInitialize()
	database.ConnectDatabase()
}

func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

	holidays := util.CreateHolidayData(url)

	controller.SaveHolidays(holidays)
}
