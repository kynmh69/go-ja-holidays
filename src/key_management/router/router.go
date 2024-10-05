package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/key_management/controller"
	"github.com/kynmh69/go-ja-holidays/logging"
)

func MakeRoute(r *gin.Engine) {
	controllers := []controller.Controller{
		controller.KeyManagement{ControllerName: "key"},
	}
	for _, v := range controllers {
		path := fmt.Sprintf("/manage/%s", v.GetControllerName())
		logging.GetLogger().Debugln(path)
		r.POST(fmt.Sprintf("%s/create", path), v.Create)
		r.GET(path, v.Retrieve)
		r.PUT(path, v.Update)
		r.POST(fmt.Sprintf("%s/delete", path), v.Delete)
	}
}
