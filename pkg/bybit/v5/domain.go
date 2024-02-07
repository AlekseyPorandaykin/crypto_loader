package v5

type CommonResponse struct {
	//https://bybit-exchange.github.io/docs/v5/error#uma--uta--futures-of-classic-account
	Code       int         `json:"retCode"`    //Success/Error code
	Message    string      `json:"retMsg"`     //Success/Error msg. OK, success, SUCCESS indicate a successful response
	Result     interface{} `json:"result"`     //Business data result
	ExtendInfo interface{} `json:"retExtInfo"` //Extend info. Most of the time, it is {}
	Time       int64       `json:"time"`       //Current timestamp (ms)
}

func (r CommonResponse) IsOk() bool {
	return r.Code == 0
}

type SymbolTick struct {
	Symbol        string `json:"symbol"`
	Bid1Price     string `json:"bid1Price"`
	Bid1Size      string `json:"bid1Size"`
	Ask1Price     string `json:"ask1Price"`
	Ask1Size      string `json:"ask1Size"`
	LastPrice     string `json:"lastPrice"`
	PrevPrice24h  string `json:"prevPrice24h"`
	Price24hPcnt  string `json:"price24hPcnt"`
	HighPrice24h  string `json:"highPrice24h"`
	LowPrice24h   string `json:"lowPrice24h"`
	Turnover24h   string `json:"turnover24h"`
	Volume24h     string `json:"volume24h"`
	UsdIndexPrice string `json:"usdIndexPrice"`
}

type TickerResult struct {
	Category string       `json:"category"`
	List     []SymbolTick `json:"list"`
}

type TickerResponse struct {
	CommonResponse
	Result TickerResult `json:"result"`
}
