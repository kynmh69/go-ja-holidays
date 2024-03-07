package util

import (
	"reflect"
	"testing"
	"time"

	"github.com/kynmh69/go-ja-holidays/model"
)

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
			downloadCSV(tt.args.url)
		})
	}
}

func TestParseCSV(t *testing.T) {
	expected := []model.HolidayData{
		{
			Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			Name: "元日",
		},
		{
			Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			Name: "成人の日",
		},
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []model.HolidayData
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "parse csv ok",
			args:    args{data: []byte("date,name\n2023/1/1,元日\n2023/1/2,成人の日")},
			want:    expected,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCSV(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateHolidayData(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want []model.HolidayData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateHolidayData(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHolidayData() = %v, want %v", got, tt.want)
			}
		})
	}
}
