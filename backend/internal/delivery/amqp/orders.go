package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mag1c0/L0/backend/internal/domain"
	"github.com/mag1c0/L0/backend/internal/service"
	"github.com/nats-io/stan.go"
)

type OrdersConsumer struct {
	service service.Orders
	stan    stan.Conn
}

func NewOrdersConsumer(service service.Orders, stan stan.Conn) *OrdersConsumer {
	return &OrdersConsumer{service: service, stan: stan}
}

func (c *OrdersConsumer) Subscribe(ctx context.Context, subject string) error {
	_, err := c.stan.Subscribe(subject, func(msg *stan.Msg) {
		var order domain.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			fmt.Println("Failed to unmarshal json")
			return
		}
		if err := c.service.CreateOrder(ctx, &order); err != nil {
			fmt.Println("Failed to create order")
			return
		}
	})

	return err
}
