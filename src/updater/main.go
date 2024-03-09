package main

import (
	"log"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/updater/controller"
	"github.com/kynmh69/go-ja-holidays/util"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[App] ")

	database.ConnectDatabase()
}

func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

	holidays := util.CreateHolidayData(url)

	controller.SaveHolidays(holidays)
}
