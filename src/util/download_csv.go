package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DownloadCSV(url string) {
	// HTTP GETリクエストでCSVデータを取得
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("can not get csv file.", err)
	}
	defer resp.Body.Close()

	// レスポンスのボディを読み込み
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("can not read csv file", err)
	}

	// データを標準出力に表示
	fmt.Println(string(data))
}
