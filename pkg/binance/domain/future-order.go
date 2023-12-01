package domain

import (
	"strconv"
	"time"
)

type OrderSide string

const (
	BuyOrderSide  OrderSide = "BUY"
	SellOrderSide OrderSide = "SELL"
)

type PositionSide string

const (
	BothPositionSide  PositionSide = "BOTH"
	LongPositionSide  PositionSide = "LONG"
	ShortPositionSide PositionSide = "SHORT"
)

type TimeInForce string

const (
	GtcTimeInForce TimeInForce = "GTC" //Good Till Cancel
	IocTimeInForce TimeInForce = "IOC" //Immediate or Cancel
	FokTimeInForce TimeInForce = "FOK" //Fill or Kill
	GtxTimeInForce TimeInForce = "GTX" //Good Till Crossing (Post Only)
	GtdTimeInForce TimeInForce = "GTD" //Good Till Date
)

type WorkingType string

const (
	MarkPriceWorkingType     WorkingType = "MARK_PRICE"
	ContractPriceWorkingType WorkingType = "CONTRACT_PRICE"
)

type TypeFutureOrder string

const (
	LimitTypeFutureOrder              TypeFutureOrder = "LIMIT"
	MarketTypeFutureOrder             TypeFutureOrder = "MARKET"
	StopTypeFutureOrder               TypeFutureOrder = "STOP"
	StopMarketTypeFutureOrder         TypeFutureOrder = "STOP_MARKET"
	TakeProfitTypeFutureOrder         TypeFutureOrder = "TAKE_PROFIT"
	TakeProfitMarketTypeFutureOrder   TypeFutureOrder = "TAKE_PROFIT_MARKET"
	TrailingStopMarketTypeFutureOrder TypeFutureOrder = "TRAILING_STOP_MARKET"
)

type ResponseType string

const (
	AckResponseType    ResponseType = "ACK"
	ResultResponseType ResponseType = "RESULT"
)

/*
TODO
quantity
reduceOnly - Reduce-Only order serves to strictly reduce your open position.
StopPrice - при TP и SL
closePosition
activationPrice
callbackRate
workingType - для stop_price
priceProtect - для stop_price
PriceMatch - установить цену по лучшим ценам в списке (не рекомендуется)
*/

type FutureOrder struct {
	Symbol                  string          `json:"symbol"`
	Side                    OrderSide       `json:"side"`
	PositionSide            PositionSide    `json:"positionSide"` //Default BOTH for One-way Mode ; LONG or SHORT for Hedge Mode. It must be sent in Hedge Mode.
	Type                    TypeFutureOrder `json:"type"`
	TimeInForce             TimeInForce     `json:"timeInForce"`
	Quantity                float64         `json:"quantity"`   //Cannot be sent with closePosition=true(Close-All)
	ReduceOnly              bool            `json:"reduceOnly"` //"true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
	Price                   float64         `json:"price"`
	StopPrice               float64         `json:"stopPrice"` //Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ClosePosition           bool            `json:"closePosition"`
	ActivationPrice         float64         `json:"activationPrice"`
	CallbackRate            float64         `json:"callbackRate"`
	WorkingType             WorkingType     `json:"workingType"`  //stopPrice triggered by: "MARK_PRICE", "CONTRACT_PRICE". Default "CONTRACT_PRICE"
	PriceProtect            bool            `json:"priceProtect"` //"TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	NewOrderRespType        ResponseType    `json:"newOrderRespType"`
	PriceMatch              string          `json:"priceMatch"`              //	only avaliable for LIMIT/STOP/TAKE_PROFIT order; can be set to OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20: /QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20; Can't be passed together with price
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"` //NONE:No STP / EXPIRE_TAKER:expire taker order when STP triggers/ EXPIRE_MAKER:expire taker order when STP triggers/ EXPIRE_BOTH:expire both orders when STP triggers; default NONE
	GoodTillDate            int64           `json:"goodTillDate"`            //order cancel time for timeInForce GTD, mandatory when timeInforce set to GTD; order the timestamp only retains second-level precision, ms part will be ignored; The goodTillDate timestamp must be greater than the current time plus 600 seconds and smaller than 253402300799000
	Timestamp               int64           `json:"timestamp"`
}

var DefaultFutureOrder = FutureOrder{
	PositionSide:     BothPositionSide,
	ReduceOnly:       false,
	WorkingType:      ContractPriceWorkingType,
	NewOrderRespType: AckResponseType,
}

func NewLimitFutureOrder(symbol string, side OrderSide, timeInForce TimeInForce, quantity, price float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:      symbol,
		Side:        side,
		Type:        LimitTypeFutureOrder,
		TimeInForce: timeInForce,
		Quantity:    quantity,
		Price:       price,
		ReduceOnly:  reduceOnly,
	}
}

func NewMarketFutureOrder(symbol string, side OrderSide, quantity float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:     symbol,
		Side:       side,
		Type:       MarketTypeFutureOrder,
		Quantity:   quantity,
		ReduceOnly: reduceOnly,
	}
}

func NewStopFutureOrder(symbol string, side OrderSide, quantity, price, stopPrice float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:      symbol,
		Side:        side,
		Type:        StopTypeFutureOrder,
		Quantity:    quantity,
		Price:       price,
		StopPrice:   stopPrice,
		TimeInForce: GtxTimeInForce,
		ReduceOnly:  reduceOnly,
	}
}

func NewTakeProfitFutureOrder(symbol string, side OrderSide, quantity, price, stopPrice float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:      symbol,
		Side:        side,
		Type:        TakeProfitTypeFutureOrder,
		Quantity:    quantity,
		Price:       price,
		StopPrice:   stopPrice,
		TimeInForce: GtxTimeInForce,
		ReduceOnly:  reduceOnly,
	}
}

func NewStopMarketFutureOrder(symbol string, side OrderSide, quantity, stopPrice float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:     symbol,
		Side:       side,
		Type:       StopMarketTypeFutureOrder,
		StopPrice:  stopPrice,
		Quantity:   quantity,
		ReduceOnly: reduceOnly,
	}
}

func NewTakeProfitMarketFutureOrder(symbol string, side OrderSide, quantity, stopPrice float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:     symbol,
		Side:       side,
		Type:       TakeProfitMarketTypeFutureOrder,
		StopPrice:  stopPrice,
		Quantity:   quantity,
		ReduceOnly: reduceOnly,
	}
}

func NewTrailingStopMarketFutureOrder(symbol string, side OrderSide, callbackRate float64, reduceOnly bool) FutureOrder {
	return FutureOrder{
		Symbol:       symbol,
		Side:         side,
		Type:         TrailingStopMarketTypeFutureOrder,
		CallbackRate: callbackRate,
		ReduceOnly:   reduceOnly,
	}
}

func (order FutureOrder) ToMap() map[string]string {
	res := make(map[string]string)
	res["symbol"] = order.Symbol
	res["side"] = string(order.Side)
	res["type"] = string(order.Type)
	res["priceProtect"] = strconv.FormatBool(order.PriceProtect)
	if order.Quantity > 0 {
		res["quantity"] = floatToString(order.Quantity)
	}
	if order.TimeInForce == GtdTimeInForce {
		order.GoodTillDate = int64(time.Now().Add(30 * 24 * time.Hour).Second())
	}
	switch order.Type {
	case LimitTypeFutureOrder:
		res["timeInForce"] = string(order.TimeInForce)
		res["price"] = floatToString(order.Price)
	case StopTypeFutureOrder, TakeProfitTypeFutureOrder:
		res["timeInForce"] = string(order.TimeInForce)
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
