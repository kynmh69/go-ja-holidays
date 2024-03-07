package main

import (
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/util"
)

func init() {
	database.ConnectDatabase()
}

func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

	util.CreateHolidayData(url)
	// log.Println(holidays)
}
