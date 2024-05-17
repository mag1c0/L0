package service

import (
	"context"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/internal/repository"
)

type OrdersService struct {
	repo repository.Orders
}

func NewOrdersService(repo repository.Orders) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) GetByID(ctx context.Context, uid string) (*domain.Order, error) {
	return s.repo.GetByID(ctx, uid)
}

func (s *OrdersService) GetAll(ctx context.Context) (*[]domain.Order, error) {
	return s.repo.GetAll(ctx)
}

func (s *OrdersService) CreateOrder(ctx context.Context, order domain.Order) error {
	return s.repo.CreateOrder(ctx, order)
}
