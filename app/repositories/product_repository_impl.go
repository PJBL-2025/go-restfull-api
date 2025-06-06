package repositories

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (repo *ProductRepositoryImpl) AddProduct(product *model.Product) (int, error) {
	err := repo.db.Table("products").Create(product).Error
	if err != nil {
		return 0, err
	}
	return product.Id, nil
}

func (repo *ProductRepositoryImpl) AddProductSize(productSize map[string]interface{}) error {
	err := repo.db.Table("product_size").Create(productSize).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) AddProductImage(productImage map[string]interface{}) error {

	err := repo.db.Table("product_images").Create(productImage).Error
	if err != nil {
		return err
	}
	return nil

}

func (repo *ProductRepositoryImpl) AddProductCategory(productCategory map[string]interface{}) error {
	err := repo.db.Table("product_category").Create(productCategory).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) UpdateProduct(product *model.Product) error {
	return repo.db.Table("products").Where("id = ?", product.Id).Updates(product).Error
}

func (repo *ProductRepositoryImpl) UpdateProductImage(productImage map[string]interface{}) error {
	return repo.db.Table("product_images").
		Where("product_id = ?", productImage["product_id"]).
		Updates(productImage).Error
}

func (repo *ProductRepositoryImpl) DeleteProductSizeByProductID(productId int) error {
	return repo.db.Table("product_size").Where("product_id = ?", productId).Delete(nil).Error
}

func (repo *ProductRepositoryImpl) DeleteProductCategoryByProductID(productId int) error {
	return repo.db.Table("product_category").Where("product_id = ?", productId).Delete(nil).Error
}

func (repo *ProductRepositoryImpl) FindRequestProductByID(id int) (*dto.RequestProduct, error) {
	var product model.Product
	err := repo.db.Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	var image struct {
		ImagePath string
	}
	repo.db.Table("product_images").Where("product_id = ?", id).First(&image)

	var sizes []struct {
		SizeID int
	}
	repo.db.Table("product_size").Where("product_id = ?", id).Find(&sizes)

	var categories []struct {
		CategoryID int
	}
	repo.db.Table("product_category").Where("product_id = ?", id).Find(&categories)

	// Konversi ukuran dan kategori ke slice int
	sizeIds := []int{}
	for _, s := range sizes {
		sizeIds = append(sizeIds, s.SizeID)
	}
	categoryIds := []int{}
	for _, c := range categories {
		categoryIds = append(categoryIds, c.CategoryID)
	}

	return &dto.RequestProduct{
		Id:          product.Id,
		Image:       image.ImagePath,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Weight:      product.Weight,
		Size:        sizeIds,
		CategoryId:  categoryIds,
		CreatedAt:   product.CreatedAt,
	}, nil
}
