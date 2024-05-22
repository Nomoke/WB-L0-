package storage

import (
	"context"
	"fmt"

	"github.com/Nomoke/wb-test-app/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDBRepository interface {
	SaveOrder(ctx context.Context, order models.Order) (*models.Order, error)
	GetOrderById(ctx context.Context, id uuid.UUID) (*models.Order, error)
	SaveDelivery(ctx context.Context, delivery models.Delivery) (int, error)
	GetDeliveryById(ctx context.Context, deliveryId int) (*models.Delivery, error)
	SavePayment(ctx context.Context, payment models.Payment) (int, error)
	GetPaymentById(ctx context.Context, paymentId int) (*models.Payment, error)
	SaveItems(ctx context.Context, items []models.OrderItem) error
	GetItemByTrackNumber(ctx context.Context, trackNumber string) ([]models.OrderItem, error)
	GetAll(ctx context.Context) ([]models.Order, error)
}

type DataBase struct {
	*gorm.DB
}

func New(conn *gorm.DB) *DataBase {
	return &DataBase{conn}
}

func (db *DataBase) Close() {
	d, _ := db.DB.DB()
	d.Close()
}

func (db *DataBase) SaveOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	fmt.Println("saving [order]...")
	op := "storage.postgress.SaveOrder"

	deliveryID, err := db.SaveDelivery(ctx, order.Delivery)
	if err != nil {
		return &order, fmt.Errorf("%s: %w", op, err)
	}

	paymentID, err := db.SavePayment(ctx, order.Payment)
	if err != nil {
		return &order, fmt.Errorf("%s: %w", op, err)
	}

	sql := `INSERT INTO orders (order_uid, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard, track_number, entry)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	result := db.WithContext(ctx).Exec(sql, order.OrderUID, deliveryID, paymentID, order.Locale, order.InternalSig, order.CustomerID, order.DeliveryService,
		order.ShardKey, order.SMID, order.DateCreated, order.OofShard, order.TrackNumber, order.Entry)

	if result.Error != nil {
		return &order, fmt.Errorf("%s: %w", op, result.Error)
	}

	err = db.SaveItems(ctx, order.Items)
	if err != nil {
		return &order, fmt.Errorf("%s: %w", op, err)
	}

	order.Payment.ID = paymentID
	order.Delivery.ID = deliveryID

	fmt.Println("[order] saved successfully")
	return &order, nil
}

func (db *DataBase) GetOrderById(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	fmt.Println("getting [order] from postgres")
	op := "storage.postgress.GetOrderById"

	order := &models.Order{}

	sql := `SELECT * FROM orders WHERE order_uid = $1`
	result := db.WithContext(ctx).Raw(sql, id).Scan(&order)

	if result.Error != nil {
		return order, result.Error
	}

	delivery, err := db.GetDeliveryById(ctx, order.DeliveryID)
	if err != nil {
		return order, fmt.Errorf("%s: %w", op, err)
	}

	order.Delivery = *delivery

	payment, err := db.GetPaymentById(ctx, order.PaymentID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	order.Payment = *payment

	items, err := db.GetItemByTrackNumber(ctx, order.TrackNumber)
	if err != nil {
		return order, fmt.Errorf("%s: %w", op, err)
	}

	order.Items = items

	return order, nil
}

func (db *DataBase) SaveDelivery(ctx context.Context, delivery models.Delivery) (int, error) {
	fmt.Println("saving [delivery]...")
	op := "storage.postgress.SaveDelivery"

	sql := `INSERT INTO delivery (name, phone, zip, city, address, region, email) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var deliveryId int
	result := db.WithContext(ctx).Raw(sql, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email).Scan(&deliveryId)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	fmt.Println("[delevery] saved successfully")

	return deliveryId, nil

}

func (db *DataBase) GetDeliveryById(ctx context.Context, deliveryId int) (*models.Delivery, error) {
	fmt.Println("getting [delivery] from postgres")
	op := "storage.postgress.GetDeliveryById"

	delivery := models.Delivery{}

	sql := `SELECT * FROM delivery WHERE id = $1`
	result := db.WithContext(ctx).Raw(sql, deliveryId).Scan(&delivery)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &delivery, nil
}

func (db *DataBase) SavePayment(ctx context.Context, payment models.Payment) (int, error) {
	fmt.Println("saving [payment]...")
	op := "storage.postgress.SavePayment"

	sql := `INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	var paymentId int
	result := db.WithContext(ctx).Raw(sql, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee).Scan(&paymentId)

	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	fmt.Println("[payment] saved successfully")

	return paymentId, nil

}

func (db *DataBase) GetPaymentById(ctx context.Context, paymentId int) (*models.Payment, error) {
	fmt.Println("getting [payment] from postgres")
	op := "storage.postgress.GetPaymentById"

	payment := models.Payment{}

	sql := `SELECT * FROM payments WHERE id = $1`
	result := db.WithContext(ctx).Raw(sql, paymentId).Scan(&payment)

	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", op, result.Error)
	}

	return &payment, nil
}

func (db *DataBase) SaveItems(ctx context.Context, items []models.OrderItem) error {
	fmt.Println("saving [order items]...")
	op := "storage.postgress.SaveItems"

	sql := `INSERT INTO order_items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	for _, i := range items {
		result := db.WithContext(ctx).Exec(sql, i.ChrtId, i.TrackNumber, i.Price, i.RID, i.Name, i.Sale, i.Size, i.TotalPrice, i.NMID, i.Brand, i.Status)
		if result.Error != nil {
			return fmt.Errorf("%s: %w", op, result.Error)
		}
	}

	fmt.Println("[order items] saved successfully")
	return nil
}

func (db *DataBase) GetItemByTrackNumber(ctx context.Context, trackNumber string) ([]models.OrderItem, error) {
	fmt.Println("getting [order items] from postgres")
	op := "storage.postgress.GetItemByTrackNumber"

	items := []models.OrderItem{}

	sql := `SELECT * FROM order_items WHERE track_number = $1`
	result, err := db.WithContext(ctx).Raw(sql, trackNumber).Rows()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for result.Next() {
		item := models.OrderItem{}
		err := result.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NMID, &item.Brand, &item.Status)
		if err != nil {
			return items, fmt.Errorf("%s: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}

// используется для восстановления в случае падения сервиса
func (db *DataBase) GetAll(ctx context.Context) ([]models.Order, error) {
	fmt.Println("getting all [order] from postgres")
	op := "storage.postgress.GetAll"

	orders, err := db.getAllOrders(ctx)

	if err != nil {
		return orders, fmt.Errorf("%s: %w", op, err)
	}

	for i := 0; i < len(orders); i++ {
		delivery, err := db.GetDeliveryById(ctx, orders[i].Delivery.ID)
		if err != nil {
			return orders, fmt.Errorf("%s: %w", op, err)
		}
		orders[i].Delivery = *delivery

		items, err := db.GetItemByTrackNumber(ctx, orders[i].TrackNumber)
		if err != nil {
			return orders, fmt.Errorf("%s: %w", op, err)
		}
		orders[i].Items = items

		payment, err := db.GetPaymentById(ctx, orders[i].PaymentID)
		if err != nil {
			return orders, fmt.Errorf("%s: %w", op, err)
		}
		orders[i].Payment = *payment
	}

	return orders, nil
}

// Приватная функция для получения всех заказов
func (db *DataBase) getAllOrders(ctx context.Context) ([]models.Order, error) {
	orders := []models.Order{}
	op := "storage.postgress.getAllOrders"

	sql := `SELECT * FROM orders`

	rows, err := db.WithContext(ctx).Raw(sql).Rows()
	if err != nil {
		return orders, fmt.Errorf("%s: %w", op, err)
	}

	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Delivery.ID, &order.Payment.ID,
			&order.Locale, &order.CustomerID, &order.DeliveryService, &order.ShardKey,
			&order.SMID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return orders, fmt.Errorf("%s: %w", op, err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
