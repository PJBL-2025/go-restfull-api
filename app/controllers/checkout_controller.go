package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type CheckoutsController interface {
	CreateOrderProduct(ctx *fiber.Ctx) error
	UpdateStatusCheckout(ctx *fiber.Ctx) error
	SetDelivery(ctx *fiber.Ctx) error
	SetStatusDelivery(ctx *fiber.Ctx) error
	GetCheckout(ctx *fiber.Ctx) error
	GetDetailProductCheckout(ctx *fiber.Ctx) error
	GetDetailProductCheckoutAdmin(ctx *fiber.Ctx) error
	GetCheckoutsAdmin(ctx *fiber.Ctx) error
}
