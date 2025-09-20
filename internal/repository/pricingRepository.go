package repository

import "agahi-plus-plus/internal/model"

type PricingRepository interface {
	ListWithFirstDiscount(serviceID uint) ([]*model.PricingLogic, error)
	ListWithOutFirstDiscount(serviceID uint) ([]*model.PricingLogic, error)
	Get(id int) (*model.PricingLogic, error)
	FindByPrice(price uint) (*model.PricingLogic, error)
}
