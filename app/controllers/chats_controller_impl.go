package controllers

import (
	"github.com/gofiber/fiber/v2"
	"restfull-api-pjbl-2025/app/services"
	"restfull-api-pjbl-2025/helper"
	"restfull-api-pjbl-2025/model"
)

type ChatsControllerImpl struct {
	chatService services.ChatsService
}

func NewChatsControllerImpl(chatService services.ChatsService) *ChatsControllerImpl {
	return &ChatsControllerImpl{chatService: chatService}
}

func (controller *ChatsControllerImpl) GetAllChatsUser(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)
	role := ctx.Locals("role").(string)
	queryId := ctx.Query("id", "0")

	data, err := controller.chatService.GetAllChatsUser(userId, role, queryId)
	if err != nil {
		return helper.ErrorResponse(ctx, 404, "Fail Get Chats", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Chats")
}

func (controller *ChatsControllerImpl) CreateChatsUser(ctx *fiber.Ctx) error {
	var chat *model.Chat
	userId := ctx.Locals("userId").(int)
	role := ctx.Locals("role").(string)
	queryId := ctx.Query("id", "0")

	err := ctx.BodyParser(&chat)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Create Chats User", err)
	}

	err = controller.chatService.CreateChatsUser(chat, userId, role, queryId)
	if err != nil {
		return helper.ErrorResponse(ctx, 404, "Fail Create Chats", err)
	}

	return helper.SuccessResponse(ctx, chat, "Success Create Chats")
}

func (controller *ChatsControllerImpl) DeleteChatUser(ctx *fiber.Ctx) error {
	chatId, err := ctx.ParamsInt("chatId")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get ChatId", err)
	}

	err = controller.chatService.DeleteChatsUser(chatId)
	if err != nil {
		return helper.ErrorResponse(ctx, 404, "Fail Delete Chats", err)
	}

	return helper.SuccessResponse(ctx, 0, "Success Delete Chats")
}
