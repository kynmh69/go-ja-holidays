package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/util"
	"net/http/httptest"
	"os"
	"testing"
)

func TestIsHoliday(t *testing.T) {
	_ = os.Setenv("DATABASE", "unittest")
	_ = os.Setenv("PSQL_HOSTNAME", "localhost")
	util.SetUp()
	r := gin.Default()
	r.GET("/holidays/:date", IsHoliday)

	tests := []struct {
		name   string
		target string
	}{
		{
			name:   "TestIsHoliday",
			target: "/holidays/2021-01-01",
		},
		{
			name:   "TestIsHolidayInvalid",
			target: "/holidays/2021-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", tt.target, nil)
			context := gin.CreateTestContextOnly(w, r)
			context.Request = req
			IsHoliday(context)
		})
	}
}
