package util

import "testing"

func TestDownloadCSV(t *testing.T) {
	okUrl := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	ngUrl := "https://www8.cao.go.jp/chosei/shukujitsu/sy"
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "download csv ok", args: args{url: okUrl}},
		{name: "download csv ng", args: args{url: ngUrl}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownloadCSV(tt.args.url)
		})
	}
}
