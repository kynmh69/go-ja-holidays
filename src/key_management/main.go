package main

import (
	"os"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/key_management/router"
	"github.com/kynmh69/go-ja-holidays/key_management/template"
	"github.com/kynmh69/go-ja-holidays/middleware"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func init() {
	util.LoggerInitialize()
	database.ConnectDatabase()
}

func main() {
	e := echo.New()
	util.EchoLoggerInitialize(e)
	middleware.SetMiddleware(e)
	e.Use(mid.Static("./static"))
	logger := e.Logger

	t := template.NewTemplate("view/*.html")
	e.Renderer = t
	e.HTTPErrorHandler = util.CustomHTTPErrorHandler
	router.MakeRoute(e)
	wd, _ := os.Getwd()
	logger.Debug(wd)
	logger.Fatal(e.Start(":80"))
}
