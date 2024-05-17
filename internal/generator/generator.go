package generator

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/mag1c0/L0/internal/domain"
)

func GenerateOrder() domain.Order {
	orderUID := gofakeit.UUID()
	trackNumber := gofakeit.DigitN(15)

	goodsTotal := gofakeit.Number(100, 50000)
	deliveryCost := gofakeit.Number(0, 2000)

	items := make([]domain.Item, 0)
	items = append(items, domain.Item{
		ChrtId:      gofakeit.Int(),
		TrackNumber: trackNumber,
		Price:       goodsTotal,
		Rid:         gofakeit.UUID(),
		Name:        gofakeit.ProductName(),
		Sale:        15,
		Size:        "0",
		TotalPrice:  goodsTotal,
		NmId:        gofakeit.Int(),
		Brand:       gofakeit.Company(),
		Status:      202,
	})

	return domain.Order{
		OrderUID:    orderUID,
		TrackNumber: trackNumber,
		Entry:       "WBIL",
		Delivery: domain.Delivery{
			Name:    gofakeit.Name(),
			Phone:   gofakeit.Phone(),
			Zip:     gofakeit.Address().Zip,
			City:    gofakeit.Address().City,
			Address: gofakeit.Address().Address,
			Region:  gofakeit.Address().Country,
			Email:   gofakeit.Email(),
		},
		Payment: domain.Payment{
			Transaction:  orderUID,
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       goodsTotal + deliveryCost,
			PaymentDt:    gofakeit.Int64(),
			Bank:         "alpha",
			DeliveryCost: deliveryCost,
			GoodsTotal:   goodsTotal,
			CustomFee:    0,
		},
		Items:             nil,
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmId:              "99",
		DateCreated:       gofakeit.Date(),
		OofShard:          "1",
	}
}
