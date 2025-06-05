package model

import "time"

type Product struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Price       int
	Quantity    int
	Description string
	Weight      int
	CreatedAt   time.Time
}
