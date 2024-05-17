package repository

import (
	"context"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/pkg/db"
)

const (
	tableName = "orders"
)

type BannersRepo struct {
	db db.Client
}

func NewOrdersRepo(db db.Client) *BannersRepo {
	return &BannersRepo{
		db: db,
	}
}

func (r *BannersRepo) GetByID(ctx context.Context, id string) (domain.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *BannersRepo) GetAll(ctx context.Context) ([]domain.Order, error) {
	//TODO implement me
	panic("implement me")
}
