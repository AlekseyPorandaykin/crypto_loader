package domain

import (
	"context"
	"time"
)

type SymbolPrice struct {
	Exchange string    `json:"exchange" db:"exchange"`
	Symbol   string    `json:"symbol" db:"symbol"`
	Price    string    `json:"price" db:"price"`
	Date     time.Time `json:"date" db:"datetime"`
}

type AfgCoefficient struct {
	Date      time.Time `json:"date" db:"datetime"`
	Symbol    string    `json:"symbol" db:"symbol"`
	Exchange  string    `json:"exchange" db:"exchange"`
	AfgValue  int64     `json:"afgValue" db:"afg_value"`
	Price     float64   `json:"price" db:"price"`
	PrevPrice float64   `json:"prevPrice" db:"prev_price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type PriceStorage interface {
	SavePrices(ctx context.Context, prices []SymbolPrice) error
	LastPrices(ctx context.Context) ([]SymbolPrice, error)
}
