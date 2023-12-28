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
	return &Market{
		host: urlHost,
	}, nil
}

func (r *Market) AllTickers(ctx context.Context) (*http.Request, error) {
	return createRequest(ctx, CreateApiRequest(
		http.MethodGet,
		r.host,
		"/api/v1/market/allTickers",
		nil,
	))
}
