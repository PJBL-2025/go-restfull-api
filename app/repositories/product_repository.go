package repositories

import (
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type ProductRepository interface {
	AddProduct(product *model.Product) (int, error)
	AddProductSize(productSize map[string]interface{}) error
	AddProductImage(productImage map[string]interface{}) error
	AddProductCategory(productCategory map[string]interface{}) error

	UpdateProduct(product *model.Product) error
	DeleteProductSizeByProductID(productId int) error
	DeleteProductCategoryByProductID(productId int) error
	UpdateProductImage(productImage map[string]interface{}) error
	FindRequestProductByID(productId int) (*dto.RequestProduct, error)

	GetAllCategories() ([]dto.Category, error)
}
