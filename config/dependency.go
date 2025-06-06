package config

import (
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/app/controllers"
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/app/services"
)

func DependencyInjection(db *gorm.DB, snapClient *snap.Client) (*controllers.ChatsControllerImpl, *controllers.CheckoutsControllerImpl, *controllers.ProductControllerImpl) {
	chatRepository := repositories.NewChatRepositoryImpl(db)
	chatService := services.NewChatServiceImpl(chatRepository)
	chatController := controllers.NewChatsControllerImpl(chatService)

	checkoutRepository := repositories.NewCheckoutRepositoryImpl(db)
	checkoutService := services.NewCheckoutServiceImpl(checkoutRepository, snapClient)
	checkoutController := controllers.NewCheckoutsControllerImpl(checkoutService)

	productRepository := repositories.NewProductRepositoryImpl(db)
	productService := services.NewProductServiceImpl(productRepository)
	productController := controllers.NewProductControllerImpl(productService)

	return chatController, checkoutController, productController
}
