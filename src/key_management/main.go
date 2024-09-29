package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/key_management/router"
	"github.com/kynmh69/go-ja-holidays/logging"
	"os"
)

func init() {
	logging.LoggerInitialize()
	database.ConnectDatabase()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*.html")
	r.Static("/css", "./static/css")
	logger := logging.GetLogger()

	//e.HTTPErrorHandler = util.CustomHTTPErrorHandler
	router.MakeRoute(r)
	wd, _ := os.Getwd()
	logger.Debug(wd)
	logger.Fatal(r.Run(":8080"))
}
