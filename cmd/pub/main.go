package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Nomoke/wb-test-app/cmd/pub/data"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx := context.Background()

	// Подключение к серверу NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Failed to connect to NATS server: ", err)
	}
	defer nc.Close()
	log.Println("Connected to NATS server")

	// Получение доступа к JetStream
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal("Failed to create JetStream connection: ", err)
	}

	// Создание стрима
	cfg := jetstream.StreamConfig{
		Name:        "orders",
		Description: "orders stream",
		Subjects:    []string{"ORDERS.*"},
	}
	_, err = js.CreateOrUpdateStream(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to create stream: ", err)
	}

	// Публикация сообщения
	for {
		order := data.RandomOrder()
		orderJSON, err := json.Marshal(order)
		if err != nil {
			log.Fatal("Failed to marshal order: ", err)
		}

		_, err = js.Publish(ctx, "ORDERS.NEW", orderJSON)
		if err != nil {
			log.Fatal("Failed to publish order: ", err)
		}

		log.Println("Order published successfully")

		// Пауза между публикациями
		time.Sleep(10 * time.Second)
	}
}
