package request

import (
	"context"
	"net/http"
	"net/url"
)

type Market struct {
	host *url.URL
}

func NewMarket(host *url.URL) *Market {
	return &Market{host: host}
}

func (r *Market) GetTickers(ctx context.Context, category string) (*http.Request, error) {
	return createRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/market/tickers").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "category", Value: category}},
	})
}
