package order

import (
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"strings"
)

type Exchange interface {
	CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error)
}

type Order struct {
	exchanges map[string]Exchange
}

func NewOrder() *Order {
	return &Order{
		exchanges: make(map[string]Exchange),
	}
}

func (o *Order) AddExchange(name string, exchange Exchange) {
	o.exchanges[name] = exchange
}

func (o *Order) CreateFutureOrder(orderReq dto.FutureOrderRequest) ([]dto.CreateOrderDTO, error) {
	var (
		sideOrder domain.SideOrder
		typeOrder domain.TypeOrder
		orders    []dto.CreateOrderDTO
	)
	switch strings.ToUpper(orderReq.FutureOrder.Side) {
	case "BUY":
		sideOrder = domain.BuySideOrder
	case "SELL":
		sideOrder = domain.SellSideOrder
	default:
		return nil, errors.New("not allowed side newOrder")

	}
	switch strings.ToUpper(orderReq.FutureOrder.Type) {
	case "LIMIT":
		typeOrder = domain.LimitTypeOrder
	default:
		return nil, errors.New("not allowed type newOrder")
	}
	newOrder := domain.FutureOrder{
		Symbol:     orderReq.FutureOrder.Symbol,
		Quantity:   orderReq.FutureOrder.Quantity,
		Price:      orderReq.FutureOrder.Price,
		TakeProfit: orderReq.FutureOrder.TakeProfit,
		StopLoss:   orderReq.FutureOrder.StopLoss,
		Side:       sideOrder,
		Type:       typeOrder,
		Leverage:   orderReq.FutureOrder.Leverage,
	}
	for _, exCred := range orderReq.Exchanges {
		exchange, has := o.exchanges[exCred.Exchange]
		if !has {
			continue
		}
		createdOrders, err := exchange.CreateFutureOrder(
			domain.ExchangeCredential{UserUID: exCred.UserUID, APIKey: exCred.APIKey, ApiSecret: exCred.ApiSecret},
			newOrder,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, dto.CreateOrderDTO{
			SourceOrder:  orderReq.FutureOrder,
			CreatedOrder: createdOrders,
			Exchange:     exCred.Exchange,
		})
	}
	return orders, nil
}
