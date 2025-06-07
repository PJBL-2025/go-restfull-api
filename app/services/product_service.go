package services

import "restfull-api-pjbl-2025/model/dto"

type ProductService interface {
	AddProduct(product *dto.RequestProduct) error
	UpdateProduct(product *dto.RequestProduct) error
	GetAllCategories() ([]dto.Category, error)
}
