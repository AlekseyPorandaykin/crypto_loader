package bitget

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

func (c *Client) GetTicker(ctx context.Context) (TickersResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.host.JoinPath("/api/v2/spot/market/tickers").String(),
		nil,
	)
	if err != nil {
		return TickersResponse{}, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return TickersResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return TickersResponse{}, errors.New("empty body response")
	}
	var result TickersResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return TickersResponse{}, errors.Wrap(err, "error decode response")
	}

	return result, nil
}
