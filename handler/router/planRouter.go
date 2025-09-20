package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type planRouter struct {
	planController controller.PlanController
}

func NewPlanRouter(planController controller.PlanController) Router {
	return &planRouter{planController: planController}
}

func (r planRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	plan := router.Group("v1").Group("plan")
	plan.GET("", middlewares.Jwt(config), r.planController.List)
}
