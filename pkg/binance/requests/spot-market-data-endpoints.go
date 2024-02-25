package requests

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func (c *SpotRequest) SymbolPriceTicker(ctx context.Context) (*http.Request, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.host.JoinPath("/api/v3/ticker/price").String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 4, nil
}
func (c *SpotRequest) ExchangeInformation(ctx context.Context) (*http.Request, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.host.JoinPath("/api/v3/exchangeInfo").String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 4, nil
}

func (c *SpotRequest) RollingWindowPriceChangeStatistics(
	ctx context.Context, symbols []string, widowSize string,
) (*http.Request, int, error) {
	if len(symbols) >= 100 {
		return nil, 0, errors.New("the maximum number of symbols allowed in a request is 100")
	}

	urlReq := c.host.JoinPath("/api/v3/ticker")
	symbolsReq, err := json.Marshal(symbols)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error decode symbols")
	}
	q := urlReq.Query()
	q.Set("symbols", string(symbolsReq))
	q.Set("windowSize", widowSize)
	urlReq.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, urlReq.String(), nil)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)

	return req, 4 * len(symbols), nil
}
