package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type pricingRouter struct {
	pricingController controller.PricingController
}

func NewPricingRouter(pricingController controller.PricingController) Router {
	return &pricingRouter{pricingController: pricingController}
}

func (r pricingRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	post := router.Group("v1").Group("pricing")
	post.GET("", middlewares.Jwt(config), r.pricingController.List)
}
