package controller

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type PlanController interface {
	List(ctx *gin.Context)
}
type planController struct {
	planService service.PlanService
	config      *helper.ServiceConfig
	logger      *zap.Logger
}

func (c planController) List(ctx *gin.Context) {
	serviceString := ctx.Query("service_id")
	serviceID, _ := strconv.Atoi(serviceString)

	post, err := c.planService.List(ctx, uint(serviceID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": post})
}

func NewPlanController(
	planService service.PlanService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) PlanController {
	return &planController{
		planService: planService,
		config:      config,
		logger:      logger,
	}
}
