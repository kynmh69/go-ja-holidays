package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/util"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logging.LoggerInitialize()
	_ = os.Setenv("DATABASE", "unittest")
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	util.SetUp()
	code := m.Run()
	if code > 0 {
		logging.GetLogger().Error("exitcode: ", code)
	}
}

func TestCountHolidays(t *testing.T) {
	r := gin.Default()
	r.GET("/holidays/count", CountHolidays)
	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "TestCountHolidays",
			target: "/holidays/count",
		},
		{
			name:   "TestCountHolidaysWithQuery",
			target: "/holidays/count?start_day=2021-01-01&end_day=2021-12-31",
		},
		{
			name:   "TestCountHolidaysWithQueryStartDay",
			target: "/holidays/count?start_day=2021-01-01",
		},
		{
			name:   "TestCountHolidaysWithQueryEndDay",
			target: "/holidays/count?end_day=2021-12-31",
		},
		{
			name:   "TestCountHolidaysWithQueryInvalid",
			target: "/holidays/count?start_day=2021-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.target, nil)
			context := gin.CreateTestContextOnly(w, r)
			context.Request = req
			CountHolidays(context)
		})
	}
}
