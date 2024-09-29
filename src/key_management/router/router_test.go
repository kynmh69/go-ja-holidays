package router

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMakeRoute(t *testing.T) {
	r := gin.Default()
	type args struct {
		e *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test ok",
			args: args{e: r},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeRoute(tt.args.e)
		})
	}
}
