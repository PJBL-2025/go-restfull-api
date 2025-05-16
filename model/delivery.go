package model

import "time"

type Delivery struct {
	Id            int        `json:"id" gorm:"column:id"`
	SendStartTime *time.Time `json:"send_start_time" gorm:"column:send_start_time"`
	SendEndTime   *time.Time `json:"send_end_time" gorm:"column:send_end_time"`
	CheckoutId    int        `json:"checkout_id" gorm:"column:checkout_id"`
}
