package controller

import (
	"errors"
	"net/http"

	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/labstack/echo/v4"
)

const (
	TOP_PAGE_NAME = "top.html"
	TOP_PATH      = "/manage/key"
)

type KeyManagement struct {
	ControllerName string
}

func (k KeyManagement) Retrieve(c echo.Context) error {
	logger := c.Logger()
	apiKeys, _ := model.GetApiKeys()
	logger.Debug("APIKEYS", apiKeys)
	return c.Render(http.StatusOK, TOP_PAGE_NAME, apiKeys)
}

func (k KeyManagement) Create(c echo.Context) error {
	model.CreateApiKey(c)
	return c.Redirect(http.StatusFound, TOP_PATH)
}

func (k KeyManagement) Update(c echo.Context) error {
	return errors.New("not implemented")
}

func (k KeyManagement) Delete(c echo.Context) error {
	model.DeleteApiKey(c)
	return c.Redirect(http.StatusFound, TOP_PATH)
}

func (k KeyManagement) GetControllerName() string {
	return k.ControllerName
}

func NewKeyManagement(controllerName string) *KeyManagement {
	return &KeyManagement{ControllerName: controllerName}
}
