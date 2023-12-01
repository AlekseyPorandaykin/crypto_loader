package clients

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	binance_domain "github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/domain"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"time"
)

type Binance struct {
	client *binance.Manager
}

func NewBinance(client *binance.Manager) *Binance {
	return &Binance{client: client}
}

func (c *Binance) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 2500)
	var binancePrices []binance_domain.PriceSymbolDTO
	err := backoff.Retry(func() error {
		var err error
		binancePrices, err = c.client.GetPrice(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from binance")
	}
	currentTime := time.Now()
	for _, price := range binancePrices {
		result = append(result, domain.SymbolPrice{
			Exchange: "binance",
			Symbol:   price.Symbol,
			Price:    price.Price,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *Binance) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	var (
		mainSide, secondSide binance_domain.OrderSide
		ordersCreated        = make([]dto.OrderDTO, 0)
	)
	switch order.Side {
	case domain.BuySideOrder:
		mainSide = binance_domain.BuyOrderSide
		secondSide = binance_domain.SellOrderSide
	case domain.SellSideOrder:
		mainSide = binance_domain.SellOrderSide
		secondSide = binance_domain.BuyOrderSide
	default:
		return nil, errors.New("not found order side")
	}
	binanceCred := binance_domain.CredentialDTO{APIKey: cred.APIKey, ApiSecret: cred.ApiSecret}
	if _, err := c.client.FutureChangeInitialLeverage(binanceCred, order.Symbol, order.Leverage); err != nil {
		return nil, err
	}

	mainOrder := binance_domain.NewLimitFutureOrder(
		order.Symbol,
		mainSide,
		binance_domain.GtcTimeInForce,
		order.Quantity,
		order.Price,
		false,
	)

	mainOrder.PositionSide = binance_domain.BothPositionSide
	var orders []binance_domain.FutureOrder

	takeProfitOrder := binance_domain.NewTakeProfitMarketFutureOrder(
		order.Symbol, secondSide, order.Quantity, order.TakeProfit, true,
	)
	takeProfitOrder.PositionSide = binance_domain.BothPositionSide
	takeProfitOrder.TimeInForce = binance_domain.GtcTimeInForce
	takeProfitOrder.WorkingType = binance_domain.MarkPriceWorkingType
	takeProfitOrder.PriceProtect = true

	stopMarketOrder := binance_domain.NewStopMarketFutureOrder(
		order.Symbol, secondSide, order.Quantity, order.StopLoss, true,
	)
	stopMarketOrder.PositionSide = binance_domain.BothPositionSide
	stopMarketOrder.TimeInForce = binance_domain.GtcTimeInForce
	stopMarketOrder.WorkingType = binance_domain.MarkPriceWorkingType
	stopMarketOrder.PriceProtect = true

	orders = append(orders, mainOrder, takeProfitOrder, stopMarketOrder)

	binanceOrders, err := c.client.FuturesPlaceMultipleOrders(binanceCred, orders)
	if err != nil {
		return nil, errors.Wrap(err, "error create second_orders")
	}
	for _, binanceOrder := range binanceOrders {
		ordersCreated = append(ordersCreated, dto.OrderDTO{
			ID:         binanceOrder.OrderId,
			Symbol:     binanceOrder.Symbol,
			Status:     binanceOrder.Status,
			Price:      binanceOrder.Price,
			StopPrice:  binanceOrder.StopPrice,
			Quantity:   binanceOrder.OrigQty,
			Type:       binanceOrder.Type,
			UpdateTime: binanceOrder.UpdateTime,
			ExternalID: binanceOrder.ClientOrderId,
			Raw:        binanceOrder,
		})
	}

	return ordersCreated, nil
}
