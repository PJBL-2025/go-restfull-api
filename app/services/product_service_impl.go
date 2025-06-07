package services

import (
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
	"time"
)

type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
}

func NewProductServiceImpl(productRepository repositories.ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (service *ProductServiceImpl) AddProduct(product *dto.RequestProduct) error {
	product.CreatedAt = time.Now()

	addProduct := &model.Product{
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: product.Description,
		Weight:      product.Weight,
		CreatedAt:   time.Now(),
	}
	productId, err := service.productRepository.AddProduct(addProduct)
	if err != nil {
		return err
	}

	// Simpan gambar
	productImage := map[string]interface{}{
		"product_id": productId,
		"image_path": product.Image,
	}
	if err := service.productRepository.AddProductImage(productImage); err != nil {
		return err
	}

	// Simpan ukuran produk
	for _, size := range product.Size {
		productSize := map[string]interface{}{
			"product_id": productId,
			"size_id":    size,
		}
		if err := service.productRepository.AddProductSize(productSize); err != nil {
			return err
		}
	}

	for _, cat := range product.CategoryId {
		productCategory := map[string]interface{}{
			"product_id":  productId,
			"category_id": cat,
		}
		if err := service.productRepository.AddProductCategory(productCategory); err != nil {
			return err
		}
	}

	return nil
}

func (service *ProductServiceImpl) UpdateProduct(product *dto.RequestProduct) error {
	// Ambil data sebelumnya
	existing, err := service.productRepository.FindRequestProductByID(product.Id)
	if err != nil {
		return err
	}

	// Jika field kosong, gunakan nilai sebelumnya
	if product.Name == "" {
		product.Name = existing.Name
	}
	if product.Image == "" {
		product.Image = existing.Image
	}
	if product.Description == "" {
		product.Description = existing.Description
	}
	if product.Price == 0 {
		product.Price = existing.Price
	}
	if product.Quantity == 0 {
		product.Quantity = existing.Quantity
	}
	if product.Weight == 0 {
		product.Weight = existing.Weight
	}
	if len(product.Size) == 0 {
		product.Size = existing.Size
	}
	if len(product.CategoryId) == 0 {
		product.CategoryId = existing.CategoryId
	}

	// Update produk utama
	updateData := &model.Product{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: product.Description,
		Weight:      product.Weight,
	}

	if err := service.productRepository.UpdateProduct(updateData); err != nil {
		return err
	}

	// Update gambar
	productImage := map[string]interface{}{
		"product_id": product.Id,
		"image_path": product.Image,
	}
	if err := service.productRepository.UpdateProductImage(productImage); err != nil {
		return err
	}

	// Hapus & insert ulang ukuran produk
	if err := service.productRepository.DeleteProductSizeByProductID(product.Id); err != nil {
		return err
	}
	for _, size := range product.Size {
		sizeData := map[string]interface{}{
			"product_id": product.Id,
			"size_id":    size,
		}
		if err := service.productRepository.AddProductSize(sizeData); err != nil {
			return err
		}
	}

	// Hapus & insert ulang kategori produk
	if err := service.productRepository.DeleteProductCategoryByProductID(product.Id); err != nil {
		return err
	}
	for _, cat := range product.CategoryId {
		catData := map[string]interface{}{
			"product_id":  product.Id,
			"category_id": cat,
		}
		if err := service.productRepository.AddProductCategory(catData); err != nil {
			return err
		}
	}

	return nil
}

func (service *ProductServiceImpl) GetAllCategories() ([]dto.Category, error) {
	data, err := service.productRepository.GetAllCategories()
	if err != nil {
		return nil, err
	}

	return data, nil
}
