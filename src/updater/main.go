package main

import (
	"log"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/updater/controller"
	"github.com/kynmh69/go-ja-holidays/util"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[App] ")

	database.ConnectDatabase()

	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

	holidays := util.CreateHolidayData(url)
	// log.Println(holidays)

	controller.SaveHolidays(holidays)
}
