package helper

import "github.com/gofiber/fiber/v2"

func SuccessResponse(ctx *fiber.Ctx, data any, message string) error {
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
		"message": message,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, message string, err error) error {
	return ctx.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   err,
	})
}
