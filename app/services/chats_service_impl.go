package services

import (
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
	"strconv"
)

type ChatServiceImpl struct {
	chatRepository repositories.ChatsRepository
}

func NewChatServiceImpl(chatRepository repositories.ChatsRepository) *ChatServiceImpl {
	return &ChatServiceImpl{chatRepository: chatRepository}
}

func (service *ChatServiceImpl) GetAllChatsUser(UserId int, role string, queryId string) ([]*dto.ResponseChat, error) {
	AdminId, err := service.chatRepository.GetAdminId()
	if err != nil {
		return nil, err
	}

	queryIdInt, err := strconv.Atoi(queryId)
	if err != nil {
		return nil, err
	}

	if role == "admin" {
		UserId = queryIdInt
	}

	chats, err := service.chatRepository.GetAllChats(UserId, AdminId)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (service *ChatServiceImpl) CreateChatsUser(chat *model.Chat, UserId int, role string, queryId string) error {
	AdminId, err := service.chatRepository.GetAdminId()
	if err != nil {
		return err
	}

	queryIdInt, err := strconv.Atoi(queryId)

	if err != nil {
		return err
	}

	if role == "user" {
		chat.AdminId = AdminId
		chat.UserId = UserId
	} else {
		chat.AdminId = UserId
		chat.UserId = queryIdInt
	}

	err = service.chatRepository.CreateChats(chat)
	if err != nil {
		return err
	}

	return nil
}

func (service *ChatServiceImpl) DeleteChatsUser(chatId int) error {
	return service.chatRepository.DeleteChats(chatId)
}
