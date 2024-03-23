package util

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ServerError(c echo.Context, err error) error {
	return c.HTML(http.StatusInternalServerError, "Server Error."+err.Error())
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	logger := c.Logger()
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	logger.Error("Error: ", code)
	c.Logger().Error(err)
	errorPage := fmt.Sprintf("view/error/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
}
