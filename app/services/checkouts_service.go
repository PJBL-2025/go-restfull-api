package services

import (
	"restfull-api-pjbl-2025/model/dto"
)

type CheckoutsService interface {
	CreateOrderUser(userId int, checkout *map[string]interface{}) (string, string, error)
	CreateOrderCustom(productCheckout []interface{}, checkoutId int) error
	UpdateStatusCheckout(checkout *dto.RequestUpdateCheckout) error
	SetDelivery(delivery *dto.SetDelivery, deliveryId int) error
	SetStatusDelivery(status map[string]interface{}) error
	GetCheckout(param string, userId int) ([]map[string]interface{}, error)
	GetDetailCheckoutProduct(productCheckoutId int) (map[string]interface{}, error)
}
