package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kynmh69/go-ja-holidays/logging"
	"net/http"

	"github.com/kynmh69/go-ja-holidays/model"
)

const (
	TopPageName = "top.html"
	TopPath     = "/manage/key"
)

type KeyManagement struct {
	ControllerName string
}

func (k KeyManagement) Retrieve(c *gin.Context) {
	logger := logging.GetLogger()
	apiKeys, _ := model.GetApiKeys()
	logger.Debug("APIKEY", apiKeys)
	c.HTML(http.StatusOK, TopPageName, apiKeys)
}

func (k KeyManagement) Create(c *gin.Context) {
	err := model.CreateApiKey()
	if err != nil {
		c.HTML(500, "error", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, TopPath)
}

func (k KeyManagement) Update(c *gin.Context) {
	c.HTML(500, "update", nil)
}

func (k KeyManagement) Delete(c *gin.Context) {
	err := model.DeleteApiKey()
	if err != nil {
		c.HTML(500, "error", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, TopPath)
}

func (k KeyManagement) GetControllerName() string {
	return k.ControllerName
}

func NewKeyManagement(controllerName string) *KeyManagement {
	return &KeyManagement{ControllerName: controllerName}
}
