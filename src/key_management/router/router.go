package router

import (
	"fmt"

	"github.com/kynmh69/go-ja-holidays/key_management/controller"
	"github.com/labstack/echo/v4"
)

func MakeRoute(e *echo.Echo) {
	controllers := []controller.Controller{
		controller.KeyManagement{ControllerName: "key"},
	}
	for _, v := range controllers {
		path := fmt.Sprintf("/manage/%s", v.GetControllerName())
		e.POST(path, v.Create)
		e.GET(path, v.Retrieve)
		e.PUT(path, v.Update)
		e.DELETE(path, v.Delete)
	}
}
