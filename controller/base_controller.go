package controller

import (
	"github.com/gin-gonic/gin"
	"todos-rest/logger"
)

type BaseControllerInterface interface {
}

type BaseController struct {
	logger *logger.Logger
}

func NewBaseController(logger *logger.Logger) *BaseController {
	return &BaseController{logger: logger}
}

func (c BaseController) response(context *gin.Context, obj interface{}, code int) {
	context.JSON(code, obj)
}
