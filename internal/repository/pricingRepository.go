package repository

import "agahi-plus-plus/internal/model"

type PlanRepository interface {
	ListWithFirstDiscount(serviceID uint) ([]*model.Plan, error)
	ListWithOutFirstDiscount(serviceID uint) ([]*model.Plan, error)
	Get(id int) (*model.Plan, error)
	FindByPrice(price uint) (*model.Plan, error)
}
