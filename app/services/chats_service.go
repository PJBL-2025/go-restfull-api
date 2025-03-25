package services

import "restfull-api-pjbl-2025/model"

type ChatsService interface {
	GetAllChatsUser(UserId int, role string, queryId string) ([]*model.Chat, error)
	CreateChatsUser(chat *model.Chat, UserId int, role string, queryId string) error
	DeleteChatsUser(chatId int) error
}
