package router

import (
	"github.com/kynmh69/go-ja-holidays/api/controller"
	"github.com/labstack/echo/v4"
)

func MakeRoute(e *echo.Echo) {
	e.GET("/holidays", controller.GetHolidays)
}