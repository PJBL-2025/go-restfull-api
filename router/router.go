package router

import (
	"github.com/gofiber/fiber/v2"
	"restfull-api-pjbl-2025/app/controllers"
	"restfull-api-pjbl-2025/app/middleware"
)

func SetUpRoutes(app *fiber.App, chatController controllers.ChatsController) {
	api := app.Group("/api", middleware.AuthMiddleware())
	api.Post("/chat/user", chatController.CreateChatsUser)
	api.Get("/chat/user", chatController.GetAllChatsUser)
}
