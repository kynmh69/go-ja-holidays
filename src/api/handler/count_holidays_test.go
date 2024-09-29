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
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/holidays/count", nil)
	context := gin.CreateTestContextOnly(w, r)
	context.Request = req
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestCountHolidays",
			args: args{c: context},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CountHolidays(tt.args.c)
		})
	}
}
