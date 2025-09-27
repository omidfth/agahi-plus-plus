package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"

	"github.com/gin-gonic/gin"
)

type postRouter struct {
	postController controller.PostController
}

func NewPostRouter(postController controller.PostController) Router {
	return &postRouter{postController: postController}
}

func (r postRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	post := router.Group("v1").Group("post")
	post.GET(":service", middlewares.Jwt(config), r.postController.Get)
}
