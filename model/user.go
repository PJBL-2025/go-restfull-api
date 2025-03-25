package model

type User struct {
	Id           int     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name         string  `json:"name"`
	Role         string  `json:"role"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
}
