package main

import (
	"github.com/kynmh69/go-ja-holidays/api/router"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/middleware"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

func init() {
	util.LoggerInitialize()
	database.ConnectDatabase()
}
func main() {
	e := initEcho()
	middleware.SetMiddleware(e)
	logger := e.Logger
	router.MakeRoute(e)
	logger.Fatal(e.Start(":80"))
}

func initEcho() *echo.Echo {
	return echo.New()
}
