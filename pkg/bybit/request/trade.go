package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"net/http"
	"net/url"
)

type Trade struct {
	host *url.URL
}

func NewTrad(host string) (*Trade, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Trade{host: urlHost}, nil
}

func (r *Trade) GetOpenOrders(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{Url: r.host.JoinPath("/v5/order/realtime").String(), Method: http.MethodGet},
		apiKey,
		apiSecret,
	)
}

func (r *Trade) GetOrderHistory(
	ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory,
) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/order/history").String(),
			Method: http.MethodGet,
			Params: []Param{{Key: "category", Value: string(category)}},
		},
		apiKey,
		apiSecret,
	)
}

func (r *Trade) GetTradeHistory(
	ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory,
) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{
			Url:    r.host.JoinPath("/v5/execution/list").String(),
			Method: http.MethodGet,
			Params: []Param{{Key: "category", Value: string(category)}},
		},
		apiKey,
		apiSecret,
	)
}
