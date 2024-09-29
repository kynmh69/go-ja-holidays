package logging

import (
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.SugaredLogger
	}{
		// TODO: Add test cases.
		{
			name: "TestGetLogger",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerInitialize(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "testLoggerInitialize",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerInitialize()
		})
	}
}
