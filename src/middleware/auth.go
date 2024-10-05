package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"github.com/kynmh69/go-ja-holidays/model"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logging.GetLogger()
		// get api key from request header
		key := c.GetHeader("X-API-KEY")
		if key == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "API key is required."})
			return
		}
		apiKey, err := model.GetApiKey(key)
		if err != nil {
			logger.Warnln("API key is invalid. ", key)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "API key is invalid."})
			return
		}
		logger.Debugln("API key is valid. ", apiKey.Id)
		c.Next()
	}
}
