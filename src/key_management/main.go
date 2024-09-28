package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"os"

	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/key_management/router"
	"github.com/kynmh69/go-ja-holidays/key_management/template"
)

func init() {
	logging.LoggerInitialize()
	database.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	logger := logging.GetLogger()

	r.HTMLRender = template.Render("view/*.html")
	//e.HTTPErrorHandler = util.CustomHTTPErrorHandler
	router.MakeRoute(r)
	wd, _ := os.Getwd()
	logger.Debug(wd)
	logger.Fatal(e.Start(":80"))
}
