package util

import (
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
)

func SetUp() {
	logging.LoggerInitialize()
	database.ConnectDatabase()
}
