package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type paymentRouter struct {
	paymentController controller.PaymentController
}

func NewPaymentRouter(paymentController controller.PaymentController) Router {
	return &paymentRouter{paymentController: paymentController}
}

func (r paymentRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	payment := router.Group("v1").Group("payment")
	payment.GET("purchase", middlewares.Jwt(config), r.paymentController.NewPayment)
	payment.GET("verify", r.paymentController.Verify)
}
