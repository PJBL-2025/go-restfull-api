package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"restfull-api-pjbl-2025/app/controllers"
	"restfull-api-pjbl-2025/app/middleware"
	"restfull-api-pjbl-2025/ws"
)

func SetUpRoutes(app *fiber.App, chatController controllers.ChatsController, checkoutController controllers.CheckoutsController, productController controllers.ProductController) {
	api := app.Group("/api", middleware.AuthMiddleware())
	api.Post("/chat/user", chatController.CreateChatsUser)
	api.Get("/chat/user", chatController.GetAllChatsUser)
	api.Post("/order", checkoutController.CreateOrderProduct)
	api.Patch("/order/update", checkoutController.UpdateStatusCheckout)
	api.Get("/order/:status", checkoutController.GetCheckout)
	api.Get("/order/detail/:id", checkoutController.GetDetailProductCheckout)
	api.Get("/order", checkoutController.GetCheckoutsAdmin)

	api.Patch("/order/delivery/:id", checkoutController.SetDelivery)
	api.Post("/order/delivery/status", checkoutController.SetStatusDelivery)
	
	api.Post("/product", productController.AddProductCheckout)
	api.Patch("/product/:id", productController.UpdateProduct)

	app.Get("/ws/chat", websocket.New(ws.WebSocketHub.HandleConnections))
}
