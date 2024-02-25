package domain

import "context"

const (
	Binance = "binance"
	ByBit   = "bybit"
	KuKoin  = "kukoin"
	Okx     = "okx"
	GateIo  = "gate.io"
	Kraken  = "kraken"
	BitGet  = "bitget"
	Mexc    = "mexc"
)

type SymbolInfo struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
}

type ExchangeStorage interface {
	SaveSymbolInfo(ctx context.Context, data []SymbolInfo) error
	InfoBySymbol(ctx context.Context, symbol string) (SymbolInfo, error)
}
