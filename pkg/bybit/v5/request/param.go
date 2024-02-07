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
