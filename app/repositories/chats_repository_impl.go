package repositories

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
	"restfull-api-pjbl-2025/ws"
)

type ChatRepositoryImpl struct {
	db *gorm.DB
}

func NewChatRepositoryImpl(db *gorm.DB) *ChatRepositoryImpl {
	return &ChatRepositoryImpl{
		db: db,
	}
}

func (repo *ChatRepositoryImpl) GetAllChats(UserId int, AdminId int) ([]*dto.ResponseChat, error) {
	var Chats []*dto.ResponseChat
	err := repo.db.Table("chats").
		Joins("LEFT JOIN products on products.id = chats.product_id").
		Joins("LEFT JOIN product_images on product_images.product_id = products.id").
		Select("chats.id as id, chats.message as message, chats.user_id as user_id,products.name as name, products.price as price,chats.role_message as role_message, MIN(product_images.image_path) as image_path").
		Where("chats.user_id = ? AND chats.admin_id = ?", UserId, AdminId).
		Group("chats.id").
		Find(&Chats).Error

	if err != nil {
		return nil, err
	}

	return Chats, nil
}

func (repo *ChatRepositoryImpl) GetAdminId() (int, error) {
	var Admin model.Chat
	err := repo.db.Table("users").Where("role", "admin").First(&Admin).Error
	if err != nil {
		return 0, err
	}

	return Admin.Id, nil
}

func (repo *ChatRepositoryImpl) CreateChats(chat *model.Chat) error {
	err := repo.db.Table("chats").Create(chat).Error
	if err != nil {
		return err
	}

	message := map[string]interface{}{
		"id":           chat.Id,
		"user_id":      chat.UserId,
		"admin_id":     chat.AdminId,
		"role_message": chat.RoleMessage,
		"message":      chat.Message,
		"product_id":   chat.ProductId,
	}
	ws.WebSocketHub.BroadcastMessage(message, chat.UserId, chat.AdminId)

	return nil
}

func (repo *ChatRepositoryImpl) DeleteChats(chatId int) error {
	err := repo.db.Table("chats").Where("chat_id = ?", chatId).Delete(&model.Chat{}).Error
	if err != nil {
		return err
	}
	return nil
}
