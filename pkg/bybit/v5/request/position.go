package request

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

type Position struct {
	host *url.URL
}

func NewPosition(host *url.URL) *Position {
	return &Position{host: host}
}

func (p *Position) GetPositionInfo(
	ctx context.Context, cred CredentialParam, param PositionInfoParam,
) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    p.host.JoinPath("/v5/position/list").String(),
		Method: http.MethodGet,
		Params: param.Params(),
	}, cred.ApiKey, cred.ApiSecret)
}

func (p *Position) GetClosedPnL(ctx context.Context, cred CredentialParam, param ClosedPnlParam) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    p.host.JoinPath("/v5/position/closed-pnl").String(),
		Method: http.MethodGet,
		Params: param.Params(),
	}, cred.ApiKey, cred.ApiSecret)
}

func (p *Position) GetMovePositionHistory(
	ctx context.Context, cred CredentialParam, param MovePositionHistoryParam,
) (*http.Request, error) {
	var params []Param
	if param.Category != "" {
		params = append(params, Param{Key: "category", Value: string(param.Category)})
	}
	if param.Symbol != "" {
		params = append(params, Param{Key: "symbol", Value: param.Symbol})
	}
	if !param.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(param.StartTime.UnixMilli()))})
	}
	if !param.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(param.EndTime.UnixMilli()))})
	}
	if param.Status != "" {
		params = append(params, Param{Key: "status", Value: param.Status})
	}
	if param.BlockTradeId != "" {
		params = append(params, Param{Key: "blockTradeId", Value: param.BlockTradeId})
	}
	if param.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(param.Limit)})
	}
	if param.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: param.Cursor})
	}
	return personalRequest(ctx, Request{
		Url:    p.host.JoinPath("/v5/position/move-history").String(),
		Method: http.MethodGet,
	}, cred.ApiKey, cred.ApiSecret)
}
