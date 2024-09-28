package router

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestMakeRoute(t *testing.T) {
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeRoute(tt.args.r)
		})
	}
}
