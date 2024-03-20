package middleware

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func SetPrometheus() echo.MiddlewareFunc{
	return echoprometheus.NewMiddleware("app")
}
