package gateio

type Tick struct {
	CurrencyPair     string `json:"currency_pair"` //Currency pair
	LastTradingPrice string `json:"last"`
	RecentLowestAsk  string `json:"lowest_ask"`
	RecentHighestBid string `json:"highest_bid"`
	ChangePercentage string `json:"change_percentage"` //Change percentage in the last 24h
	ChangeUTC0       string `json:"change_utc0"`       //utc0 timezone, the percentage change in the last 24 hours
	ChangeUTC8       string `json:"change_utc8"`       //utc8 timezone, the percentage change in the last 24 hours
	BaseVolume       string `json:"base_volume"`       //Base currency trade volume in the last 24h
	QuoteVolume      string `json:"quote_volume"`      //Quote currency trade volume in the last 24h
	High24H          string `json:"high_24h"`          //Highest price in 24h
	Low24H           string `json:"low_24h"`           //Lowest price in 24h
}
