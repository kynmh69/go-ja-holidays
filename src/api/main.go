package main

import (
	"fmt"
	"io"
	originLog "log"
	"os"
	"path/filepath"
	"time"

	"github.com/kynmh69/go-ja-holidays/api/router"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func init() {
	util.LoggerInitialize()
	database.ConnectDatabase()
}
func main() {
	e := initEcho()
	loggerInitialize(e)
	logger := e.Logger
	router.MakeRoute(e)
	logger.Fatal(e.Start(":80"))
}

func initEcho() *echo.Echo {
	return echo.New()
}

func loggerInitialize(e *echo.Echo) {
	logger := e.Logger
	logger.SetPrefix("[APP]")
	logger.SetLevel(log.DEBUG)
	logger.SetOutput(initWriter())
}

func initWriter() io.Writer {
	logDir, ok := os.LookupEnv("LOG_DIR")
	if !ok {
		logDir = "./log/"
	}
	os.Mkdir(logDir, 0755)

	return io.MultiWriter(os.Stdout, createFile(logDir))
}

func createFile(logDir string) *os.File {
	nowStr := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("app_%s.log", nowStr)
	logFile := filepath.Join(logDir, logFileName)
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		originLog.Fatalln("Not open log file", err)
	}
	return file
}
