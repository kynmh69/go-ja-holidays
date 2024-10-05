package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BadRequest struct {
	Message string `json:"message"`
}

func BadRequestJson(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, BadRequest{msg})
}
