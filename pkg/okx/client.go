package okx

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

func (c *Client) Tickers(ctx context.Context) (TickersResponse, error) {
	urlReq := c.host.JoinPath("/api/v5/market/tickers")
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
		return TickersResponse{}, errors.Wrap(err, "create request")
	}
	req.WithContext(ctx)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return TickersResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return TickersResponse{}, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	result := TickersResponse{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return TickersResponse{}, errors.Wrap(err, "decode response")
	}
	return result, nil
}
