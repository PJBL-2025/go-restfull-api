package config

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/app/controllers"
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/app/services"
)

func DependencyInjection(db *gorm.DB) *controllers.ChatsControllerImpl {
	chatRepository := repositories.NewChatRepositoryImpl(db)
	chatService := services.NewChatServiceImpl(chatRepository)
	chatController := controllers.NewChatsControllerImpl(chatService)

	return chatController
}
