package repositories

import "restfull-api-pjbl-2025/model"

type CheckoutsRepository interface {
	CreatePayment(order *model.Checkout) error
	UpdateStatusPayment(orderId int, status string) error
	GetPaymentById(orderID int) (*model.Checkout, error)
}
