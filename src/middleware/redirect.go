package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Redirect() echo.MiddlewareFunc {
	registeredPaths := []string{
		"/manage/key",
		"/css",
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path
			c.Logger().Debug(c.Path())
			// 登録されているパス以外の場合はトップページにリダイレクト
			if !contains(registeredPaths, path) {
				return c.Redirect(http.StatusMovedPermanently, "/manage/key")
			}

			return next(c)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
