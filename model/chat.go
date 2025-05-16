package model

type Chat struct {
	Id         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int    `json:"user_id"`
	AdminId    int    `json:"admin_id"`
	ChatUserId int    `json:"chat_user_id"`
	Message    string `json:"message"`
}
