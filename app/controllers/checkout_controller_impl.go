package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"restfull-api-pjbl-2025/app/services"
	"restfull-api-pjbl-2025/helper"
	"restfull-api-pjbl-2025/model/dto"
)

type CheckoutsControllerImpl struct {
	checkoutService services.CheckoutsService
}

func NewCheckoutsControllerImpl(checkoutService services.CheckoutsService) *CheckoutsControllerImpl {
	return &CheckoutsControllerImpl{checkoutService: checkoutService}
}

func (controller *CheckoutsControllerImpl) CreateOrderProduct(ctx *fiber.Ctx) error {
	var checkout map[string]interface{}
	userId := ctx.Locals("userId").(int)

	err := ctx.BodyParser(&checkout)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request", err)
	}

	paymentURL, snapToken, err := controller.checkoutService.CreateOrderUser(userId, &checkout)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Create order", err)
	}

	response := map[string]interface{}{
		"payment_url": paymentURL,
		"snap_token":  snapToken,
	}

	return helper.SuccessResponse(ctx, response, "Success Checkout")
}

func (controller *CheckoutsControllerImpl) UpdateStatusCheckout(ctx *fiber.Ctx) error {
	var checkout *dto.RequestUpdateCheckout
	err := ctx.BodyParser(&checkout)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request fails", err)
	}

	err = controller.checkoutService.UpdateStatusCheckout(checkout)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Update checkout", err)
	}

	return helper.SuccessResponse(ctx, "", "Success Update Status")
}

func (controller *CheckoutsControllerImpl) SetDelivery(ctx *fiber.Ctx) error {
	var delivery dto.SetDelivery

	deliveryId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parse delivery id fail", err)
	}

	err = ctx.BodyParser(&delivery)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request fails", err)
	}

	err = controller.checkoutService.SetDelivery(&delivery, deliveryId)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Update checkout", err)
	}
	return helper.SuccessResponse(ctx, "", "Success Set Delivery")
}

func (controller *CheckoutsControllerImpl) SetStatusDelivery(ctx *fiber.Ctx) error {
	var status map[string]interface{}

	err := ctx.BodyParser(&status)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request fails", err)
	}

	err = controller.checkoutService.SetStatusDelivery(status)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Update Status Delivery", err)
	}

	return helper.SuccessResponse(ctx, "", "Success Set Status Delivery")

}

func (controller *CheckoutsControllerImpl) GetCheckout(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)
	param := ctx.Params("status")
	fmt.Println(param)

	data, err := controller.checkoutService.GetCheckout(param, userId)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get checkout", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Checkout")
}

func (controller *CheckoutsControllerImpl) GetDetailProductCheckout(ctx *fiber.Ctx) error {
	productCheckoutId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parse product checkout id fail", err)
	}

	data, err := controller.checkoutService.GetDetailCheckoutProduct(productCheckoutId)

	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get checkout product", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Checkout Product")
}

func (controller *CheckoutsControllerImpl) GetDetailProductCheckoutAdmin(ctx *fiber.Ctx) error {
	productCheckoutId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parse product checkout id fail", err)
	}

	data, err := controller.checkoutService.GetDetailCheckoutProductAdmin(productCheckoutId)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get checkout product", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Checkout Product")
}

func (controller *CheckoutsControllerImpl) GetCheckoutsAdmin(ctx *fiber.Ctx) error {
	data, err := controller.checkoutService.GetCheckoutsAdmin()
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get checkout admin", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Checkout Admin")
}
