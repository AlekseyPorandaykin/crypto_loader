package kucoin

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

func (c *Client) GetAllTickers(ctx context.Context) (AllTickersResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.host.JoinPath("/api/v1/market/allTickers").String(),
		nil,
	)
	if err != nil {
		return AllTickersResponse{}, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return AllTickersResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return AllTickersResponse{}, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	result := AllTickersResponse{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return AllTickersResponse{}, errors.Wrap(err, "error decode response")
	}

	return result, nil
}
