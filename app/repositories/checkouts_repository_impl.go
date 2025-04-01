package repositories

import (
	"gorm.io/gorm"
	"restfull-api-pjbl-2025/model"
)

type CheckoutRepositoryImpl struct {
	db *gorm.DB
}

func NewCheckoutRepositoryImpl(db *gorm.DB) *CheckoutRepositoryImpl {
	return &CheckoutRepositoryImpl{
		db: db,
	}
}

func (repo *CheckoutRepositoryImpl) CreatePayment(order *model.Checkout) error {
	err := repo.db.Table("checkout").Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckoutRepositoryImpl) UpdateStatusPayment(orderId int, status string) error {
	err := repo.db.Table("checkout").Where("id = ?", orderId).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CheckoutRepositoryImpl) GetPaymentById(orderID int) (*model.Checkout, error) {
	var checkout *model.Checkout
	err := repo.db.Table("checkout").Where("id = ?", orderID).First(&checkout).Error
	if err != nil {
		return nil, err
	}
	return checkout, nil
}
