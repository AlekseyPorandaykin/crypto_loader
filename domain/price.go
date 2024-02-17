package domain

import (
	"context"
	"time"
)

type SymbolPrice struct {
	Exchange  string    `json:"exchange" db:"exchange"`
	Symbol    string    `json:"symbol" db:"symbol"`
	Price     string    `json:"price" db:"price"`
	Date      time.Time `json:"date" db:"datetime"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type PriceStorage interface {
	SavePrices(ctx context.Context, prices []SymbolPrice) error
	LastPrices(ctx context.Context) ([]SymbolPrice, error)
	SymbolPrice(ctx context.Context, symbol string) ([]SymbolPrice, error)
	ExchangePrice(ctx context.Context, exchange string) ([]SymbolPrice, error)
}
