package model

import "time"

type Chat struct {
	Id        int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int       `json:"user_id"`
	AdminId   int       `json:"admin_id"`
	Role      string    `json:"role"`
	Message   string    `json:"message"`
	ProductId *int      `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}
