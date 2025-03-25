package repositories

import (
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type ChatsRepository interface {
	GetAllChats(UserId int, AdminId int) ([]*dto.ResponseChat, error)
	GetAdminId() (int, error)
	CreateChats(chat *model.Chat) error
	DeleteChats(chatId int) error
}
