package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BadRequest struct {
	Message string `json:"message"`
}

func BadRequestJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, BadRequest{msg})
}
