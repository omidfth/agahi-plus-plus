package router

import (
	"agahi-plus-plus/internal/helper"
	"github.com/gin-gonic/gin"
)

type Router interface {
	HandleRoutes(router *gin.Engine, config *helper.ServiceConfig)
}
