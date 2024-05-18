package service

import (
	"context"
	"github.com/mag1c0/L0/backend/internal/cache"
	"github.com/mag1c0/L0/backend/internal/domain"
	"github.com/mag1c0/L0/backend/internal/repository"
)

type Orders interface {
	GetByID(ctx context.Context, uid string) (*domain.Order, error)
	GetAll(ctx context.Context) (*[]domain.Order, error)
	CreateOrder(ctx context.Context, order *domain.Order) error
}

type Services struct {
	Orders Orders
}

type Deps struct {
	Repos *repository.Repositories
	Cache *cache.Cache
}

func NewServices(deps Deps) *Services {
	return &Services{
		Orders: NewOrdersService(deps.Repos.Orders, deps.Cache),
	}
}
