package domain

type SideOrder string

const (
	BuySideOrder  SideOrder = "BUY"
	SellSideOrder SideOrder = "SELL"
)

type TypeOrder string

const (
	LimitTypeOrder              TypeOrder = "LIMIT"
	MarketTypeOrder             TypeOrder = "MARKET"
	StopTypeOrder               TypeOrder = "STOP"
	StopMarketTypeOrder         TypeOrder = "STOP_MARKET"
	TakeProfitTypeOrder         TypeOrder = "TAKE_PROFIT"
	TakeProfitMarketTypeOrder   TypeOrder = "TAKE_PROFIT_MARKET"
	TrailingStopMarketTypeOrder TypeOrder = "TRAILING_STOP_MARKET"
)

type FutureOrder struct {
	Symbol     string    `json:"symbol"`
	Quantity   float64   `json:"quantity"`
	Price      float64   `json:"price"`
	TakeProfit float64   `json:"take_profit"`
	StopLoss   float64   `json:"stop_loss"`
	Side       SideOrder `json:"side"`
	Type       TypeOrder `json:"type"`
	Leverage   int       `json:"leverage"`
}
