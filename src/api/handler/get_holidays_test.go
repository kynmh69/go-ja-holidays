package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/util"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestGetHolidays(t *testing.T) {
	_ = os.Setenv("DATABASE", "unittest")
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	util.SetUp()
	r := gin.Default()
	r.GET("/holidays", GetHolidays)
	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "TestGetHolidays",
			target: "/holidays",
		},
		{
			name:   "TestGetHolidaysWithQuery",
			target: "/holidays?start_day=2021-01-01&end_day=2021-12-31",
		},
		{
			name:   "TestGetHolidaysWithQueryStartDay",
			target: "/holidays?start_day=2021-01-01",
		},
		{
			name:   "TestGetHolidaysWithQueryEndDay",
			target: "/holidays?end_day=2021-12-31",
		},
		{
			name:   "TestGetHolidaysWithQueryInvalid",
			target: "/holidays?start_day=2021-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.target, nil)
			context := gin.CreateTestContextOnly(w, r)
			context.Request = req
			GetHolidays(context)
		})
	}
}

func TestHolidaysRequest_String(t *testing.T) {
	util.SetUp()
	now := time.Now()
	type fields struct {
		StartDay time.Time
		EndDay   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestHolidaysRequest_String",
			fields: fields{
				StartDay: now,
				EndDay:   now,
			},
			want: "StartDay: \"" + now.String() + "\", EndDay: \"" + now.String() + "\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := HolidaysRequest{
				StartDay: tt.fields.StartDay,
				EndDay:   tt.fields.EndDay,
			}
			if got := receiver.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
