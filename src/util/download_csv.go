package util

import (
	"bytes"
	"encoding/csv"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
	"io"
	"log"
	"net/http"
	"time"
)

func downloadCSV(url string) []byte {
	logger := logging.GetLogger()
	// HTTP GETリクエストでCSVデータを取得
	resp, err := http.Get(url)
	if err != nil {
		logger.Panicln("can not get csv file.", err)
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み込み
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Panicln("can not read csv file", err)
	}

	// データを標準出力に表示
	// log.Println(string(data))

	// UTF-8に変換
	utf8Data := ShiftJISToUTF8(data)
	// log.Println(string(utf8Data))
	return utf8Data
}

func CreateHolidayData(url string) []*model.HolidayData {
	byteData := downloadCSV(url)

	var holidayData []*model.HolidayData

	// CSVデータをパース
	holidayData, err := parseCSV(byteData)
	if err != nil {
		log.Fatalln("can not parse csv file", err)
	}

	return holidayData
}

func parseCSV(data []byte) ([]*model.HolidayData, error) {
	logger := logging.GetLogger()
	var holidays []*model.HolidayData
	// CSVデータをパース
	reader := csv.NewReader(bytes.NewReader(data))
	// ロケーションを取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	// CSVデータを一行ずつ読み込み
	for {
		record, err := reader.Read()
		if err != nil {
			break // ファイルの終わりに達したかエラーが発生した
		}

		// 日付を解析
		date, err := time.ParseInLocation("2006/1/2", record[0], jst)
		if err != nil {
			logger.Warnln("日付の解析に失敗:", err)
			continue
		}
		if date.Year() < 1970 {
			continue // 1970年以前のデータは無視する
		}

		// HolidayData構造体に変換して追加
		holidays = append(holidays, &model.HolidayData{
			Date: date,
			Name: record[1],
		})
	}
	// log.Println(holidays)
	return holidays, nil
}
