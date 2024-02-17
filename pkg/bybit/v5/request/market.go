package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"net/http"
	"net/url"
)

type Market struct {
	host *url.URL
}

func NewMarket(host *url.URL) *Market {
	return &Market{host: host}
}

func (r *Market) GetTickers(ctx context.Context, category domain.OrderCategory) (*http.Request, error) {
	return createRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/market/tickers").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "category", Value: string(category)}},
	})
}

func (r *Market) GetInstrumentsInfo(ctx context.Context, category domain.OrderCategory) (*http.Request, error) {
	return createRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/market/instruments-info").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "category", Value: string(category)}},
	})
}
