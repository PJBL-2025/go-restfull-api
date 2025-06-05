package services

import (
	"restfull-api-pjbl-2025/model/dto"
	"time"
)

type CheckoutsService interface {
	CreateOrderUser(userId int, checkout *map[string]interface{}) (string, string, error)
	CreateOrderCustom(productCheckout []interface{}, checkoutId int) error
	UpdateStatusCheckout(checkout *dto.RequestUpdateCheckout) error
	SetDelivery(delivery *dto.SetDelivery, deliveryId int) error
	SetStatusDelivery(status string, createdAt time.Time, deliveryId int) error
	GetCheckout(param string, userId int) ([]map[string]interface{}, error)
	GetDetailCheckoutProduct(productCheckoutId int) (map[string]interface{}, error)
	GetCheckoutsAdmin() ([]map[string]interface{}, error)
	AddProduct(product *dto.RequestProduct) error
}
