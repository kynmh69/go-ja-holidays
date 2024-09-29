package logging

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// LoggerInitialize is a function to initialize logger
func LoggerInitialize() {
	if logger != nil {
		return
	}
	ginMode := gin.Mode()
	var i *zap.Logger
	switch ginMode {
	case gin.DebugMode:
		i, _ = zap.NewDevelopment()
	case gin.TestMode, gin.ReleaseMode:
		i, _ = zap.NewProduction()
	default:
		i, _ = zap.NewDevelopment()
	}
	logger = i.Sugar()
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
