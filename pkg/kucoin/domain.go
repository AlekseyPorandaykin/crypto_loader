package kucoin

type CommonResponse struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func (r CommonResponse) IsOk() bool {
	return r.Code == "200000"
}

type Ticker struct {
	Symbol           string `json:"symbol"`           // symbol
	SymbolName       string `json:"symbolName"`       // Name of trading pairs, it would change after renaming
	Buy              string `json:"buy"`              // bestAsk
	Sell             string `json:"sell"`             // bestBid
	ChangeRate       string `json:"changeRate"`       // 24h change rate
	ChangePrice      string `json:"changePrice"`      // 24h change price
	High             string `json:"high"`             // 24h highest price
	Low              string `json:"low"`              // 24h lowest price
	Volume24         string `json:"vol"`              // 24h volume，the aggregated trading volume in BTC
	Volume24Total    string `json:"volValue"`         // 24h total, the trading volume in quote currency of last 24 hours
	LastPrice        string `json:"last"`             // last price
	AveragePrice     string `json:"averagePrice"`     // 24h average transaction price yesterday
	TakerFeeRate     string `json:"takerFeeRate"`     // Basic Taker Fee
	MakerFeeRate     string `json:"makerFeeRate"`     // Basic Maker Fee
	TakerCoefficient string `json:"takerCoefficient"` // Taker Fee Coefficient
	MakerCoefficient string `json:"makerCoefficient"` // Maker Fee Coefficient
}
type TickersData struct {
	Time   int64    `json:"time"` //timestamp
	Ticker []Ticker `json:"ticker"`
}
type AllTickersResponse struct {
	CommonResponse
	Data TickersData `json:"data"`
}
