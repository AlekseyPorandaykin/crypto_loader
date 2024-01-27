package request

import (
	"context"
	"net/http"
	"net/url"
)

type Market struct {
	host *url.URL
}

func NewMarket(host string) (*Market, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Market{host: urlHost}, nil
}

func (r *Market) GetTickers(ctx context.Context, category string) (*http.Request, error) {
	return createRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/market/tickers").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "category", Value: category}},
	})
}
