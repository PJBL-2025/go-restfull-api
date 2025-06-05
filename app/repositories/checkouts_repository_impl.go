package repositories

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
	"time"
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

func (repo *CheckoutRepositoryImpl) SetStatusDelivery(status string, createdAt time.Time, deliveryId int) error {
	data := model.StatusDelivery{
		Status:       status,
		CreatedAt:    createdAt,
		DeliveriesId: deliveryId,
	}

	err := repo.db.Table("delivery_status").Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckoutRepositoryImpl) GetCheckout(param string, userId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("checkouts").
		Joins("LEFT JOIN product_checkout on product_checkout.checkout_id = checkouts.id").
		Joins("LEFT JOIN products on products.id = product_checkout.product_id").
		Joins("LEFT JOIN product_images on product_images.product_id = products.id").
		Select("checkouts.snap_token as snap_token,checkouts.id as id,product_checkout.type as type, checkouts.order_id as order_id, checkouts.total_price as total_price, products.name as name, product_checkout.price as price, product_checkout.quantity as quantity, MIN(product_images.image_path) as image_path").
		Where("checkouts.user_id = ? AND checkouts.status = ?", userId, param).
		Group("checkouts.order_id, checkouts.id,product_checkout.type,product_checkout.id,products.id, products.name, product_checkout.price, product_checkout.quantity").
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
		Joins("LEFT JOIN addresses ON addresses.id = checkouts.address_id").
		Select("products.name as product_name,addresses.address as address,products.price	as price,checkouts.order_id as order_id,deliveries.send_end_time as date,users.name as user_name,checkouts.status as checkout_status,delivery_status.status as delivery_status,MIN(product_images.image_path) as image_path").
		Where("product_checkout.id = ?", productCheckoutId).
		Group("products.name,products.price,checkouts.order_id,deliveries.send_end_time,users.name,checkouts.status,delivery_status.status,addresses.address ").
		Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *CheckoutRepositoryImpl) GetDetailProductCheckoutAdmin(productCheckoutId int) ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("product_checkout").
		Joins("LEFT JOIN checkouts ON checkouts.id = product_checkout.checkout_id").
		Joins("LEFT JOIN users ON users.id = checkouts.user_id").
		Joins("LEFT JOIN products ON products.id = product_checkout.product_id").
		Joins("LEFT JOIN product_custom ON product_custom.product_id = product_checkout.id").
		Joins("LEFT JOIN product_images ON product_images.product_id = products.id").
		Joins("LEFT JOIN deliveries ON deliveries.checkout_id = checkouts.id").
		Joins("LEFT JOIN delivery_status ON delivery_status.deliveries_id = deliveries.id").
		Joins("LEFT JOIN addresses ON addresses.id = checkouts.address_id").
		Select("product_checkout.id as id,"+
			"checkouts.order_id as order_id,"+
			"checkouts.status as status,"+
			"checkouts.total_price as total_price,"+
			"checkouts.snap_token as snap_token,"+
			"product_checkout.quantity as quantity,"+
			"product_checkout.size as size,"+
			"product_checkout.type as type,"+
			"product_checkout.color as color,"+
			"product_checkout.price as price,"+
			"product_custom.front_image_path as front_image_path,"+
			"product_custom.back_image_path as back_image_path,"+
			"product_custom.back_width as back_width,"+
			"product_custom.front_width as front_width,"+
			"users.name as name,"+
			"users.username as username,"+
			"addresses.address as address,"+
			"addresses.zip_code as zip_code,"+
			"addresses.destination_code as destination_code,"+
			"addresses.receiver_area as receiver_area,"+
			"deliveries.send_start_time as send_start_time,"+
			"deliveries.send_end_time as send_end_time,"+
			"delivery_status.status as delivery_status,"+
			"MIN(product_images.image_path) as image_path").
		Where("checkouts.id = ?", productCheckoutId).
		Group(`product_checkout.id, 
	checkouts.order_id, checkouts.status, checkouts.total_price, checkouts.snap_token,
	product_checkout.quantity, product_checkout.size, product_checkout.type, 
	product_checkout.color, product_checkout.price,
	product_custom.front_image_path, product_custom.back_image_path, 
	product_custom.back_width, product_custom.front_width,
	users.name, users.username,
	addresses.address, addresses.zip_code, addresses.destination_code, addresses.receiver_area,
	deliveries.send_start_time, deliveries.send_end_time,
	delivery_status.status,
	products.name, products.price`).
		Find(&data).Error

	if err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *CheckoutRepositoryImpl) GetCheckoutsAdmin() ([]map[string]interface{}, error) {
	var data []map[string]interface{}

	err := repo.db.Table("checkouts").
		Joins("LEFT JOIN users ON users.id = checkouts.user_id").
		Select("checkouts.order_id as order_id," +
			"checkouts.id as id," +
			"checkouts.status as status," +
			"checkouts.total_price as total_price," +
			"checkouts.created_at as date," +
			"users.name as name," +
			"users.username as username").
		Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}
