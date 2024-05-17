package repository

import (
	"context"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/pkg/db"
)

type Orders interface {
	GetByID(ctx context.Context, uid string) (*domain.Order, error)
	GetAll(ctx context.Context) (*[]domain.Order, error)
	CreateOrder(ctx context.Context, order domain.Order) error
}

type Repositories struct {
	Orders Orders
}

func NewRepositories(db db.Client) *Repositories {
	return &Repositories{
		Orders: NewOrdersRepo(db),
	}
}
