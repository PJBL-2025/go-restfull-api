package dto

type RequestUpdateCheckout struct {
	Status     string `json:"status"`
	CheckoutId int    `json:"checkout_id"`
}
