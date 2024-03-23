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
	_, err := model.GetApiKeys()
	if err != nil {
		return util.ServerError(c, err)
	}
	return c.HTML(http.StatusOK, "")
}

func (k KeyManagement) Create(c echo.Context) error {
	_, err := model.CreateApiKey(c)
	if err != nil {
		return util.ServerError(c, err)
	}
	return c.HTML(http.StatusCreated, "")
}

func (k KeyManagement) Update(c echo.Context) error {
	return errors.New("not implemented")
}

func (k KeyManagement) Delete(c echo.Context) error {
	_, err := model.DeleteApiKey(c)
	if err != nil {
		return err
	}
	return c.HTML(http.StatusAccepted, "")
}

func (k KeyManagement) GetControllerName() string {
	return k.ControllerName
}
