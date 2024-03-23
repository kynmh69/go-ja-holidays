package controller

import (
	"errors"
	"net/http"

	"github.com/kynmh69/go-ja-holidays/model"
	"github.com/kynmh69/go-ja-holidays/util"
	"github.com/labstack/echo/v4"
)

type KeyManagement struct {
	ControllerName string
}

func (k KeyManagement) Retrieve(c echo.Context) error {
	logger := c.Logger()
	apiKeys, err := model.GetApiKeys()
	if err != nil {
		return util.ServerError(c, err)
	}
	logger.Debug(apiKeys)
	return c.Render(http.StatusOK, "top", apiKeys)
}

func (k KeyManagement) Create(c echo.Context) error {
	apiKeys, err := model.CreateApiKey(c)
	if err != nil {
		return util.ServerError(c, err)
	}
	return c.Render(http.StatusCreated, "top", apiKeys)
}

func (k KeyManagement) Update(c echo.Context) error {
	return errors.New("not implemented")
}

func (k KeyManagement) Delete(c echo.Context) error {
	apikeys, err := model.DeleteApiKey(c)
	if err != nil {
		return err
	}
	return c.Render(http.StatusAccepted, "top", apikeys)
}

func (k KeyManagement) GetControllerName() string {
	return k.ControllerName
}

func NewKeyManagement(controllerName string) *KeyManagement {
	return &KeyManagement{ControllerName: controllerName}
}
