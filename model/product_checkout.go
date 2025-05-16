package model

type ProductCheckout struct {
	Id         int    `json:"id" gorm:"primary_key" gorm:"column:id" gorm:"unique"`
	Quantity   int    `json:"quantity" gorm:"column:quantity"`
	Size       string `json:"size" gorm:"column:size"`
	Type       string `json:"type" gorm:"column:type"`
	Color      string `json:"color" gorm:"column:color"`
	Price      int    `json:"price" gorm:"column:price"`
	CheckoutId int    `json:"checkout_id" gorm:"column:checkout_id"`
	ProductId  *int   `json:"product_id" gorm:"column:product_id"`
}
