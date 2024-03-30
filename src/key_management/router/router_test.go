package router

import (
	"testing"

	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func TestMakeRoute(t *testing.T) {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test ok",
			args: args{e: e},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeRoute(tt.args.e)
		})
	}
}
