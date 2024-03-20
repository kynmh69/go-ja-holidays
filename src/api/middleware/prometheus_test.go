package middleware

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestSetPrometheus(t *testing.T) {
	tests := []struct {
		name string
		want echo.MiddlewareFunc
	}{
		{
			name: "test ok",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
