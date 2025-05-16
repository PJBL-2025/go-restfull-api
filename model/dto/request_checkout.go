package dto

import "restfull-api-pjbl-2025/model"

type RequestCheckout struct {
	AddressId       int                     `json:"address_id"`
	TotalPrice      int                     `json:"total_price"`
	ProductCheckout []model.ProductCheckout `json:"product_checkout"`
}
