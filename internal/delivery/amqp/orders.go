package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/internal/service"
	"github.com/nats-io/stan.go"
)

type OrdersConsumer struct {
	service service.Orders
	stan    stan.Conn
}

func NewOrdersConsumer(service service.Orders, stan stan.Conn) *OrdersConsumer {
	return &OrdersConsumer{service: service, stan: stan}
}

func (c *OrdersConsumer) Subscribe(subject string) error {
	_, err := c.stan.Subscribe(subject, func(msg *stan.Msg) {
		fmt.Println("received message")
		var order domain.Order
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			fmt.Println("failed to unmarshal json")
			return
		}
		if err := c.service.CreateOrder(context.Background(), order); err != nil {
			fmt.Println("failed to create order")
			return
		}
		fmt.Println("created order")
	})
	return err
}
