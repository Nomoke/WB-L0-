package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Nomoke/wb-test-app/internal/models"
	"github.com/Nomoke/wb-test-app/internal/storage"
	"github.com/Nomoke/wb-test-app/internal/storage/cache"
	"github.com/google/uuid"
)

type Order struct {
	OrderDataBase storage.OrderDBRepository
	OrderCache    cache.OrderCacheRepository
}

func New(db storage.OrderDBRepository, cache cache.OrderCacheRepository) *Order {
	return &Order{
		OrderDataBase: db,
		OrderCache:    cache,
	}
}

func (srv *Order) SaveOrder(ctx context.Context, order models.Order) error {
	const op = "service.Order.SaveOrder"

	// Сохраняем в БД
	ord, err := srv.OrderDataBase.SaveOrder(ctx, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Сохраняем в кеш
	srv.OrderCache.Set(*ord)
	return nil
}

func (srv *Order) GetOrderById(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	op := "service.Order.GetOrderById"
	
	var order *models.Order
	var err error

	order, err = srv.getOrderFromCache(id)
	if err == nil {
		// Возвращаем из кеша
		return order, nil
	}

	if err.Error() != "not found" {
		return order, fmt.Errorf("%s: %w", op, err)
	}

	order, err = srv.OrderDataBase.GetOrderById(ctx, id)
	if err != nil {
		// Возвращаем из БД
		return order, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (srv *Order) getOrderFromCache(id uuid.UUID) (*models.Order, error) {
	op := "service.Order.getOrderFromCache"

	order := models.Order{}
	orderAny, err := srv.OrderCache.Get(id)

	if err != nil {
		return &order, fmt.Errorf("%s: %w", op, err)
	}

	if orderAny == nil {
		return nil, errors.New("not found")
	}

	order, ok := orderAny.(models.Order)

	if !ok {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &order, nil
}

func (srv *Order) RecoverOrderCache(ctx context.Context) error {
	op := "service.Order.RecoverOrderCache"

	orders, err := srv.OrderDataBase.GetAll(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Сохраняем в кеш все заказы
	srv.OrderCache.SetAll(orders)
	return nil
}

func (srv *Order) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	op := "service.Order.GetAllOrders"
	orders, err := srv.OrderDataBase.GetAll(ctx)
	if err != nil {
		return orders, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}
