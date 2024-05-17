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

func (s *OrdersService) GetByID(ctx context.Context, id string) (domain.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *OrdersService) GetAll(ctx context.Context) ([]domain.Order, error) {
	return s.repo.GetAll(ctx)
}
