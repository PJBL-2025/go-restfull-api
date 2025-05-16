package repositories

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
)

type CheckoutRepositoryImpl struct {
	db *gorm.DB
}

func NewCheckoutRepositoryImpl(db *gorm.DB) *CheckoutRepositoryImpl {
	return &CheckoutRepositoryImpl{
		db: db,
	}
}

func (repo *CheckoutRepositoryImpl) CreateCheckout(userId int, totalPrice int, addressId int, orderId int) (int, error) {
	checkout := model.Checkout{
		TotalPrice: totalPrice,
		AddressId:  addressId,
		UserId:     userId,
		Status:     "pending",
		OrderId:    orderId,
	}

	err := repo.db.Table("checkouts").Create(&checkout).Error
	if err != nil {
		return 0, err
	}
	return checkout.Id, nil
}

func (repo *CheckoutRepositoryImpl) InsertSnapToken(checkoutId int, snap string) error {
	err := repo.db.Table("checkouts").Where("id", checkoutId).Update("snap_token", snap).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CheckoutRepositoryImpl) CreateOrder(order *model.ProductCheckout) (int, error) {

	err := repo.db.Table("product_checkout").Create(&order).Error
	if err != nil {
		return 0, err
	}
	return order.Id, nil
}

func (repo *CheckoutRepositoryImpl) CreateProductCustom(custom map[string]interface{}, productId int) error {
	custom["product_id"] = productId

	err := repo.db.Table("product_custom").Create(custom).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CheckoutRepositoryImpl) CreateDelivery(checkoutId int) (int, error) {
	delivery := model.Delivery{
		CheckoutId: checkoutId,
	}

	err := repo.db.Table("deliveries").Create(&delivery).Error
	if err != nil {
		return 0, err
	}
	return delivery.Id, nil
}

func (repo *CheckoutRepositoryImpl) CreateStatusDelivery(status string, deliveryId int) error {
	statusDelivery := model.StatusDelivery{
		Status:       status,
		DeliveriesId: deliveryId,
	}

	err := repo.db.Table("delivery_status").Create(&statusDelivery).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CheckoutRepositoryImpl) UpdateStatusCheckout(checkout *dto.RequestUpdateCheckout) error {
	err := repo.db.Table("checkouts").Where("checkouts.id = ?", checkout.CheckoutId).Update("status", checkout.Status).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CheckoutRepositoryImpl) SetDelivery(delivery *dto.SetDelivery, deliveryId int) error {
	err := repo.db.Table("deliveries").Where("id", deliveryId).Updates(&delivery).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CheckoutRepositoryImpl) SetStatusDelivery(status map[string]interface{}) error {
	err := repo.db.Table("delivery_status").Create(status).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckoutRepositoryImpl) GetCheckoutPending(param string, userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("checkouts").
		Joins("LEFT JOIN product_checkout on product_checkout.checkout_id = checkouts.id").
		Joins("LEFT JOIN products on products.id = product_checkout.product_id").
		Joins("LEFT JOIN product_images on product_images.product_id = products.id").
		Select("checkouts.id as id,product_checkout.type as type, checkouts.order_id as order_id, checkouts.total_price as total_price, products.name as name, product_checkout.price as price, product_checkout.quantity as quantity, MIN(product_images.image_path) as image_path").
		Where("checkouts.user_id = ? AND checkouts.status = ?", userId, param).
		Group("checkouts.order_id, checkouts.id,product_checkout.type,product_checkout.id,products.id, products.name, product_checkout.price, product_checkout.quantity").
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *CheckoutRepositoryImpl) GetCheckoutNotPending(param string, userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("checkouts").
		Joins("LEFT JOIN product_checkout ON product_checkout.checkout_id = checkouts.id").
		Joins("LEFT JOIN products ON products.id = product_checkout.product_id").
		Joins("LEFT JOIN product_images ON product_images.product_id = products.id").
		Select("checkouts.status as status,product_checkout.type as type,product_checkout.id as product_checkout_id,products.id as product_id, products.name as name, product_checkout.price as price, product_checkout.quantity as quantity, MIN(product_images.image_path) as image_path").
		Where("checkouts.user_id = ? AND checkouts.status = ?", userId, param).
		Group("checkouts.status, products.name,product_checkout.type,product_checkout.id,products.id, product_checkout.price, product_checkout.quantity").
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (repo *CheckoutRepositoryImpl) GetDetailProductCheckout(productCheckoutId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("product_checkout").
		Joins("LEFT JOIN checkouts ON checkouts.id = product_checkout.checkout_id").
		Joins("LEFT JOIN users ON users.id = checkouts.user_id").
		Joins("LEFT JOIN products ON products.id = product_checkout.product_id").
		Joins("LEFT JOIN product_images ON product_images.product_id = products.id").
		Joins("LEFT JOIN deliveries ON deliveries.checkout_id = checkouts.id").
		Joins("LEFT JOIN delivery_status ON delivery_status.deliveries_id = deliveries.id").
		Select("products.name as product_name,products.price	as price,checkouts.order_id as order_id,deliveries.send_end_time as date,users.name as user_name,checkouts.status as checkout_status,delivery_status.status as delivery_status,MIN(product_images.image_path) as image_path").
		Where("product_checkout.id = ?", productCheckoutId).
		Group("products.name,products.price,checkouts.order_id,deliveries.send_end_time,users.name,checkouts.status,delivery_status.status").
		Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}
