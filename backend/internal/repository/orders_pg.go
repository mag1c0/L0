package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/mag1c0/L0/backend/internal/domain"
	"github.com/mag1c0/L0/backend/pkg/db"
)

const (
	ordersTableName         = "orders"
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

	deliveriesTableName = "order_deliveries"
	nameColumn          = "name"
	phoneColumn         = "phone"
	zipColumn           = "zip"
	cityColumn          = "city"
	addressColumn       = "address"
	regionColumn        = "region"
	emailColumn         = "email"

	paymentsTableName  = "order_payments"
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

	itemsTableName    = "order_items"
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
		From(ordersTableName).
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

	err = r.db.DB().ScanOneContext(ctx, &order, q, args...)
	if err != nil {
		return nil, err
	}

	order.Delivery, err = r.GetDeliveryByID(ctx, order.OrderUID)
	if err != nil {
		return nil, err
	}

	order.Payment, err = r.GetPaymentByID(ctx, order.OrderUID)
	if err != nil {
		return nil, err
	}

	order.Items, err = r.GetItemsByID(ctx, order.OrderUID)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrdersRepo) GetAll(ctx context.Context) (*[]domain.Order, error) {
	builder := sq.Select(orderUidColumn, trackNumberColumn, entryColumn, localeColumn, internalSignatureColumn, customerIdColumn, deliveryServiceColumn, shardkeyColumn, smIdColumn, dateCreatedColumn, oofShardColumn).
		PlaceholderFormat(sq.Dollar).
		From(ordersTableName)

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

	for i := range list {
		list[i].Delivery, err = r.GetDeliveryByID(ctx, list[i].OrderUID)
		if err != nil {
			return nil, err
		}

		list[i].Payment, err = r.GetPaymentByID(ctx, list[i].OrderUID)
		if err != nil {
			return nil, err
		}

		list[i].Items, err = r.GetItemsByID(ctx, list[i].OrderUID)
		if err != nil {
			return nil, err
		}
	}

	return &list, nil
}

func (r *OrdersRepo) GetDeliveryByID(ctx context.Context, uid string) (*domain.Delivery, error) {
	builder := sq.Select(nameColumn, phoneColumn, zipColumn, cityColumn, addressColumn, regionColumn, emailColumn).
		PlaceholderFormat(sq.Dollar).
		From(deliveriesTableName).
		Where(sq.Eq{orderUidColumn: uid}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "orders_repository.GetDeliveryByID",
		QueryRaw: query,
	}

	var delivery domain.Delivery

	err = r.db.DB().ScanOneContext(ctx, &delivery, q, args...)
	if err != nil {
		return nil, err
	}

	return &delivery, nil
}

func (r *OrdersRepo) GetPaymentByID(ctx context.Context, uid string) (*domain.Payment, error) {
	builder := sq.Select(transactionColumn, requestIdColumn, currencyColumn, providerColumn, amountColumn, paymentDtColumn, bankColumn, deliveryCostColumn, goodsTotalColumn, customFeeColumn).
		PlaceholderFormat(sq.Dollar).
		From(paymentsTableName).
		Where(sq.Eq{transactionColumn: uid}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "orders_repository.GetPaymentByID",
		QueryRaw: query,
	}

	var payment domain.Payment

	err = r.db.DB().ScanOneContext(ctx, &payment, q, args...)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *OrdersRepo) GetItemsByID(ctx context.Context, uid string) (*[]domain.Item, error) {
	builder := sq.Select(chrtIdColumn, trackNumberColumn, priceColumn, ridColumn, nameProductColumn, saleColumn, sizeColumn, totalPriceColumn, nmIdColumn, brandColumn, statusColumn).
		PlaceholderFormat(sq.Dollar).
		From(itemsTableName).
		Where(sq.Eq{orderUidColumn: uid})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "orders_repository.GetItemsByID",
		QueryRaw: query,
	}

	listItem := make([]domain.Item, 0)

	err = r.db.DB().ScanAllContext(ctx, &listItem, q, args...)
	if err != nil {
		return nil, err
	}

	return &listItem, nil
}

func (r *OrdersRepo) CreateOrder(ctx context.Context, order *domain.Order) error {
	builder := sq.Insert(ordersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(orderUidColumn, trackNumberColumn, entryColumn, localeColumn, internalSignatureColumn, customerIdColumn, deliveryServiceColumn, shardkeyColumn, smIdColumn, dateCreatedColumn, oofShardColumn).
		Values(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "orders_repository.CreateOrder",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	err = r.CreateOrderDelivery(ctx, order)
	if err != nil {
		return err
	}

	err = r.CreateOrderPayment(ctx, order)
	if err != nil {
		return err
	}

	err = r.CreateOrderItem(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrdersRepo) CreateOrderDelivery(ctx context.Context, order *domain.Order) error {
	builder := sq.Insert(deliveriesTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(orderUidColumn, nameColumn, phoneColumn, zipColumn, cityColumn, addressColumn, regionColumn, emailColumn).
		Values(order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "orders_repository.CreateOrderDelivery",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrdersRepo) CreateOrderPayment(ctx context.Context, order *domain.Order) error {
	builder := sq.Insert(paymentsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(transactionColumn, requestIdColumn, currencyColumn, providerColumn, amountColumn, paymentDtColumn, bankColumn, deliveryCostColumn, goodsTotalColumn, customFeeColumn).
		Values(order.OrderUID, order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "orders_repository.CreateOrderPayment",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *OrdersRepo) CreateOrderItem(ctx context.Context, order *domain.Order) error {
	for _, item := range *order.Items {
		builder := sq.Insert(itemsTableName).
			PlaceholderFormat(sq.Dollar).
			Columns(orderUidColumn, chrtIdColumn, trackNumberColumn, priceColumn, ridColumn, nameProductColumn, saleColumn, sizeColumn, totalPriceColumn, nmIdColumn, brandColumn, statusColumn).
			Values(order.OrderUID, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)

		query, args, err := builder.ToSql()
		if err != nil {
			return err
		}

		q := db.Query{
			Name:     "orders_repository.CreateOrderItem",
			QueryRaw: query,
		}

		_, err = r.db.DB().ExecContext(ctx, q, args...)
		if err != nil {
			return err
		}
	}

	return nil
}
