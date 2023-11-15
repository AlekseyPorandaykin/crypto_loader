package client

import (
	"time"
)

type SymbolPrice struct {
	Exchange  string    `json:"exchange" db:"exchange"`
	Symbol    string    `json:"symbol" db:"symbol"`
	Price     float32   `json:"price" db:"price"`
	Date      time.Time `json:"date" db:"datetime"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Batch struct {
	prices []*SymbolPrice
}

func NewBatch(size int) Batch {
	return Batch{prices: make([]*SymbolPrice, 0, size)}
}

func (b *Batch) Append(price *SymbolPrice) {
	b.prices = append(b.prices, price)
}

func (b *Batch) Prices() []*SymbolPrice {
	return b.prices
}
