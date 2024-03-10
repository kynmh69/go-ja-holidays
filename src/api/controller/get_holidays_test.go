package controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()
	res := m.Run()
	os.Exit(res)
}
func TestGetHolidays(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("end-day", "2023-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getStartDate(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("start-day", "2023-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	ti := time.Date(2023, 1, 1, 0, 0, 0, 0, loc)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				c: c,
			},
			want:    &ti,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStartDate(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStartDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStartDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEndDate(t *testing.T) {
	// Setup
	e := echo.New()
	q := make(url.Values)
	q.Set("end-day", "2023-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	ti := time.Date(2023, 1, 1, 0, 0, 0, 0, loc)
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				c: c,
			},
			want:    &ti,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getEndDate(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("getEndDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEndDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ok",
			args: args{
				"2023-01-01",
			},
			want: true,
		},
		{
			name: "ng",
			args: args{
				"2023-01-0",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDate(tt.args.dateStr); got != tt.want {
				t.Errorf("isValidDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createTime(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/holidays", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	loc, _ := time.LoadLocation("Asia/Tokyo")
	ti := time.Date(2023, 1, 1, 0, 0, 0, 0, loc)
	type args struct {
		c       echo.Context
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *time.Time
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				c:       c,
				dateStr: "2023-01-01",
			},
			want:    &ti,
			wantErr: false,
		},
		{
			name: "NG",
			args: args{
				c:       c,
				dateStr: "2023-01-0",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createTime(tt.args.c, tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("createTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setUp() {
	os.Setenv("PSQL_HOSTNAME", "localhost")
	os.Setenv("DATABASE", "unittest")
	url := "https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
	database.ConnectDatabase()
	util.CreateHolidayData(url)
}

func tearDown() {
	os.Unsetenv("PSQL_HOSTNAME")
	os.Unsetenv("DATABASE")
	db := database.GetDbConnection()
	if _, err := db.Delete("holidays_jp").Executor().Exec(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Teardown.")
}
