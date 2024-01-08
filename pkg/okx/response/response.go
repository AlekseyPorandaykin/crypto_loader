package response

type CommonResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r CommonResponse) IsOk() bool {
	//https://www.okx.com/docs-v5/en/?shell#financial-product-earn-post-cancel-purchases-redemptions
	return r.Code == "0"
}

type Ticker struct {
	InstrumentType  string `json:"instType"`
	InstrumentID    string `json:"instId"`
	LastTradedPrice string `json:"last"`
	LastTradedSize  string `json:"lastSz"`
	BestAskPrice    string `json:"askPx"`
	BestAskSize     string `json:"askSz"`
	BestBidPrice    string `json:"bidPx"`
	BestBidSize     string `json:"bidSz"`
	Open24H         string `json:"open24h"` //Open price in the past 24 hours
	High24h         string `json:"high24h"` //Highest price in the past 24 hours
	Low24h          string `json:"low24h"`  //Lowest price in the past 24 hours

	//24h trading volume, with a unit of currency.
	//If it is a derivatives contract, the value is the number of base currency.
	//If it is SPOT/MARGIN, the value is the quantity in quote currency.
	VolCcy24h string `json:"volCcy24h"`

	//24h trading volume, with a unit of contract.
	//If it is a derivatives contract, the value is the number of contracts.
	//If it is SPOT/MARGIN, the value is the quantity in base currency.
	Vol24h string `json:"vol24h"`

	SodUtc0   string `json:"sodUtc0"` //Open price in the UTC 0
	SodUtc8   string `json:"sodUtc8"` //Open price in the UTC 8
	Timestamp string `json:"ts"`      //Ticker data generation time, Unix timestamp format in milliseconds, e.g. 1597026383085
}

type TickersResponse struct {
	CommonResponse
	Data []Ticker `json:"data"`
}

type FundingBalance struct {
	AvailableBalance string `json:"availBal"`
	Balance          string `json:"bal"`
	Currency         string `json:"ccy"`
	FrozenBalance    string `json:"frozenBal"`
}

type FundingBalanceResponse struct {
	CommonResponse
	Data []FundingBalance `json:"data"`
}

type TradingAccountBalance struct {
	AdjustedEffectiveEquity           string        `json:"adjEq"`
	PotentialBorrowingIMR             string        `json:"borrowFroz"`
	Details                           []interface{} `json:"details"`
	InitialMarginRequirement          string        `json:"imr"`
	IsolatedMarginEquity              string        `json:"isoEq"`
	MarginRatio                       string        `json:"mgnRatio"`
	MaintenanceMarginRequirement      string        `json:"mmr"`
	NotionalValueOfPositions          string        `json:"notionalUsd"`
	MarginFrozenForPendingCrossOrders string        `json:"ordFroz"`
	TotalAmountEquity                 string        `json:"totalEq"`
	LatestTime                        string        `json:"uTime"`
}

type TradingAccountBalanceResponse struct {
	CommonResponse
	Data []TradingAccountBalance `json:"data"`
}
