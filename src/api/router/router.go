package router

import (
	"github.com/kynmh69/go-ja-holidays/api/controller"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func MakeRoute(e *echo.Echo) {
	e.GET("/holidays", controller.GetHolidays)
	e.GET("/holidays/:day", controller.IsHoliday)
	e.GET("/holidays/count", controller.CountHolidays)
}

func SetPrometheusHandler(e *echo.Echo) {
	e.GET("/metrics", echoprometheus.NewHandler())
}
