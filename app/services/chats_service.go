package services

import (
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type ChatsService interface {
	GetAllChatsUser(UserId int, role string, queryId string) ([]*dto.ResponseChat, error)
	CreateChatsUser(chat *model.Chat, UserId int, role string, queryId string) error
	DeleteChatsUser(chatId int) error
}
