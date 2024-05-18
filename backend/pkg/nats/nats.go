package nats

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NsClient struct {
	Sc stan.Conn
}

func New(url, clusterID, clientID string) (*NsClient, error) {
	Client, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL(url),
	)
	if err != nil {
		fmt.Printf("failed to connect to NATS streaming server: %s", err.Error())
		return nil, err
	}

	return &NsClient{
		Sc: Client,
	}, nil
}
