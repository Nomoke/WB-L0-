package cache

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Nomoke/wb-test-app/internal/logger"
	"github.com/Nomoke/wb-test-app/internal/models"
	"github.com/google/uuid"
)

type OrderCacheRepository interface {
	Get(id uuid.UUID) (any, error)
	Set(order models.Order)
	SetAll(orders []models.Order)
}

type OrderCache struct {
	sync.Map
	log *logger.Logger
}

func NewOrder(log *logger.Logger) *OrderCache {
	return &OrderCache{}
}

func (cache *OrderCache) Get(id uuid.UUID) (any, error) {
	cache.log.Info("getting [order] from cache id = ", id)

	value, ok := cache.Load(id)
	if !ok {
		fmt.Println("not found [order] in cache")
		return nil, errors.New("not found")
	}

	cache.log.Info("returned [order] from cache")
	return value, nil
}

func (cache *OrderCache) Set(order models.Order) {
	cache.log.Info("cache save [order] key = ", order.OrderUID)
	cache.Store(order.OrderUID, order)
}

func (cache *OrderCache) SetAll(orders []models.Order) {
	len := len(orders)
	var wg sync.WaitGroup
	wg.Add(len)

	for _, ord := range orders {
		go func(o models.Order) {
			cache.Set(o)
		}(ord)
	}

	cache.log.Info("totally recovered %d orders\n", len)
}
