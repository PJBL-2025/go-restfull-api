package dto

import "time"

type RequestProduct struct {
	Id          int       `gorm:"primary_key;auto_increment"`
	Image       string    `json:"image"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Quantity    int       `json:"quantity"`
	Weight      int       `json:"weight"`
	Size        []int     `json:"size"`
	CategoryId  []int     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
