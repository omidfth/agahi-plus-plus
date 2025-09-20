package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type promptRouter struct {
	promptController controller.PromptController
}

func NewPromptRouter(promptController controller.PromptController) Router {
	return &promptRouter{promptController: promptController}
}

func (r promptRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	prompt := router.Group("v1").Group("prompt")
	prompt.POST("generate/:service", middlewares.Jwt(config), r.promptController.Generate)
}
