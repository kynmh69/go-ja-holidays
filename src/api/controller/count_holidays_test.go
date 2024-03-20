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
	q.Set("end-day", "2023-01-01")
	req := httptest.NewRequest(http.MethodGet, "/holidays/count", nil)
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
		// TODO: Add test cases.
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
