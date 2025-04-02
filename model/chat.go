package model

type Chat struct {
	Id          int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId      int    `json:"user_id"`
	AdminId     int    `json:"admin_id"`
	Message     string `json:"message"`
	RoleMessage string `json:"role_message"`
	ProductId   *int   `json:"product_id" gorm:"column:product_id"`
}
