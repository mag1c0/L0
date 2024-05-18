package amqp

import (
	"context"
	"github.com/mag1c0/L0/backend/internal/service"
	"github.com/nats-io/stan.go"
)

type Orders interface {
	Subscribe(ctx context.Context, subject string) error
}

type Consumer struct {
	Orders Orders
}

type Deps struct {
	Services *service.Services
	Stan     stan.Conn
}

func NewConsumer(deps Deps) *Consumer {
	return &Consumer{
		Orders: NewOrdersConsumer(deps.Services.Orders, deps.Stan),
	}
}
