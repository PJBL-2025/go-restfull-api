package repositories

import (
	"restfull-api-pjbl-2025/model"
)

type ChatsRepository interface {
	GetIdentifyChats(UserId int, AdminId int) ([]map[string]interface{}, error)
	GetChats(productId int) (map[string]interface{}, error)
	GetAdminId() (int, error)
	CreateChats(chat *model.Chat) error
	DeleteChats(chatId int) error
}
