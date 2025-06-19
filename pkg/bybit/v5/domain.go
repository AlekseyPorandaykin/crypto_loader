package v5

import "time"

var ExchangeCreated = time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)

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

// SymbolTick - Поля могут отличаться в зависимости от типа акаунта
type SymbolTick struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`              // Последняя торговая цена.
	IndexPrice             string `json:"indexPrice"`             // Индексная цена.
	MarkPrice              string `json:"markPrice"`              // Марк-цена. Используется для расчета P&L и ликвидации.
	PrevPrice24H           string `json:"prevPrice24h"`           // Цена за 24 часа до текущего момента.
	Price24HPcnt           string `json:"price24hPcnt"`           // Процентное изменение цены за 24 часа.
	HighPrice24H           string `json:"highPrice24h"`           // Самая высокая цена за последние 24 часа.
	LowPrice24H            string `json:"lowPrice24h"`            // Самая низкая цена за последние 24 часа.
	PrevPrice1H            string `json:"prevPrice1h"`            // Цена за 1 часа до текущего момента.
	OpenInterest           string `json:"openInterest"`           // Объем открытого интереса (общее количество открытых позиций) для этого символа.
	OpenInterestValue      string `json:"openInterestValue"`      // Стоимость открытого интереса в базовой валюте.
	Turnover24H            string `json:"turnover24h"`            // Общий объем торгов за 24 часа в базовой валюте (например, USD для BTCUSD, USDT для BTCUSDT).
	Volume24H              string `json:"volume24h"`              // Общий объем торгов за 24 часа в контрактах или единицах актива.
	FundingRate            string `json:"fundingRate"`            // Текущая ставка финансирования. Положительное значение означает, что лонги платят шортам; отрицательное — шорты платят лонгам.
	NextFundingTime        string `json:"nextFundingTime"`        // Время следующего обмена финансированием в Unix timestamp (миллисекунды).
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"` // Predicated delivery price. It has a value 30 mins before delivery.
	BasisRate              string `json:"basisRate"`              // Базовая ставка
	DeliveryFeeRate        string `json:"deliveryFeeRate"`        // Ставка комиссии за доставку (актуально для квартальных контрактов).
	DeliveryTime           string `json:"deliveryTime"`           // Время доставки (актуально для квартальных контрактов) в Unix timestamp (миллисекунды).
	Ask1Size               string `json:"ask1Size"`               // Размер лучшего спроса (количество базовой валюты для продажи).
	Bid1Price              string `json:"bid1Price"`              // Лучшая цена предложения (цена покупки).
	Ask1Price              string `json:"ask1Price"`              // Лучшая цена спроса (цена продажи).
	Bid1Size               string `json:"bid1Size"`               // Размер лучшего предложения (количество базовой валюты для покупки).
	Basis                  string `json:"basis"`                  //Basis
	PreOpenPrice           string `json:"preOpenPrice"`           //
	PreQty                 string `json:"preQty"`                 //
}

type TickerResult struct {
	Category string       `json:"category"`
	List     []SymbolTick `json:"list"`
}

type TickerResponse struct {
	CommonResponse
	Result TickerResult `json:"result"`
}
