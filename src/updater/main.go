package main

import "github.com/kynmh69/go-ja-holidays/util"

func main() {
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"

	util.CreateHolidayData(url)
}
