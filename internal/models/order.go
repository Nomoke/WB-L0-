package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderUID        uuid.UUID   `json:"order_uid"`
	TrackNumber     string      `json:"track_number"`
	Entry           string      `json:"entry"`
	DeliveryID      int         `json:"delivery_id"`
	PaymentID       int         `json:"payment_id"`
	Locale          string      `json:"locale"`
	InternalSig     string      `json:"internal_signature"`
	CustomerID      string      `json:"customer_id"`
	DeliveryService string      `json:"delivery_service"`
	ShardKey        string      `json:"shard_key"`
	SMID            int         `json:"sm_id"`
	DateCreated     time.Time   `json:"date_created"`
	OofShard        string      `json:"oof_shard"`
	Delivery        Delivery    `json:"delivery"`
	Payment         Payment     `json:"payment"`
	Items           []OrderItem `gorm:"-" json:"items"`
}
