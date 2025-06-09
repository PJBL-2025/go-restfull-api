package services

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"restfull-api-pjbl-2025/app/repositories"
	"restfull-api-pjbl-2025/helper"
	"restfull-api-pjbl-2025/model"
	"restfull-api-pjbl-2025/model/dto"
	"strconv"
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

func (service *CheckoutServiceImpl) CreateOrderUser(userId int, checkout *map[string]interface{}) (int, string, string, error) {
	fmt.Println("CreateOrderUser")
	orderId := helper.GenerateOrderID()
	totalPrice := int((*checkout)["total_price"].(float64))
	addressId := int((*checkout)["address_id"].(float64))
	productCheckout := (*checkout)["product_checkout"].([]interface{})

	checkoutId, err := service.checkoutRepository.CreateCheckout(userId, totalPrice, addressId, orderId)
	if err != nil {
		return 0, "", "", err
	}

	fmt.Println(checkoutId)
	err = service.CreateOrderCustom(productCheckout, checkoutId)
	if err != nil {
		return 0, "", "", err
	}

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(orderId),
			GrossAmt: int64(totalPrice),
		},
	}
	snapResponse, err := service.snapClient.CreateTransaction(req)
	fmt.Println(err)
	err = service.checkoutRepository.InsertSnapToken(checkoutId, snapResponse.Token)
	if err != nil {
		return 0, "", "", err
	}
	return checkoutId, snapResponse.RedirectURL, snapResponse.Token, nil
}

func (service *CheckoutServiceImpl) CreateOrderCustom(productCheckout []interface{}, checkoutId int) error {
	for i := range productCheckout {
		product := productCheckout[i].(map[string]interface{})

		if product["type"] == "reguler" {
			productId := int(product["product_id"].(float64))
			orderProduct := model.ProductCheckout{
				Quantity:   int(product["quantity"].(float64)),
				Size:       product["size"].(string),
				Color:      product["color"].(string),
				Type:       product["type"].(string),
				Price:      int(product["price"].(float64)),
				ProductId:  &productId,
				CheckoutId: checkoutId,
			}

			productCustomId, err := service.checkoutRepository.CreateOrder(&orderProduct)
			if err != nil {
				return err
			}
			fmt.Println(productCustomId)
		}

		if product["type"] == "custom" {
			productId := int(product["product_id"].(float64))
			orderProduct := model.ProductCheckout{
				Quantity:   int(product["quantity"].(float64)),
				Size:       product["size"].(string),
				Color:      product["color"].(string),
				Type:       product["type"].(string),
				Price:      int(product["price"].(float64)),
				ProductId:  &productId,
				CheckoutId: checkoutId,
			}

			productCustomId, err := service.checkoutRepository.CreateOrder(&orderProduct)
			err = service.checkoutRepository.CreateProductCustom(product["product_custom"].(map[string]interface{}), productCustomId)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (service *CheckoutServiceImpl) UpdateStatusCheckout(checkout *dto.RequestUpdateCheckout) error {
	fmt.Println("cek")
	if checkout.Status == "processing" {
		deliveryId, err := service.checkoutRepository.CreateDelivery(checkout.CheckoutId)
		if err != nil {
			return err
		}

		fmt.Println("test")
		err = service.checkoutRepository.CreateStatusDelivery("Pesanan Sedang Di Kemas", deliveryId)
		if err != nil {
			return err
		}
	}

	err := service.checkoutRepository.UpdateStatusCheckout(checkout)
	if err != nil {
		return err
	}

	return nil
}

func (service *CheckoutServiceImpl) SetDelivery(delivery *dto.SetDelivery, deliveryId int) error {
	return service.checkoutRepository.SetDelivery(delivery, deliveryId)
}

func (service *CheckoutServiceImpl) SetStatusDelivery(status string, createdAt time.Time, deliveryId int) error {
	createdAt = time.Now()
	err := service.checkoutRepository.SetStatusDelivery(status, createdAt, deliveryId)
	if err != nil {
		return err
	}
	return nil
}
func (service *CheckoutServiceImpl) GetCheckout(param string, userId int) ([]map[string]interface{}, error) {
	var flatData []map[string]interface{}
	var err error

	if param != "all" {
		flatData, err = service.checkoutRepository.GetCheckout(param, userId)
		if err != nil {
			return nil, err
		}
	} else {
		flatData, err = service.checkoutRepository.GetCheckoutAll(userId)
		if err != nil {
			return nil, err
		}
	}

	grouped := map[int]map[string]interface{}{}

	for _, row := range flatData {
		id := int(row["id"].(int32))

		if _, exists := grouped[id]; !exists {
			grouped[id] = map[string]interface{}{
				"id":          id,
				"order_id":    row["order_id"],
				"total_price": row["total_price"],
				"snap_token":  row["snap_token"],
				"product":     []map[string]interface{}{},
			}
		}
		product := map[string]interface{}{
			"name":       row["name"],
			"price":      row["price"],
			"quantity":   row["quantity"],
			"type":       row["type"],
			"image_path": row["image_path"],
		}
		grouped[id]["product"] = append(grouped[id]["product"].([]map[string]interface{}), product)
	}

	var result []map[string]interface{}
	for _, item := range grouped {
		result = append(result, item)
	}

	return result, nil

}

func (service *CheckoutServiceImpl) GetDetailCheckoutProduct(productCheckoutId int) (map[string]interface{}, error) {
	data, err := service.checkoutRepository.GetDetailProductCheckoutAdmin(productCheckoutId)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	result := map[string]interface{}{
		"id":          data[0]["id"],
		"order_id":    data[0]["order_id"],
		"status":      data[0]["status"],
		"total_price": data[0]["total_price"],
		"snap_token":  data[0]["snap_token"],
	}

	// Data pengguna
	result["user"] = map[string]interface{}{
		"name":     data[0]["name"],
		"username": data[0]["username"],
	}

	// Alamat pengiriman
	result["address"] = map[string]interface{}{
		"address":          data[0]["address"],
		"zip_code":         data[0]["zip_code"],
		"destination_code": data[0]["destination_code"],
		"receiver_area":    data[0]["receiver_area"],
	}

	// Status pengiriman (menghindari duplikat)
	deliveryStatusSet := map[string]struct{}{}
	var deliveryStatuses []string
	for _, item := range data {
		if statusStr, ok := item["delivery_status"].(string); ok {
			if _, exists := deliveryStatusSet[statusStr]; !exists {
				deliveryStatusSet[statusStr] = struct{}{}
				deliveryStatuses = append(deliveryStatuses, statusStr)
			}
		}
	}

	// Data pengiriman
	result["deliveries"] = map[string]interface{}{
		"send_start_time": data[0]["send_start_time"],
		"send_end_time":   data[0]["send_end_time"],
		"delivery_status": deliveryStatuses,
		"id":              data[0]["delivery_id"],
	}

	// Menyusun data product_checkout tanpa duplikat
	productCheckouts := []map[string]interface{}{}
	productCheckoutMap := make(map[interface{}]bool) // untuk mengecek ID yang sudah ditambahkan

	for _, item := range data {
		productCheckoutID := item["product_checkout_id"] // hasil dari alias "product_checkout.id as id"

		if _, exists := productCheckoutMap[productCheckoutID]; exists {
			continue // skip jika sudah pernah ditambahkan
		}

		product := map[string]interface{}{
			"quantity": item["quantity"],
			"size":     item["size"],
			"color":    item["color"],
			"type":     item["type"],
			"price":    item["price"],
			"image":    item["image_path"],
			"name":     item["product_name"],
			"id":       item["product_checkout_id"], // id dari product_checkout
		}

		if item["type"] == "custom" {
			product["product_custom"] = map[string]interface{}{
				"front_image_path": item["front_image_path"],
				"back_image_path":  item["back_image_path"],
				"front_width":      item["front_width"],
				"back_width":       item["back_width"],
			}
		}

		productCheckouts = append(productCheckouts, product)
		productCheckoutMap[productCheckoutID] = true
	}

	result["product_checkout"] = productCheckouts

	return result, nil
}

func (service *CheckoutServiceImpl) GetCheckoutsAdmin() ([]map[string]interface{}, error) {
	result, err := service.checkoutRepository.GetCheckoutsAdmin()

	if err != nil {
		return nil, err
	}

	return result, nil
}
