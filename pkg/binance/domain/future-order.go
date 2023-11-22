package domain

import "strconv"

type TypeFutureOrder string

const (
	LimitTypeFutureOrder              TypeFutureOrder = "LIMIT"
	MarketTypeFutureOrder             TypeFutureOrder = "MARKET"
	StopTypeFutureOrder               TypeFutureOrder = "STOP"
	TakeProfitTypeFutureOrder         TypeFutureOrder = "TAKE_PROFIT"
	StopMarketTypeFutureOrder         TypeFutureOrder = "STOP_MARKET"
	TakeProfitMarketTypeFutureOrder   TypeFutureOrder = "TAKE_PROFIT_MARKET"
	TrailingStopMarketTypeFutureOrder TypeFutureOrder = "TRAILING_STOP_MARKET"
)

const DefaultTimeInForce = "GTC"

type FutureOrder struct {
	Symbol                  string          `json:"symbol"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"` //Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode.
	Type                    TypeFutureOrder `json:"type"`
	TimeInForce             string          `json:"timeInForce"`
	Quantity                float64         `json:"quantity"`   //Cannot be sent with closePosition=true(Close-All)
	ReduceOnly              bool            `json:"reduceOnly"` //"true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
	Price                   float64         `json:"price"`
	StopPrice               float64         `json:"stopPrice"` //Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ClosePosition           bool            `json:"closePosition"`
	ActivationPrice         float64         `json:"activationPrice"`
	CallbackRate            float64         `json:"callbackRate"`
	WorkingType             string          `json:"workingType"`             //stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE". Default "CONTRACT_PRICE"
	PriceProtect            bool            `json:"priceProtect"`            //"TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	PriceMatch              string          `json:"priceMatch"`              //	only avaliable for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"` //NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire taker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers; default NONE
	GoodTillDate            int64           `json:"goodTillDate"`            //order cancel time for timeInForce GTD, mandatory when timeInforce set to GTD; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
	Timestamp               int64           `json:"timestamp"`
}

func NewLimitFutureOrder(symbol, side, timeInForce string, quantity, price float64) FutureOrder {
	return FutureOrder{
		Symbol:      symbol,
		Side:        side,
		Type:        LimitTypeFutureOrder,
		TimeInForce: timeInForce,
		Quantity:    quantity,
		Price:       price,
	}
}

func NewMarketFutureOrder(symbol, side string, quantity float64) FutureOrder {
	return FutureOrder{
		Symbol:   symbol,
		Side:     side,
		Type:     MarketTypeFutureOrder,
		Quantity: quantity,
	}
}

func NewStopFutureOrder(symbol, side string, quantity, price, stopPrice float64) FutureOrder {
	return FutureOrder{
		Symbol:    symbol,
		Side:      side,
		Type:      StopTypeFutureOrder,
		Quantity:  quantity,
		Price:     price,
		StopPrice: stopPrice,
	}
}

func NewTakeProfitFutureOrder(symbol, side string, quantity, price, stopPrice float64) FutureOrder {
	return FutureOrder{
		Symbol:    symbol,
		Side:      side,
		Type:      TakeProfitTypeFutureOrder,
		Quantity:  quantity,
		Price:     price,
		StopPrice: stopPrice,
	}
}

func NewStopMarketFutureOrder(symbol, side string, quantity, stopPrice float64) FutureOrder {
	return FutureOrder{
		Symbol:    symbol,
		Side:      side,
		Type:      StopMarketTypeFutureOrder,
		StopPrice: stopPrice,
		Quantity:  quantity,
	}
}

func NewTakeProfitMarketFutureOrder(symbol, side string, stopPrice float64) FutureOrder {
	return FutureOrder{
		Symbol:    symbol,
		Side:      side,
		Type:      TakeProfitMarketTypeFutureOrder,
		StopPrice: stopPrice,
	}
}

func NewTrailingStopMarketFutureOrder(symbol, side string, callbackRate float64) FutureOrder {
	return FutureOrder{
		Symbol:       symbol,
		Side:         side,
		Type:         TrailingStopMarketTypeFutureOrder,
		CallbackRate: callbackRate,
	}
}

func (order FutureOrder) ToMap() map[string]string {
	res := make(map[string]string)
	if order.TimeInForce == "" {
		order.TimeInForce = DefaultTimeInForce
	}
	res["symbol"] = order.Symbol
	res["side"] = order.Side
	res["type"] = string(order.Type)
	if order.Quantity > 0 {
		res["quantity"] = floatToString(order.Quantity)
	}
	switch order.Type {
	case LimitTypeFutureOrder:
		res["timeInForce"] = order.TimeInForce
		res["price"] = floatToString(order.Price)
	case StopTypeFutureOrder, TakeProfitTypeFutureOrder:
		res["timeInForce"] = order.TimeInForce
		res["price"] = floatToString(order.Price)
		res["stopPrice"] = floatToString(order.StopPrice)
	case StopMarketTypeFutureOrder, TakeProfitMarketTypeFutureOrder:
		res["stopPrice"] = floatToString(order.StopPrice)
	case TrailingStopMarketTypeFutureOrder:
		res["callbackRate"] = floatToString(order.CallbackRate)
	}
	return res
}

func floatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}
