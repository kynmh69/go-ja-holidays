package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/api/router"
	"github.com/kynmh69/go-ja-holidays/database"
	"github.com/kynmh69/go-ja-holidays/logging"
)

func init() {
	logging.LoggerInitialize()
	database.ConnectDatabase()
}
func main() {
	logger := logging.GetLogger()
	r := initGin()
	router.MakeRoute(r)
	err := r.Run(":8080")
	if err != nil {
		// Runできなかったらエラーログを出力して終了
		logger.Error(err)
		return
	}
}

func initGin() *gin.Engine {
	return gin.Default()
}
