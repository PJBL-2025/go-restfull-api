package services

import (
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/model"
	"strconv"
	"time"
)

type ChatServiceImpl struct {
	chatRepository repositories.ChatsRepository
}

func NewChatServiceImpl(chatRepository repositories.ChatsRepository) *ChatServiceImpl {
	return &ChatServiceImpl{chatRepository: chatRepository}
}

func (service *ChatServiceImpl) GetAllChatsUser(UserId int, role string, queryId string) (map[string]interface{}, error) {
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

	chats, err := service.chatRepository.GetIdentifyChats(UserId, AdminId)
	if err != nil {
		return nil, err
	}

	var userName string
	if len(chats) > 0 {
		// Ambil name dari chat pertama
		if name, ok := chats[0]["username"].(string); ok {
			userName = name
		}
	}

	// Buang 'name' dari setiap chat karena sudah di luar
	for i := range chats {
		delete(chats[i], "name")
	}

	// Jika ada product_id, ambil detail product dan masukkan ke chat["product"]
	for i, chat := range chats {
		if pid, ok := chat["product_id"].(int64); ok && pid != 0 {
			product, err := service.chatRepository.GetChats(int(pid))
			if err == nil {
				chat["product"] = product
			}
			delete(chat, "product_id")
			chats[i] = chat
		}
	}

	// Bentuk respons sesuai format yang kamu mau
	response := map[string]interface{}{
		"name":  userName,
		"chats": chats,
	}

	return response, nil
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
	chat.Role = role
	chat.CreatedAt = time.Now()

	err = service.chatRepository.CreateChats(chat)
	if err != nil {
		return err
	}

	return nil
}

func (service *ChatServiceImpl) DeleteChatsUser(chatId int) error {
	return service.chatRepository.DeleteChats(chatId)
}
