package service

import (
	"agahi-plus-plus/internal/helper"
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlanService interface {
	List(ctx *gin.Context, serviceID uint) ([]*model.Plan, error)
	Get(id int) (*model.Plan, error)
	FindByPrice(price uint) (*model.Plan, error)
}

type planService struct {
	planRepo    repository.PlanRepository
	userService UserService
	config      *helper.ServiceConfig
	logger      *zap.Logger
}

func NewPlanService(
	planRepo repository.PlanRepository,
	userService UserService,
	config *helper.ServiceConfig,
	logger *zap.Logger,
) PlanService {
	return &planService{
		planRepo:    planRepo,
		userService: userService,
		config:      config,
		logger:      logger,
	}
}

func (p planService) List(ctx *gin.Context, serviceID uint) ([]*model.Plan, error) {
	user, err := p.userService.GetUserWithContext(ctx)
	if err != nil {
		return nil, err
	}

	if user.IsUse == true {
		return p.planRepo.ListWithOutFirstDiscount(serviceID)
	}

	return p.planRepo.ListWithFirstDiscount(serviceID)
}

func (p planService) Get(id int) (*model.Plan, error) {
	return p.planRepo.Get(id)
}

func (p planService) FindByPrice(price uint) (*model.Plan, error) {
	return p.planRepo.FindByPrice(price)
}
