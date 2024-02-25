package request

import (
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"strconv"
	"time"
)

type CredentialParam struct {
	ApiKey    string
	ApiSecret string
}

type GetDepositRecordParam struct {
	Cursor    string
	Coin      string
	StartTime time.Time
	EndTime   time.Time
	Limit     int
}

type MovePositionHistoryParam struct {
	Category     domain.OrderCategory
	Symbol       string
	StartTime    time.Time
	EndTime      time.Time
	Status       string
	BlockTradeId string
	Limit        int
	Cursor       string
}

type TradeOrderHistoryParam struct {
	Category    domain.OrderCategory
	Symbol      string
	BaseCoin    string
	SettleCoin  string
	OrderID     string
	OrderLinkID string
	OrderFilter string
	OrderStatus string
	StartTime   time.Time
	EndTime     time.Time
	Limit       int
	Cursor      string
}

func (p TradeOrderHistoryParam) Params() []Param {
	var params []Param
	params = append(params, Param{Key: "category", Value: string(p.Category)})
	if p.Symbol != "" {
		params = append(params, Param{Key: "symbol", Value: p.Symbol})
	}
	if p.BaseCoin != "" {
		params = append(params, Param{Key: "baseCoin", Value: p.BaseCoin})
	}
	if p.SettleCoin != "" {
		params = append(params, Param{Key: "settleCoin", Value: p.SettleCoin})
	}
	if p.OrderID != "" {
		params = append(params, Param{Key: "orderId", Value: p.OrderID})
	}
	if p.OrderLinkID != "" {
		params = append(params, Param{Key: "orderLinkId", Value: p.OrderLinkID})
	}
	if p.OrderFilter != "" {
		params = append(params, Param{Key: "orderFilter", Value: p.OrderFilter})
	}
	if p.OrderStatus != "" {
		params = append(params, Param{Key: "orderStatus", Value: p.OrderStatus})
	}
	if !p.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(p.StartTime.UnixMilli()))})
	}
	if !p.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(p.EndTime.UnixMilli()))})
	}
	if p.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(p.Limit)})
	}
	if p.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: p.Cursor})
	}

	return params
}

type AssetWithdrawalRecordsParam struct {
	WithdrawID   int
	TxID         string
	StartTime    time.Time
	EndTime      time.Time
	Coin         string
	WithdrawType int
	Cursor       string
	Limit        int
}

func (p AssetWithdrawalRecordsParam) Params() []Param {
	var params []Param
	if p.WithdrawID > 0 {
		params = append(params, Param{Key: "withdrawID", Value: strconv.Itoa(p.WithdrawID)})
	}
	if p.TxID != "" {
		params = append(params, Param{Key: "txID", Value: p.TxID})
	}

	if p.Coin != "" {
		params = append(params, Param{Key: "coin", Value: p.Coin})
	}
	if p.WithdrawType > 0 {
		params = append(params, Param{Key: "withdrawType", Value: strconv.Itoa(p.WithdrawType)})
	}

	if !p.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(p.StartTime.UnixMilli()))})
	}
	if !p.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(p.EndTime.UnixMilli()))})
	}
	if p.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(p.Limit)})
	}
	if p.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: p.Cursor})
	}

	return params
}

type TradeHistoryParam struct {
	Category    domain.OrderCategory
	Symbol      string
	OrderID     string
	OrderLinkID string
	BaseCoin    string
	StartTime   time.Time
	EndTime     time.Time
	ExecType    domain.ExecType
	Limit       int
	Cursor      string
}

func (p TradeHistoryParam) Params() []Param {
	var params []Param
	params = append(params, Param{Key: "category", Value: string(p.Category)})
	if p.Symbol != "" {
		params = append(params, Param{Key: "symbol", Value: p.Symbol})
	}
	if p.OrderID != "" {
		params = append(params, Param{Key: "orderId", Value: p.OrderID})
	}
	if p.OrderLinkID != "" {
		params = append(params, Param{Key: "orderLinkId", Value: p.OrderLinkID})
	}
	if p.BaseCoin != "" {
		params = append(params, Param{Key: "baseCoin", Value: p.BaseCoin})
	}
	if p.ExecType != "" {
		params = append(params, Param{Key: "execType", Value: string(p.ExecType)})
	}
	if !p.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(p.StartTime.UnixMilli()))})
	}
	if !p.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(p.EndTime.UnixMilli()))})
	}
	if p.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(p.Limit)})
	}
	if p.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: p.Cursor})
	}

	return params
}

type AccountTransactionLogParam struct {
	AccountType domain.AccountType
	Category    domain.OrderCategory
	Currency    string
	BaseCoin    string
	Type        string
	StartTime   time.Time
	EndTime     time.Time
	Limit       int
	Cursor      string
}

func (p AccountTransactionLogParam) Params() []Param {
	var params []Param
	if p.AccountType != "" {
		params = append(params, Param{Key: "accountType", Value: string(p.AccountType)})
	}
	if p.Category != "" {
		params = append(params, Param{Key: "category", Value: string(p.Category)})
	}
	if p.Currency != "" {
		params = append(params, Param{Key: "orderId", Value: p.Currency})
	}
	if p.BaseCoin != "" {
		params = append(params, Param{Key: "baseCoin", Value: p.BaseCoin})
	}
	if p.Type != "" {
		params = append(params, Param{Key: "execType", Value: p.Type})
	}
	if !p.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(p.StartTime.UnixMilli()))})
	}
	if !p.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(p.EndTime.UnixMilli()))})
	}
	if p.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(p.Limit)})
	}
	if p.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: p.Cursor})
	}

	return params
}

type PlaceOrderParam struct {
	Category       domain.OrderCategory
	Symbol         string
	IsLeverage     bool
	Side           domain.Side
	OrderType      domain.OrderType
	Qty            float64
	MarketUnit     string
	Price          float64
	OrderFilter    string
	TriggerPrice   string
	TriggerBy      string
	OrderIv        string
	TimeInForce    string
	PositionIdx    string
	OrderLinkId    string
	TakeProfit     string
	StopLoss       string
	TpTriggerBy    string
	SlTriggerBy    string
	ReduceOnly     string //What is a reduce-only order? true means your position can only reduce in size if this order is triggered.
	CloseOnTrigger string //What is a close on trigger order? For a closing order. It can only reduce your position, not increase it. If the account has insufficient available balance when the closing order is triggered, then other active orders of similar contracts will be cancelled or reduced. It can be used to ensure your stop loss reduces your position regardless of current available margin.
	SmpType        string //With the Self Match Prevention function users can choose the execution method when placing an order. When the counterparty is the same UID or belongs to the same specified SMP group, the execution can be effected accordingly
	Mmp            string //Market Maker Protection (MMP) is an automated mechanism designed to protect market makers (MM) against liquidity risks and over-exposure in the market. It prevents simultaneous trade executions on quotes provided by the MM within a short time span. The MM can automatically pull their quotes if the number of contracts traded for an underlying asset exceeds the configured threshold within a certain time frame. Once MMP is triggered, any pre-existing MMP orders will be automatically canceled, and new orders tagged as MMP will be rejected for a specific duration — known as the frozen period — so that MM can reassess the market and modify the quotes.
	TpslMode       string
	TpLimitPrice   string
	SlLimitPrice   string
	TpOrderType    string
	SlOrderType    string
}

func (p PlaceOrderParam) Params() []Param {
	var params []Param
	params = append(params, Param{Key: "category", Value: string(p.Category)})
	params = append(params, Param{Key: "symbol", Value: p.Symbol})
	if p.IsLeverage {
		params = append(params, Param{Key: "isLeverage", Value: "1"})
	} else {
		params = append(params, Param{Key: "isLeverage", Value: "0"})
	}
	params = append(params, Param{Key: "side", Value: string(p.Side)})
	params = append(params, Param{Key: "orderType", Value: string(p.OrderType)})

	params = append(params, Param{Key: "qty", Value: strconv.FormatFloat(p.Qty, 'f', -1, 32)})

	if p.MarketUnit != "" {
		params = append(params, Param{Key: "marketUnit", Value: p.MarketUnit})
	}
	if p.Price > 0 {
		params = append(params, Param{Key: "price", Value: strconv.FormatFloat(p.Price, 'f', -1, 32)})
	}
	if p.OrderFilter != "" {
		params = append(params, Param{Key: "orderFilter", Value: p.OrderFilter})
	}
	if p.TriggerPrice != "" {
		params = append(params, Param{Key: "triggerPrice", Value: p.TriggerPrice})
	}
	if p.TriggerBy != "" {
		params = append(params, Param{Key: "triggerBy", Value: p.TriggerBy})
	}
	if p.OrderIv != "" {
		params = append(params, Param{Key: "orderIv", Value: p.OrderIv})
	}
	if p.TimeInForce != "" {
		params = append(params, Param{Key: "timeInForce", Value: p.TimeInForce})
	}
	if p.PositionIdx != "" {
		params = append(params, Param{Key: "positionIdx", Value: p.PositionIdx})
	}
	if p.OrderLinkId != "" {
		params = append(params, Param{Key: "orderLinkId", Value: p.OrderLinkId})
	}
	if p.TakeProfit != "" {
		params = append(params, Param{Key: "takeProfit", Value: p.TakeProfit})
	}
	if p.StopLoss != "" {
		params = append(params, Param{Key: "stopLoss", Value: p.StopLoss})
	}
	if p.TpTriggerBy != "" {
		params = append(params, Param{Key: "tpTriggerBy", Value: p.TpTriggerBy})
	}
	if p.SlTriggerBy != "" {
		params = append(params, Param{Key: "slTriggerBy", Value: p.SlTriggerBy})
	}
	if p.ReduceOnly != "" {
		params = append(params, Param{Key: "reduceOnly", Value: p.ReduceOnly})
	}
	if p.CloseOnTrigger != "" {
		params = append(params, Param{Key: "closeOnTrigger", Value: p.CloseOnTrigger})
	}
	if p.SmpType != "" {
		params = append(params, Param{Key: "smpType", Value: p.SmpType})
	}
	if p.Mmp != "" {
		params = append(params, Param{Key: "mmp", Value: p.Mmp})
	}
	if p.TpslMode != "" {
		params = append(params, Param{Key: "tpslMode", Value: p.TpslMode})
	}
	if p.TpLimitPrice != "" {
		params = append(params, Param{Key: "tpLimitPrice", Value: p.TpLimitPrice})
	}
	if p.SlLimitPrice != "" {
		params = append(params, Param{Key: "slLimitPrice", Value: p.SlLimitPrice})
	}
	if p.TpOrderType != "" {
		params = append(params, Param{Key: "tpOrderType", Value: p.TpOrderType})
	}
	if p.SlOrderType != "" {
		params = append(params, Param{Key: "slOrderType", Value: p.SlOrderType})
	}

	return params
}
