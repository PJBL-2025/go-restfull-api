package model

import "time"

type StatusDelivery struct {
	Id           int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	DeliveriesId int       `json:"deliveries_id" gorm:"column:deliveries_id"`
}
