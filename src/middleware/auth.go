package middleware

import (
	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Auth() echo.MiddlewareFunc {
	return mid.KeyAuthWithConfig(mid.KeyAuthConfig{
		KeyLookup: "header:X-API-KEY",
		Validator: func(auth string, c echo.Context) (bool, error) {
			apiKey, err := model.GetApiKey(auth)
			if err != nil {
				return false, err
			}
			return auth == apiKey.Key, nil
		},
	})
}
