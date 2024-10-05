package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/api/handler"
	"github.com/kynmh69/go-ja-holidays/middleware"
)

func MakeRoute(r *gin.Engine) {
	r.Use(middleware.Auth())
	r.GET("/holidays", handler.GetHolidays)
	r.GET("/holidays/:date", handler.IsHoliday)
	r.GET("/holidays/count", handler.CountHolidays)
}
