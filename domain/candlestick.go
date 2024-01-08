package domain

import "time"

type Candlestick struct {
	Symbol       string              `json:"symbol"`
	Exchange     string              `json:"exchange"`
	OpenTime     time.Time           `json:"open_time"`
	CloseTime    time.Time           `json:"close_time"`
	OpenPrice    float64             `json:"open_price"`
	HighPrice    float64             `json:"high_price"`
	LowPrice     float64             `json:"low_price"`
	ClosePrice   float64             `json:"close_price"`
	Volume       float64             `json:"volume"`
	NumberTrades int                 `json:"number_trades"`
	Interval     CandlestickInterval `json:"interval"`
	CreatedAt    time.Time           `json:"created_at"`
}

type CandlestickInterval string

const (
	OneHourCandlestickInterval  CandlestickInterval = "1h"
	FourHourCandlestickInterval CandlestickInterval = "4h"
)

func HasCandlestickInterval(interval CandlestickInterval) bool {
	switch interval {
	case OneHourCandlestickInterval, FourHourCandlestickInterval:
		return true
	default:
		return false
	}
}
