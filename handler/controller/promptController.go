package controller

import (
	"agahi-plus-plus/handler/request"
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type PromptController interface {
	Generate(ctx *gin.Context)
}
type promptController struct {
	promptService service.PromptService
	config        *helper.ServiceConfig
	logger        *zap.Logger
}

func NewPromptController(
	promptService service.PromptService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) PromptController {
	return &promptController{
		promptService: promptService,
		config:        config,
		logger:        logger,
	}
}

func (c promptController) Generate(ctx *gin.Context) {
	srv := ctx.Param("service")
	var req request.GeneratePromptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	post, balance, err := c.promptService.Generate(ctx, req.SelectedImagesUrls, srv)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "post": post, "balance": balance})
}
