package subscriber

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Nomoke/wb-test-app/internal/models"
	"github.com/Nomoke/wb-test-app/internal/nats"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go/jetstream"
)

type orderService interface {
	GetOrderById(ctx context.Context, id uuid.UUID) (*models.Order, error)
	SaveOrder(ctx context.Context, order models.Order) error
	RecoverOrderCache(ctx context.Context) error
	GetAllOrders(ctx context.Context) ([]models.Order, error)
}

type Subscriber struct {
	nc      *nats.NatsConnection
	service orderService
}

func New(nc *nats.NatsConnection, service orderService) error {
	subscriber := &Subscriber{nc: nc, service: service}
	js, err := jetstream.New(subscriber.nc.Conn)

	if err != nil {
		return err
	}

	stream, err := js.Stream(context.Background(), "orders")
	if err != nil {
		return err
	}

	consumer, err := stream.CreateOrUpdateConsumer(context.Background(), jetstream.ConsumerConfig{
		Name:    "order_prosessor",
		Durable: "order_prosessor",
	})

	if err != nil {
		return err
	}

	// gorutine?
	_, err = consumer.Consume(func(msg jetstream.Msg) {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()

		var order models.Order

		err := json.Unmarshal(msg.Data(), &order)
		if err != nil {
			fmt.Println("unable to unmarshal msg: ", err)
		}

		fmt.Println(order.OrderUID)

		err = service.SaveOrder(ctx, order)
		if err != nil {
			fmt.Println("unable to save order: ", err)
		}

		msg.Ack()
	})

	if err != nil {
		return err
	}

	return nil
}

// // prosess messages
// func prosessMessages(m *jetstream.Msg) {
// 	fmt.Println("prosessing messages...")

// 	fmt.Println(string(m.).)
// }
