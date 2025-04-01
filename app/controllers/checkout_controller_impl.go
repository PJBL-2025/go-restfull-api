package controllers

import (
	"github.com/gofiber/fiber/v2"
	"restfull-api-pjbl-2025/app/services"
	"restfull-api-pjbl-2025/helper"
	"restfull-api-pjbl-2025/model"
)

type CheckoutsControllerImpl struct {
	checkoutService services.CheckoutsService
}

func NewCheckoutsControllerImpl(checkoutService services.CheckoutsService) *CheckoutsControllerImpl {
	return &CheckoutsControllerImpl{checkoutService: checkoutService}
}

func (controller *CheckoutsControllerImpl) CreatePaymentUser(ctx *fiber.Ctx) error {
	var payment *model.Checkout
	userId := ctx.Locals("userId").(int)

	err := ctx.BodyParser(&payment)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Create Payment", err)
	}

	err = controller.checkoutService.CreateOrderUser(userId, payment)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Create Payment", err)
	}

	paymentURL, snapToken, err := controller.checkoutService.CreatePaymentUser(payment.Id, payment.TotalPrice)
	if err != nil {
		return helper.ErrorResponse(ctx, 500, "Midtrans error", err)
	}

	response := map[string]interface{}{
		"payment_url": paymentURL,
		"snap_token":  snapToken,
		"order":       payment,
	}

	return helper.SuccessResponse(ctx, response, "Success Checkout")
}

func (controller *CheckoutsControllerImpl) UpdateStatusPaymentUser(ctx *fiber.Ctx) error {
	orderId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Update Payment", err)
	}

	type Request struct {
		Status string `json:"status"`
	}
	var request *Request
	err = ctx.BodyParser(&request)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Parse Data", err)
	}

	err = controller.checkoutService.UpdateStatusPaymentUser(orderId, request.Status)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Update Payment", err)
	}

	return helper.SuccessResponse(ctx, request, "Success Update Payment")
}

func (controller *CheckoutsControllerImpl) GetPaymentUserById(ctx *fiber.Ctx) error {
	orderId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get Payment", err)
	}

	data, err := controller.checkoutService.GetPaymentUserById(orderId)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Get Payment", err)
	}

	return helper.SuccessResponse(ctx, data, "Success Get Payment")
}
