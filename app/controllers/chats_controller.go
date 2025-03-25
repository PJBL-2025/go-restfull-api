package controllers

import "github.com/gofiber/fiber/v2"

type ChatsController interface {
	GetAllChatsUser(ctx *fiber.Ctx) error
	CreateChatsUser(ctx *fiber.Ctx) error
	DeleteChatUser(ctx *fiber.Ctx) error
}
