package router

import (
	"agahi-plus-plus/handler/controller"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type divarRouter struct {
	divarController controller.DivarController
}

func NewDivarRouter(divarController controller.DivarController) Router {
	return &divarRouter{divarController: divarController}
}

func (r divarRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	//user := router.Group("v1").Group("divar")
}
