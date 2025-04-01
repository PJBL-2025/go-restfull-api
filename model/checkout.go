package model

import "time"

type Checkout struct {
	Id            int        `json:"id" gorm:"column:id"`
	Status        string     `json:"status" gorm:"column:status" default:"pending"`
	Quantity      int        `json:"quantity" gorm:"column:quantity"`
	TotalPrice    int        `json:"total_price" gorm:"column:total_price"`
	PaymentMethod string     `json:"payment_method" gorm:"column:payment_method"`
	UserId        int        `json:"user_id" gorm:"column:user_id"`
	AddressId     int        `json:"address_id" gorm:"column:address_id"`
	ProductId     int        `json:"product_id" gorm:"column:product_id"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at"`
}
