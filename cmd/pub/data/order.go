package data

import (
	"github.com/Nomoke/wb-test-app/internal/models"
)

func RandomOrder() models.Order {
	trackNumber := randomString(10)

	items := randomOrderItems(randomInt(1, 5), trackNumber)

	order := models.Order{
		OrderUID:        randomUIID(),
		TrackNumber:     trackNumber,
		Entry:           "WBIL",
		Delivery:        randomDelivery(),
		Payment:         randomPayment(),
		Items:           items,
		Locale:          "en",
		InternalSig:     randomString(19),
		CustomerID:      randomString(3),
		DeliveryService: "wb-delivery",
		ShardKey:        randomString(1),
		SMID:            randomInt(1, 100),
		DateCreated:     randomTime(),
		OofShard:        randomString(1),
	}

	return order
}

func randomDelivery() models.Delivery {
	delivery := models.Delivery{
		ID:      randomInt(1, 20),
		Name:    randomString(10),
		Phone:   randomPhone(),
		Zip:     randomString(6),
		City:    randomString(6),
		Address: randomString(10),
		Region:  randomString(11),
		Email:   randomEmail(),
	}

	return delivery
}

func randomPayment() models.Payment {
	payment := models.Payment{
		ID:           randomInt(1, 20),
		Transaction:  randomUIID(),
		RequestID:    randomString(10),
		Currency:     "RUB",
		Provider:     "wbpay",
		Amount:       randomInt(100, 10000),
		PaymentDT:    randomTime(),
		Bank:         "wb-bank",
		DeliveryCost: randomInt(700, 1500),
		GoodsTotal:   randomInt(200, 1000),
		CustomFee:    0,
	}

	return payment
}

func randomOrderItem(trackNumber string) models.OrderItem {
	item := models.OrderItem{
		ChrtId:      randomInt(100, 10000),
		TrackNumber: trackNumber,
		Price:       randomInt(1000, 5000),
		RID:         randomUIID(),
		Name:        randomString(10),
		Sale:        randomInt(0, 51),
		Size:        randomString(3),
		TotalPrice:  randomInt(250, 1600),
		NMID:        randomInt(100000, 1500000),
		Brand:       randomString(10),
		Status:      randomInt(100, 200),
	}

	return item
}

func randomOrderItems(n int, trackNumber string) []models.OrderItem {
	items := []models.OrderItem{}

	for i := 0; i < n; i++ {
		items = append(items, randomOrderItem(trackNumber))
	}

	return items
}
