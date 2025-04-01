package services

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/model"
	"time"
)

type CheckoutServiceImpl struct {
	checkoutRepository repositories.CheckoutsRepository
	snapClient         *snap.Client
}

func NewCheckoutServiceImpl(checkoutRepository repositories.CheckoutsRepository, snapClient *snap.Client) *CheckoutServiceImpl {
	return &CheckoutServiceImpl{
		checkoutRepository: checkoutRepository,
		snapClient:         snapClient,
	}
}

func (service *CheckoutServiceImpl) CreateOrderUser(userId int, order *model.Checkout) error {
	timeNow := time.Now()

	order.UserId = userId
	order.CreatedAt = &timeNow
	order.Status = "pending"

	err := service.checkoutRepository.CreatePayment(order)
	if err != nil {
		return err
	}
	return nil
}

func (service *CheckoutServiceImpl) CreatePaymentUser(orderId int, totalPrice int) (string, string, error) {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("order-%d", orderId),
			GrossAmt: int64(totalPrice),
		},
	}

	snapResponse, err := service.snapClient.CreateTransaction(req)
	if err != nil {
		return "", "", err
	}

	return snapResponse.RedirectURL, snapResponse.Token, nil
}

func (service *CheckoutServiceImpl) UpdateStatusPaymentUser(orderId int, status string) error {
	err := service.checkoutRepository.UpdateStatusPayment(orderId, status)
	if err != nil {
		return err
	}
	return nil
}

func (service *CheckoutServiceImpl) GetPaymentUserById(orderID int) (*model.Checkout, error) {
	data, err := service.checkoutRepository.GetPaymentById(orderID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
