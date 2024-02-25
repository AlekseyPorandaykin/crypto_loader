package domain

type CredentialDTO struct {
	APIKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type PriceSymbolDTO struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type PriceChangeStatisticDTO struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

type ResponseOrderDTO struct {
	OrderId                 int    `json:"orderId"`
	Symbol                  string `json:"symbol"`
	Status                  string `json:"status"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	AvgPrice                string `json:"avgPrice"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ReduceOnly              bool   `json:"reduceOnly"`
	ClosePosition           bool   `json:"closePosition"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	OrigType                string `json:"origType"`
	PriceMatch              string `json:"priceMatch"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	GoodTillDate            int64  `json:"goodTillDate"`
	UpdateTime              int64  `json:"updateTime"`
	Time                    int    `json:"time"`
}

type LeverageDTO struct {
	Symbol           string `json:"symbol"`
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
}

type CandlestickBarDTO struct {
	OpenTime                 float64
	OpenPrice                string
	HighPrice                string
	LowPrice                 string
	ClosePrice               string
	Volume                   string
	CloseTime                float64
	QuoteAssetVolume         string
	NumberOfTrades           float64
	TakerBuyBaseAssetVolume  string
	TakerBuyQuoteAssetVolume string
	Ignore                   string
}

type CandlestickInterval string

const (
	OneHourCandlestickInterval  CandlestickInterval = "1h"
	FourHourCandlestickInterval CandlestickInterval = "4h"
)

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

type ExchangeInfoDTO struct {
	Timezone   string              `json:"timezone"`
	ServerTime int                 `json:"serverTime"`
	RateLimits []RateLimit         `json:"rateLimits"`
	Symbols    []SymbolExchangeDTO `json:"symbols"`
}

type SymbolExchangeDTO struct {
	Symbol                          string      `json:"symbol"`
	Status                          string      `json:"status"`
	BaseAsset                       string      `json:"baseAsset"`
	BaseAssetPrecision              int         `json:"baseAssetPrecision"`
	QuoteAsset                      string      `json:"quoteAsset"`
	QuotePrecision                  int         `json:"quotePrecision"`
	QuoteAssetPrecision             int         `json:"quoteAssetPrecision"`
	BaseCommissionPrecision         int         `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int         `json:"quoteCommissionPrecision"`
	OrderTypes                      []string    `json:"orderTypes"`
	IcebergAllowed                  bool        `json:"icebergAllowed"`
	OcoAllowed                      bool        `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool        `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool        `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool        `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool        `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool        `json:"isMarginTradingAllowed"`
	Filters                         interface{} `json:"filters"`
	Permissions                     []string    `json:"permissions"`
	DefaultSelfTradePreventionMode  string      `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string    `json:"allowedSelfTradePreventionModes"`
}
