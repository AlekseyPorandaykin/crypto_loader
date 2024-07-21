package request

import (
	"context"
	"net/http"
	"net/url"
)

type Trade struct {
	host *url.URL
}

func NewTrad(host *url.URL) *Trade {
	return &Trade{host: host}
}

func (r *Trade) GetOpenOrders(
	ctx context.Context, apiKey, apiSecret string, param TradeOpenOrdersParam,
) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/order/realtime").String(),
			Method: http.MethodGet,
			Params: param.Params(),
		},
		apiKey,
		apiSecret,
	)
}

func (r *Trade) GetOrderHistory(
	ctx context.Context, cred CredentialParam, param TradeOrderHistoryParam,
) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/order/history").String(),
			Method: http.MethodGet,
			Params: param.Params(),
		},
		cred.ApiKey,
		cred.ApiSecret,
	)
}

func (r *Trade) GetTradeHistory(
	ctx context.Context, cred CredentialParam, param TradeHistoryParam,
) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/execution/list").String(),
			Method: http.MethodGet,
			Params: param.Params(),
		},
		cred.ApiKey,
		cred.ApiSecret,
	)
}

func (r *Trade) PlaceOrder(ctx context.Context, cred CredentialParam, param PlaceOrderParam) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/order/create").String(),
			Method: http.MethodPost,
			Params: param.Params(),
		},
		cred.ApiKey,
		cred.ApiSecret,
	)
}

func (r *Trade) AmendOrder(ctx context.Context, cred CredentialParam, param AmendOrderParam) (*http.Request, error) {
	return postPersonalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/order/amend").String(),
			Method: http.MethodPost,
			Params: param.Params(),
		},
		cred.ApiKey,
		cred.ApiSecret,
	)
}
