package service

import (
	"context"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/internal/repository"
)

type Orders interface {
	GetByID(ctx context.Context, id string) (domain.Order, error)
	GetAll(ctx context.Context) ([]domain.Order, error)
}

type Services struct {
	Orders Orders
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Orders: NewOrdersService(deps.Repos.Orders),
	}
}
