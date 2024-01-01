package request

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Market struct {
	host *url.URL
}

func NewMarket(host string) (*Market, error) {
	h, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Market{host: h}, nil
}

func (m *Market) Tickers(ctx context.Context) (*http.Request, error) {
	urlReq := m.host.JoinPath("/api/v5/market/tickers")
	q := urlReq.Query()
	q.Set("instType", "SPOT")
	q.Encode()
	urlReq.RawQuery = q.Encode()
	req, err := http.NewRequest(
		http.MethodGet,
		urlReq.String(),
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "create request")
	}
	req.WithContext(ctx)

	return req, nil
}
