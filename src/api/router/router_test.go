package router

import (
	"log"
	"os"
	"testing"

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
func TestMakeRoute(t *testing.T) {
	e := echo.New()
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{e},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeRoute(tt.args.e)
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
