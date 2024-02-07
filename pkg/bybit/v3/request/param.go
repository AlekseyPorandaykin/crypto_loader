package request

import (
	"strconv"
	"time"
)

type CredentialParam struct {
	ApiKey    string
	ApiSecret string
}

type AssetWithdrawParam struct {
	WithdrawID   int
	TxID         string
	StartTime    time.Time
	EndTime      time.Time
	Coin         string
	WithdrawType string
	Cursor       string
	Limit        int
}

func (p AssetWithdrawParam) Params() []Param {
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
	if p.WithdrawType != "" {
		params = append(params, Param{Key: "withdrawType", Value: p.WithdrawType})
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
