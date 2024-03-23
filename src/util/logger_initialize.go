package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	echoLog "github.com/labstack/gommon/log"
)

func LoggerInitialize() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[App] ")
}

func GetLoggerLevel() echoLog.Lvl {
	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		level = "info"
	}
	level = strings.ToLower(level)
	levelNo := echoLog.INFO
	switch level {
	case "debug":
		levelNo = echoLog.DEBUG
	case "info":
		levelNo = echoLog.INFO
	case "warn":
		levelNo = echoLog.WARN
	case "warning":
		levelNo = echoLog.WARN
	case "error":
		levelNo = echoLog.ERROR
	default:
		levelNo = echoLog.OFF
	}
	return levelNo
}

func EchoLoggerInitialize(e *echo.Echo) {
	logger := e.Logger
	logger.SetPrefix("[APP]")
	logger.SetLevel(GetLoggerLevel())
	logger.SetOutput(InitWriter())
	logger.SetHeader("${prefix} ${time_rfc3339_nano} [${level}] ${path} Line.${line} ")
}

func InitWriter() io.Writer {
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
		log.Fatalln("Not open log file", err)
	}
	return file
}
