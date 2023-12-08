package dto

type Candlestick struct {
	Symbol       string  `json:"symbol"`
	Exchange     string  `json:"exchange"`
	OpenTime     string  `json:"open_time"`
	CloseTime    string  `json:"close_time"`
	OpenPrice    float64 `json:"open_price"`
	HighPrice    float64 `json:"high_price"`
	LowPrice     float64 `json:"low_price"`
	ClosePrice   float64 `json:"close_price"`
	Volume       float64 `json:"volume"`
	NumberTrades int     `json:"number_trades"`
	Interval     string  `json:"interval"`
	CreatedAt    string  `json:"created_at"`
}

type SymbolSnapshot struct {
	Symbol        string      `json:"symbol"`
	Exchange      string      `json:"exchange"`
	CreatedAt     string      `json:"created_at"`
	Price         string      `json:"price"`
	Candlestick4H Candlestick `json:"candlestick_4h"`
	Candlestick1H Candlestick `json:"candlestick_1h"`
}
