package domain

import "time"

var SymbolCandlestick = []string{"BTCUSDT", "ETHUSDT", "LTCUSDT", "SOLUSDT"}

type SymbolSnapshot struct {
	Symbol        string      `json:"symbol"`
	Exchange      string      `json:"exchange"`
	CreatedAt     time.Time   `json:"created_at"`
	Price         string      `json:"price"`
	Candlestick4H Candlestick `json:"candlestick_4h"`
	Candlestick1H Candlestick `json:"candlestick_1h"`
}
