package dto

import "github.com/google/uuid"

type FutureOrder struct {
	Symbol     string  `json:"symbol"`
	Quantity   float64 `json:"quantity"`
	Price      float64 `json:"price"`
	TakeProfit float64 `json:"take_profit"`
	StopLoss   float64 `json:"stop_loss"`
	Side       string  `json:"side"`
	Type       string  `json:"type"`
	Leverage   int     `json:"leverage"`
}

type ExchangeCredential struct {
	Exchange  string    `json:"exchange"`
	UserUID   uuid.UUID `json:"user_uid"`
	APIKey    string    `json:"api_key"`
	ApiSecret string    `json:"api_secret"`
}

type FutureOrderRequest struct {
	Exchanges   []ExchangeCredential `json:"exchanges"`
	FutureOrder FutureOrder          `json:"future_order"`
}

type OrderDTO struct {
	ID         int         `json:"id"`
	Symbol     string      `json:"symbol"`
	Status     string      `json:"status"`
	Price      string      `json:"price"`
	StopPrice  string      `json:"stopPrice"`
	Quantity   string      `json:"quantity"`
	Type       string      `json:"type"`
	UpdateTime int64       `json:"updateTime"`
	ExternalID string      `json:"external_id"`
	Raw        interface{} `json:"raw"`
}

type CreateOrderDTO struct {
	SourceOrder  FutureOrder `json:"source_order"`
	CreatedOrder []OrderDTO  `json:"created_order"`
	Exchange     string      `json:"exchange"`
	Errors       []error     `json:"errors"`
}
