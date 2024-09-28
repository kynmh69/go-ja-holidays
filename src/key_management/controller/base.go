package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Retrieve(r *gin.Context)
	Create(r *gin.Context)
	Update(r *gin.Context)
	Delete(r *gin.Context)
	GetControllerName() string
}
