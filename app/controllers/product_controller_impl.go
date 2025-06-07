package controllers

import (
	"github.com/gofiber/fiber/v2"
	"restfull-api-pjbl-2025/app/services"
	"restfull-api-pjbl-2025/helper"
	"restfull-api-pjbl-2025/model/dto"
)

type ProductControllerImpl struct {
	productService services.ProductService
}

func NewProductControllerImpl(productService services.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		productService: productService,
	}
}

func (controller *ProductControllerImpl) AddProductCheckout(ctx *fiber.Ctx) error {
	var product *dto.RequestProduct

	err := ctx.BodyParser(&product)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request fails", err)
	}

	err = controller.productService.AddProduct(product)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail Add checkout product", err)
	}

	return helper.SuccessResponse(ctx, product, "Success Add")
}

func (controller *ProductControllerImpl) UpdateProduct(ctx *fiber.Ctx) error {
	var product *dto.RequestProduct
	err := ctx.BodyParser(&product)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Parser request fails", err)
	}

	// Ambil ID dari parameter URL
	productId, err := ctx.ParamsInt("id")
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Invalid product ID", err)
	}
	product.Id = productId

	err = controller.productService.UpdateProduct(product)
	if err != nil {
		return helper.ErrorResponse(ctx, 400, "Fail update product", err)
	}

	return helper.SuccessResponse(ctx, product, "Success Update")
}

func (controller *ProductControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	data, err := controller.productService.GetAllCategories()
	if err != nil {
		return helper.ErrorResponse(ctx, 500, "Fail get all categories", err)
	}

	return helper.SuccessResponse(ctx, data, "Success GetAllCategories")
}
