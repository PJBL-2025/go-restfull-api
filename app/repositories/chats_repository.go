package repositories

import (
	"restfull-api-pjbl-2025/model"
)

type ChatsRepository interface {
	GetAllChats(UserId int, AdminId int) ([]*model.Chat, error)
	GetAdminId() (int, error)
	CreateChats(chat *model.Chat) error
	DeleteChats(chatId int) error
}
