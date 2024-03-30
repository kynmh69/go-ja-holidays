package middleware

import (
	"reflect"
	"testing"

	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Test_getLogFormat(t *testing.T) {
	format := "${time_rfc3339_nano} [${method}] ${uri} ${status} ${id}\n"
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test ok",
			want: format,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLogFormat(); got != tt.want {
				t.Errorf("getLogFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_confLogFormat(t *testing.T) {
	tests := []struct {
		name string
		want mid.LoggerConfig
	}{
		{
			name: "test OK",
			want: mid.LoggerConfig{Format: getLogFormat(), Output: util.InitWriter()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confLogFormat(); !reflect.DeepEqual(got.Format, tt.want.Format) {
				t.Errorf("confLogFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setLogger(t *testing.T) {
	tests := []struct {
		name string
		want echo.MiddlewareFunc
	}{
		{
			"test OK",
			mid.LoggerWithConfig(confLogFormat()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("setLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
