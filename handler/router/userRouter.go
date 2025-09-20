package router

import (
	"agahi-plus-plus/handler/controller"
	middlewares "agahi-plus-plus/handler/middleware"
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type userRouter struct {
	userController controller.UserController
}

func NewUserRouter(userController controller.UserController) Router {
	return &userRouter{userController: userController}
}

func (r userRouter) HandleRoutes(router *gin.Engine, config *helper.ServiceConfig) {
	user := router.Group("v1").Group("user")
	user.GET("login/divar/:service", r.userController.LoginWithDivar)
	user.GET("login/divar/call/:service", r.userController.CallOAuth)
	user.GET("oauth/:service", r.userController.OAuth)
	user.GET("balance", middlewares.Jwt(config), r.userController.GetBalance)
	user.POST("register", r.userController.Register)
	user.POST("login", r.userController.Login)
	user.GET("has-balance", middlewares.Jwt(config), r.userController.HasBalance)
	user.GET("login/check", middlewares.Jwt(config), r.userController.CheckLogin)
	user.GET("posts/:service", r.userController.GetPosts)
	user.GET("ads/entry/:service", r.userController.AdsEntry)
	user.GET("ads/oauth/:service", r.userController.AdsOAuth)
}
