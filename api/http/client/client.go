package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

type PriceResponse struct {
	Exchange  string    `json:"exchange"`
	Symbol    string    `json:"symbol"`
	Price     string    `json:"price"`
	Date      time.Time `json:"date"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DefaultClient() *Client {
	c, _ := NewClient("http://localhost:8081", http.DefaultClient)
	return c
}

type Client struct {
	client  *http.Client
	hostUrl *url.URL
}

func NewClient(host string, client *http.Client) (*Client, error) {
	hostUrl, err := url.Parse(host)
	if err != nil {
		return nil, errors.Wrap(err, "parse host")
	}
	return &Client{
		hostUrl: hostUrl,
		client:  client,
	}, nil
}

func (c *Client) Prices(ctx context.Context, symbol string) ([]PriceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s/%s", c.hostUrl.String(), "price", symbol), nil)
	if err != nil {
		return nil, errors.Wrap(err, "crate price request")
	}
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send price request")
	}
	defer func() { _ = resp.Body.Close() }()
	var prices []PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, errors.Wrap(err, "parse price response body")
	}
	return prices, nil
}

func (c *Client) AllSymbolPrices(ctx context.Context) ([]PriceResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.hostUrl.String(), "prices"), nil)
	if err != nil {
		return nil, errors.Wrap(err, "crate price request to loader")
	}
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send price request to loader")
	}
	defer func() { _ = resp.Body.Close() }()
	var prices []PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&prices); err != nil {
		return nil, errors.Wrap(err, "parse price response body from loader")
	}
	return prices, nil
}
