package middleware

import (
	"os"

	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func getLogFormat() string {
	format, ok := os.LookupEnv("LOG_FORMAT")
	if !ok {
		format = "${time_rfc3339_nano} [${method}] ${uri} ${status} ${id}\n"
	}
	return format
}

func confLogFormat() mid.LoggerConfig {
	return mid.LoggerConfig{Format: getLogFormat(), Output: util.InitWriter()}
}

func setLogger() echo.MiddlewareFunc {
	return mid.LoggerWithConfig(confLogFormat())
}
