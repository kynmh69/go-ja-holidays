package middleware

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestMain(m *testing.M) {
	setUp()
	defer tearDown()
	res := m.Run()
	os.Exit(res)
}

func TestSetMiddleware(t *testing.T) {
	e := echo.New()
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test OK",
			args: args{e: e},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMiddleware(tt.args.e)
		})
	}
}

func setUp() {

}

func tearDown() {

}
