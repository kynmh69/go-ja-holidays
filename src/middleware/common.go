package middleware

import (
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func SetMiddleware(e *echo.Echo) {
	e.Use(setLogger())
	e.Use(mid.CORS())
	e.Use(mid.CSRF())
	e.Use(SetPrometheus())
	e.Use(mid.Recover())
}
