package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	Retrieve(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetControllerName() string
}
