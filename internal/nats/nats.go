package nats

import (
	"github.com/nats-io/nats.go"
)

type NatsConnection struct {
	*nats.Conn
}

func ConnectToServer(url string) (*NatsConnection, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsConnection{nc}, nil
}
