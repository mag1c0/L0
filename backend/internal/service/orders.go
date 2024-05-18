package service

import (
	"context"
	"fmt"
	"github.com/mag1c0/L0/backend/internal/cache"
	"github.com/mag1c0/L0/backend/internal/domain"
	"github.com/mag1c0/L0/backend/internal/repository"
	"time"
)

type OrdersService struct {
	repo  repository.Orders
	cache *cache.Cache
}

func NewOrdersService(repo repository.Orders, cache *cache.Cache) *OrdersService {
	s := OrdersService{repo: repo, cache: cache}
	s.restoreCache()

	return &s
}

func (s *OrdersService) GetByID(ctx context.Context, uid string) (*domain.Order, error) {
	item, exist := s.cache.Get(uid)
	if !exist {
		order, err := s.repo.GetByID(ctx, uid)
		if err != nil {
			return nil, err
		}

		s.cache.Set(order.OrderUID, order, cache.DefaultExpire)

		return order, nil
	}

	return item.(*domain.Order), nil
}

func (s *OrdersService) GetAll(ctx context.Context) (*[]domain.Order, error) {
	item, exist := s.cache.Get("getAllOrder")
	if !exist {
		list, err := s.repo.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		s.cache.Set("getAllOrder", list, time.Second*10)

		return list, nil
	}

	return item.(*[]domain.Order), nil
}

func (s *OrdersService) CreateOrder(ctx context.Context, order *domain.Order) error {
	err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}

	s.cache.Set(order.OrderUID, order, cache.DefaultExpire)

	return nil
}

func (s *OrdersService) restoreCache() {
	list, err := s.GetAll(context.Background())
	if err != nil {
		fmt.Println("OrderService.restoreCache: failed to get all orders")
		return
	}

	for _, v := range *list {
		s.cache.Set(v.OrderUID, &v, cache.DefaultExpire)
	}
}
