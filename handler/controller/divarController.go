package controller

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type DivarController interface {
	EditPost(ctx *gin.Context)
}
type divarController struct {
	divarService service.DivarService
	config       *helper.ServiceConfig
	logger       *zap.Logger
}

func NewDivarController(
	divarService service.DivarService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) DivarController {
	return &divarController{
		divarService: divarService,
		config:       config,
		logger:       logger,
	}
}

func (c divarController) EditPost(ctx *gin.Context) {
	postToken := ctx.Param("post-token")
	//srv := ctx.Param("service")

	if postToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "empty post token"})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
