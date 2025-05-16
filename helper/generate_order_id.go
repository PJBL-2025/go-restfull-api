package helper

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOrderID() int {
	now := time.Now()

	datePart := now.Format("020106")

	rand.Seed(time.Now().UnixNano())
	randomPart := fmt.Sprintf("%03d", rand.Intn(1000))

	orderID, err := strconv.Atoi(datePart + randomPart)
	if err != nil {
		return 0
	}
	return orderID
}
