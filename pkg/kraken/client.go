package kraken

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

func (c *Client) Ticker(ctx context.Context) (TickerResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.host.JoinPath("/0/public/Ticker").String(),
		nil,
	)
	if err != nil {
		return TickerResponse{}, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return TickerResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return TickerResponse{}, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	var result TickerResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return TickerResponse{}, errors.Wrap(err, "error decode response")
	}
	return result, nil
}
