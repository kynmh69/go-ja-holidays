package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func TestIsHoliday(t *testing.T) {
	// Setup
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/holidays/:day")
	c.SetParamNames("day")
	c.SetParamValues("2023-01-01")

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
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
			if err := IsHoliday(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IsHoliday() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsHolidayFmErr(t *testing.T) {
	// Setup
	e := echo.New()
	util.EchoLoggerInitialize(e)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/holidays/:day")
	c.SetParamNames("day")
	c.SetParamValues("2023-01-0")

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "NG",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsHoliday(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("IsHoliday() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
