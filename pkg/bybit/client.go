package bybit

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	host       *url.URL
}

func NewClient(host string) (*Client, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		httpClient: http.DefaultClient,
		host:       urlHost,
	}, nil
}

func (c *Client) WithHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) SpotTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	urlReq := c.host.JoinPath("/v5/market/tickers")
	q := urlReq.Query()
	q.Set("category", "spot")
	urlReq.RawQuery = q.Encode()
	req, err := http.NewRequest(
		http.MethodGet,
		urlReq.String(),
		nil,
	)
	if err != nil {
		return result, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return TickerResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return result, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, errors.Wrap(err, "error decode response")
	}

	return result, nil
}
