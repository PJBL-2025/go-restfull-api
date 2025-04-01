package services

import "restfull-api-pjbl-2025/model"

type CheckoutsService interface {
	CreateOrderUser(userId int, order *model.Checkout) error
	CreatePaymentUser(orderId int, totalPrice int) (string, string, error)
	UpdateStatusPaymentUser(orderId int, status string) error
	GetPaymentUserById(orderID int) (*model.Checkout, error)
}
