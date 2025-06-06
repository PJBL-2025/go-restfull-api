package controllers

import "github.com/gofiber/fiber/v2"

type ProductController interface {
	AddProductCheckout(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
}
