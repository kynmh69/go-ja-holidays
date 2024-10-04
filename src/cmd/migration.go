package main

import (
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
)

func init() {
	// logger initialize
	logging.LoggerInitialize()
	// connect database
	database.ConnectDatabase()
}

func main() {
	// get logger
	logger := logging.GetLogger()
	// migration db
	db := database.GetDbConnection()
	// migrate db
	if err := db.AutoMigrate(
		&model.ApiKey{},
		&model.HolidayData{},
	); err != nil {
		// if failed to migrate db, panic
		logger.Panicln("failed to migrate db", err)
	}
	// if success to migrate db, log
	logger.Infoln("migrate db success")
}
