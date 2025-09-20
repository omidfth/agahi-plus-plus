package repository

import "agahi-plus-plus/internal/model"

type UserPaymentRepository interface {
	Create(*model.UserPayment) error
	List(userID uint) ([]*model.UserPayment, error)
}
