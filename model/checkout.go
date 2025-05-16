package model

import "time"

type Checkout struct {
	Id         int        `json:"id" gorm:"column:id"`
	OrderId    int        `json:"order_id" gorm:"column:order_id" gorm:"unique"`
	Status     string     `json:"status" gorm:"column:status" default:"pending"`
	TotalPrice int        `json:"total_price" gorm:"column:total_price"`
	SnapToken  string     `json:"snap_token" gorm:"column:snap_token"`
	UserId     int        `json:"user_id" gorm:"column:user_id"`
	AddressId  int        `json:"address_id" gorm:"column:address_id"`
	CreatedAt  *time.Time `json:"created_at" gorm:"column:created_at"`
}
