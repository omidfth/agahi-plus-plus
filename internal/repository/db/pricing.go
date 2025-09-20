package db

import (
	"agahi-plus-plus/internal/model"
	"agahi-plus-plus/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewPlanDB(db *gorm.DB, logger *zap.Logger) repository.PlanRepository {
	return &planDB{
		db:     db,
		logger: logger,
	}
}

type planDB struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (p planDB) ListWithFirstDiscount(serviceID uint) ([]*model.Plan, error) {
	var plan []*model.Plan

	err := p.db.Where("id != ? AND service_id = ?", 1, serviceID).Order("price ASC").Find(&plan).Error
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p planDB) ListWithOutFirstDiscount(serviceID uint) ([]*model.Plan, error) {
	var plan []*model.Plan

	err := p.db.Where("id != ? AND service_id = ?", 4, serviceID).Order("price ASC").Find(&plan).Error
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (p planDB) Get(id int) (*model.Plan, error) {
	var price model.Plan
	err := p.db.Where("id = ?", id).Find(&price).Error
	if err != nil {
		return nil, err
	}

	return &price, nil
}

func (p planDB) FindByPrice(priceAmount uint) (*model.Plan, error) {
	var price model.Plan
	err := p.db.Where("price = ?", priceAmount).Find(&price).Error
	if err != nil {
		return nil, err
	}

	return &price, nil
}
