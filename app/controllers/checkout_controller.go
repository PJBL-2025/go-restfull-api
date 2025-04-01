package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type CheckoutsController interface {
	CreatePaymentUser(ctx *fiber.Ctx) error
	UpdateStatusPaymentUser(ctx *fiber.Ctx) error
	GetPaymentUserById(ctx *fiber.Ctx) error
}
