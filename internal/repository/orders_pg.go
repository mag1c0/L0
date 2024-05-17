package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/mag1c0/L0/internal/domain"
	"github.com/mag1c0/L0/pkg/db"
)

const (
	tableName = "orders"

	orderUidColumn          = "order_uid"
	trackNumberColumn       = "track_number"
	entryColumn             = "entry"
	localeColumn            = "locale"
	internalSignatureColumn = "internal_signature"
	customerIdColumn        = "customer_id"
	deliveryServiceColumn   = "delivery_service"
	shardkeyColumn          = "shardkey"
	smIdColumn              = "sm_id"
	dateCreatedColumn       = "date_created"
	oofShardColumn          = "oof_shard"

	nameColumn    = "name"
	phoneColumn   = "phone"
	zipColumn     = "zip"
	cityColumn    = "city"
	addressColumn = "address"
	regionColumn  = "region"
	emailColumn   = "email"

	transactionColumn  = "transaction"
	requestIdColumn    = "request_id"
	currencyColumn     = "currency"
	providerColumn     = "provider"
	amountColumn       = "amount"
	paymentDtColumn    = "payment_dt"
	bankColumn         = "bank"
	deliveryCostColumn = "delivery_cost"
	goodsTotalColumn   = "goods_total"
	customFeeColumn    = "custom_fee"

	idProductColumn   = "id"
	chrtIdColumn      = "chrt_id"
	priceColumn       = "price"
	ridColumn         = "rid"
	nameProductColumn = "name"
	saleColumn        = "sale"
	sizeColumn        = "size"
	totalPriceColumn  = "total_price"
	nmIdColumn        = "nm_id"
	brandColumn       = "brand"
	statusColumn      = "status"
)

type OrdersRepo struct {
	db db.Client
}

func NewOrdersRepo(db db.Client) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (r *OrdersRepo) GetByID(ctx context.Context, uid string) (*domain.Order, error) {
	builder := sq.Select(orderUidColumn, trackNumberColumn, entryColumn, localeColumn, internalSignatureColumn, customerIdColumn, deliveryServiceColumn, shardkeyColumn, smIdColumn, dateCreatedColumn, oofShardColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{orderUidColumn: uid}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "orders_repository.GetAll",
		QueryRaw: query,
	}

	var order domain.Order

	err = r.db.DB().ScanAllContext(ctx, &order, q, args...)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrdersRepo) GetAll(ctx context.Context) (*[]domain.Order, error) {
	builder := sq.Select(orderUidColumn, trackNumberColumn, entryColumn, localeColumn, internalSignatureColumn, customerIdColumn, deliveryServiceColumn, shardkeyColumn, smIdColumn, dateCreatedColumn, oofShardColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "orders_repository.GetAll",
		QueryRaw: query,
	}

	list := make([]domain.Order, 0)

	err = r.db.DB().ScanAllContext(ctx, &list, q, args...)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, order domain.Order) error {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(orderUidColumn, trackNumberColumn, entryColumn, localeColumn, internalSignatureColumn, customerIdColumn, deliveryServiceColumn, shardkeyColumn, smIdColumn, dateCreatedColumn, oofShardColumn).
		Values(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "orders_repository.Create",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
