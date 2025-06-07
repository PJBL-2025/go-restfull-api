package dto

type Category struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Category string `gorm:"not null" json:"category"`
}
