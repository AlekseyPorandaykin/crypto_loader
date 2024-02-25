package domain

import "time"

var SymbolCandlestick = []string{"BTCUSDT", "ETHUSDT", "LTCUSDT", "SOLUSDT"}

type SymbolSnapshot struct {
	Symbol        string      `json:"symbol"`
	QuoteAsset    string      `json:"quote_asset"`
	BaseAsset     string      `json:"base_asset"`
	Exchange      string      `json:"exchange"`
	CreatedAt     time.Time   `json:"created_at"`
	Price         string      `json:"price"`
	PriceUpdated  time.Time   `json:"price_updated"`
	Candlestick4H Candlestick `json:"candlestick_4h"`
	Candlestick1H Candlestick `json:"candlestick_1h"`
}
