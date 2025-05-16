package dto

import "time"

type SetDelivery struct {
	SendStartTime time.Time `json:"send_start_time" gorm:"column:send_start_time"`
	SendEndTime   time.Time `json:"send_end_time" gorm:"column:send_end_time"`
}
