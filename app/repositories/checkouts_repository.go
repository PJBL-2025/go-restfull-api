package repositories

import (
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type CheckoutsRepository interface {
	CreateCheckout(userId int, totalPrice int, addressId int, orderId int) (int, error)
	InsertSnapToken(checkoutId int, snap string) error
	CreateOrder(order *model.ProductCheckout) (int, error)
	CreateProductCustom(custom map[string]interface{}, productId int) error
	CreateDelivery(checkoutId int) (int, error)
	CreateStatusDelivery(status string, deliveryId int) error
	UpdateStatusCheckout(checkout *dto.RequestUpdateCheckout) error
	SetDelivery(delivery *dto.SetDelivery, deliveryId int) error
	SetStatusDelivery(status map[string]interface{}) error
	GetCheckoutPending(param string, userId int) ([]map[string]interface{}, error)
	GetCheckoutNotPending(param string, userId int) ([]map[string]interface{}, error)
	GetDetailProductCheckout(productCheckoutId int) ([]map[string]interface{}, error)
	GetDetailProductCheckoutAdmin(productCheckoutId int) ([]map[string]interface{}, error)
	GetCheckoutsAdmin() ([]map[string]interface{}, error)
}
