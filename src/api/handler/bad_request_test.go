package handler

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func testHandler(c *gin.Context) {
	BadRequestJson(c, "TestBadRequestJson")
}

func TestBadRequestJson(t *testing.T) {
	r := gin.Default()
	r.GET("/", testHandler)
	w := httptest.NewRecorder()
	ctx := gin.CreateTestContextOnly(w, r)

	request := httptest.NewRequest("GET", "/", nil)
	ctx.Request = request

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestBadRequestJson",
			args: args{
				c: ctx,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testHandler(tt.args.c)
		})
	}
}
