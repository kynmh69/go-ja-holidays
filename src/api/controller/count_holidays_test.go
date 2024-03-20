package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestCountHolidays(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("start-day", "2022-01-01")
	q.Set("end-day", "2023-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count?"+q.Encode(), nil)
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
			name: "test OK",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CountHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CountHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountHolidaysStartDay(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("start-day", "2022-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count?"+q.Encode(), nil)
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
			name: "test OK",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CountHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CountHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountHolidaysEndDay(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("end-day", "2022-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count?"+q.Encode(), nil)
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
			name: "test OK",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CountHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CountHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountHolidaysStartDayErr(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("start-day", "2022-01-0")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count?"+q.Encode(), nil)
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
			name: "test err",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CountHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CountHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountHolidaysEndDayErr(t *testing.T) {
	e := echo.New()
	q := make(url.Values)
	q.Set("end-day", "2022-01-0")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count?"+q.Encode(), nil)
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
			name: "test err",
			args: args{
				c: c,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CountHolidays(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CountHolidays() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
