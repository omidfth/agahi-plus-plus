package response

import "agahi-plus-plus/internal/model"

type UserBalancePayment struct {
	Balance      int                  `json:"balance"`
	UserPayments []*model.UserPayment `json:"user_payments"`
}
