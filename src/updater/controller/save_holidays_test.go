package controller

import (
	"github.com/kynmh69/go-ja-holidays/logging"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kynmh69/go-ja-holidays/database"
)

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSaveHolidays(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}
	holidays := []HolidayDbData{
		{
			Date: time.Date(2004, 1, 4, 0, 0, 0, 0, loc),
			Name: "テスト祝日4",
		},
	}
	type args struct {
		holidays []HolidayDbData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				holidays: holidays,
			},
		},
		{
			name: "ok new submit",
			args: args{
				holidays: holidays,
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case 1:
				db := database.GetDbConnection()
				db.Delete(TABLE_HOLIDAYS_JP)
			}
			SaveHolidays(tt.args.holidays)
		})
	}
}

func Test_getLatestHoliday(t *testing.T) {
	wants, _ := getLatestHoliday()
	tests := []struct {
		name  string
		want  HolidayDbData
		want1 bool
	}{
		{
			name:  "ok",
			want:  wants,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLatestHoliday()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLatestHoliday() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLatestHoliday() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_firstInsertHolidays(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	holidaysD := []HolidayDbData{
		{Date: time.Date(2002, 1, 1, 0, 0, 0, 0, loc), Name: "テスト祝日1"},
		{Date: time.Date(2002, 1, 2, 0, 0, 0, 0, loc), Name: "テスト祝日2"},
		{Date: time.Date(2002, 1, 3, 0, 0, 0, 0, loc), Name: "テスト祝日3"},
		{Date: time.Date(2002, 1, 4, 0, 0, 0, 0, loc), Name: "テスト祝日4"},
	}
	type args struct {
		holidays []HolidayDbData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{holidaysD},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			firstInsertHolidays(tt.args.holidays)
		})
	}
}

func Test_createDiff(t *testing.T) {
	newHolidayData := []HolidayDbData{
		{Date: time.Now().AddDate(1, 0, 0), Name: "テスト祝日"},
	}
	oldData := HolidayDbData{Date: time.Now(), Name: "テスト祝日最後"}
	type args struct {
		newHolidayData []HolidayDbData
		oldHolidayData HolidayDbData
	}
	tests := []struct {
		name string
		args args
		want []HolidayDbData
	}{
		{
			name: "ok",
			args: args{
				newHolidayData: newHolidayData,
				oldHolidayData: oldData,
			},
			want: newHolidayData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createDiff(tt.args.newHolidayData, tt.args.oldHolidayData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createDiff() = %v, want %v", got, tt.want)
			}
		})
	}
	tearDown()
}

func Test_updateData(t *testing.T) {
	newHolidayData := []HolidayDbData{
		{Date: time.Now().AddDate(1, 0, 0), Name: "テスト祝日"},
	}
	oldData := HolidayDbData{Date: time.Now(), Name: "テスト祝日最後"}
	type args struct {
		newHolidayData []HolidayDbData
		oldData        HolidayDbData
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"ok",
			args{
				newHolidayData: newHolidayData,
				oldData:        oldData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateData(tt.args.newHolidayData, tt.args.oldData)
		})
	}
	tearDown()
}

func setUp() {
	logging.LoggerInitialize()
	_ = os.Setenv("DATABASE", "unittest")
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalln(err)
	}
	holidays := []HolidayDbData{
		{Date: time.Date(2004, 1, 1, 0, 0, 0, 0, loc), Name: "テスト祝日1"},
		{Date: time.Date(2004, 1, 2, 0, 0, 0, 0, loc), Name: "テスト祝日2"},
		{Date: time.Date(2004, 1, 3, 0, 0, 0, 0, loc), Name: "テスト祝日3"},
		{Date: time.Date(2004, 1, 4, 0, 0, 0, 0, loc), Name: "テスト祝日4"},
	}
	database.ConnectDatabase()
	firstInsertHolidays(holidays)
}

func tearDown() {
	db := database.GetDbConnection()
	if _, err := db.Delete(TABLE_HOLIDAYS_JP).Executor().Exec(); err != nil {
		log.Fatalln(err)
	}
	_ = os.Unsetenv("DATABASE")
	_ = os.Unsetenv("PSQL_HOSTNAME")
	log.Println("Tear down.")
}
