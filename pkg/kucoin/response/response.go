package response

type CommonResponse struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

func (r CommonResponse) IsOk() bool {
	return r.Code == "200000"
}

type Paginator struct {
	CurrentPage uint        `json:"currentPage"`
	PageSize    uint        `json:"pageSize"`
	TotalNum    uint        `json:"totalNum"`
	TotalPage   uint        `json:"totalPage"`
	Data        interface{} `json:"data"`
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

type UserInfo struct {
	Level                 int `json:"level"`
	SubQuantity           int `json:"subQuantity"`
	SpotSubQuantity       int `json:"spotSubQuantity"`
	MarginSubQuantity     int `json:"marginSubQuantity"`
	FuturesSubQuantity    int `json:"futuresSubQuantity"`
	MaxSubQuantity        int `json:"maxSubQuantity"`
	MaxDefaultSubQuantity int `json:"maxDefaultSubQuantity"`
	MaxSpotSubQuantity    int `json:"maxSpotSubQuantity"`
	MaxMarginSubQuantity  int `json:"maxMarginSubQuantity"`
	MaxFuturesSubQuantity int `json:"maxFuturesSubQuantity"`
}

type UserInfoResponse struct {
	CommonResponse
	Data UserInfo `json:"data"`
}

type Account struct {
	ID        string `json:"id"`
	Currency  string `json:"currency"`
	Type      string `json:"type"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

type AccountList []Account

type AccountListResponse struct {
	CommonResponse
	Data AccountList `json:"data"`
}
