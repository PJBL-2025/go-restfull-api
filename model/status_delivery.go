package model

type StatusDelivery struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Status       string `json:"status"`
	DeliveriesId int    `json:"deliveries_id" gorm:"column:deliveries_id"`
}
