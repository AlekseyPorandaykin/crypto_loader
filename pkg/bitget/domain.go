package bitget

type CommonResponse struct {
	Code        string      `json:"code"`
	Data        interface{} `json:"data"`
	Message     string      `json:"msg"`
	RequestTime int64       `json:"requestTime"`
}

func (r CommonResponse) IsOk() bool {
	return r.Code == "00000"
}

type Tick struct {
	Symbol       string `json:"symbol"`       //Trading pair
	HighestPrice string `json:"high24h"`      //24h highest price
	Open24H      string `json:"open"`         //24h open price
	LastPrice    string `json:"lastPr"`       //Latest price
	LowestPrice  string `json:"low24h"`       //24h lowest price
	QuoteVolume  string `json:"quoteVolume"`  //Trading volume in quote currency
	BaseVolume   string `json:"baseVolume"`   //Trading volume in base currency
	USDTVolume   string `json:"usdtVolume"`   //Trading volume in USDT
	BidPr        string `json:"bidPr"`        //Bid 1 price
	AskPr        string `json:"askPr"`        //Ask 1 price
	BidSz        string `json:"bidSz"`        //Buying 1 amount
	AskSz        string `json:"askSz"`        //Selling 1 amount
	OpenUTC      string `json:"openUtc"`      //UTC±00:00 Entry price
	Timestamp    string `json:"ts"`           //Current time Unix millisecond timestamp, e.g. 1690196141868
	ChangeUTC24H string `json:"changeUtc24h"` //Change at UTC+0, 0.01 means 1%.
	Change24H    string `json:"change24h"`    //24-hour change, 0.01 means 1%.
}

type TickersResponse struct {
	CommonResponse
	Data []Tick `json:"data"`
}
