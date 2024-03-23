package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerError(c echo.Context, err error) error {
	return c.HTML(http.StatusInternalServerError, "Server Error."+err.Error())
}
