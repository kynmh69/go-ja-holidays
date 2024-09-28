package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/api/handler"
)

func MakeRoute(r *gin.Engine) {
	r.GET("/holidays", handler.GetHolidays)
	r.GET("/holidays/:day", handler.IsHoliday)
	r.GET("/holidays/count", handler.CountHolidays)
}
