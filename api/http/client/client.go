package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

type SymbolSnapshotCandlestick struct {
	Symbol       string    `json:"symbol"`
	Exchange     string    `json:"exchange"`
	OpenTime     time.Time `json:"open_time"`
	CloseTime    time.Time `json:"close_time"`
	OpenPrice    float64   `json:"open_price"`
	HighPrice    float64   `json:"high_price"`
	LowPrice     float64   `json:"low_price"`
	ClosePrice   float64   `json:"close_price"`
	Volume       float64   `json:"volume"`
	NumberTrades int       `json:"number_trades"`
	Interval     string    `json:"interval"`
	CreatedAt    time.Time `json:"created_at"`
}

type SymbolSnapshotResponse struct {
	Symbol        string                    `json:"symbol"`
	Exchange      string                    `json:"exchange"`
	CreatedAt     time.Time                 `json:"created_at"`
	Price         string                    `json:"price"`
	Candlestick4H SymbolSnapshotCandlestick `json:"candlestick_4h"`
	Candlestick1H SymbolSnapshotCandlestick `json:"candlestick_1h"`
}

func DefaultClient() *Client {
	c, err := NewClient("http://localhost:8081", http.DefaultClient)
	if err != nil {
		zap.L().Panic("error init crypto_loader client", zap.Error(err))
	}
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

func (c *Client) SymbolSnapshot(ctx context.Context, exchange, symbol string) (SymbolSnapshotResponse, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/snapshot/%s/%s", c.hostUrl.String(), exchange, symbol),
		nil)
	if err != nil {
		return SymbolSnapshotResponse{}, errors.Wrap(err, "crate price request to symbol snapshot")
	}
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return SymbolSnapshotResponse{}, errors.Wrap(err, "send request to symbol snapshot")
	}
	defer func() { _ = resp.Body.Close() }()
	var res SymbolSnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return SymbolSnapshotResponse{}, errors.Wrap(err, "parse response body from symbol snapshot")
	}
	return res, nil
}

func (c *Client) OneHourCandlesticks(ctx context.Context, exchange, symbol string) ([]SymbolSnapshotCandlestick, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/candlesticks/1h/%s/%s", c.hostUrl.String(), exchange, symbol),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "crate price request to symbol snapshot")
	}
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send request to symbol snapshot")
	}
	defer func() { _ = resp.Body.Close() }()
	var res []SymbolSnapshotCandlestick
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, "parse response body from symbol snapshot")
	}
	return res, nil
}
func (c *Client) FourHourCandlesticks(ctx context.Context, exchange, symbol string) ([]SymbolSnapshotCandlestick, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/candlesticks/4h/%s/%s", c.hostUrl.String(), exchange, symbol),
		nil)
	if err != nil {
		return nil, errors.Wrap(err, "crate price request to symbol snapshot")
	}
	req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send request to symbol snapshot")
	}
	defer func() { _ = resp.Body.Close() }()
	var res []SymbolSnapshotCandlestick
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, "parse response body from symbol snapshot")
	}
	return res, nil
}
