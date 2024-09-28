package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Retrieve(r *gin.Context) error
	Create(r *gin.Context) error
	Update(r *gin.Context) error
	Delete(r *gin.Context) error
	GetControllerName() string
}
