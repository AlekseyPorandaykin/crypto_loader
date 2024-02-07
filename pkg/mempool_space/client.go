package mempool_space

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	host       string
}

func NewClient(host string) *Client {
	return &Client{httpClient: http.DefaultClient, host: host}
}

func (c *Client) WithHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) Price(ctx context.Context) (PriceDTO, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/prices", c.host), nil)
	if err != nil {
		return PriceDTO{}, errors.Wrap(err, "create request")
	}
	rawData := make(map[string]float64)
	if err := c.sendRequest(req, &rawData); err != nil {
		return PriceDTO{}, err
	}
	currencies := make(map[string]float64)
	for key, val := range rawData {
		if key == "time" {
			continue
		}
		currencies[key] = val
	}
	return PriceDTO{
		CreatedAt:  time.Unix(int64(rawData["time"]), 0),
		Currencies: currencies,
	}, nil
}

func (c *Client) RecommendedFees(ctx context.Context) (RecommendedFeesDTO, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/fees/recommended", c.host), nil)
	if err != nil {
		return RecommendedFeesDTO{}, errors.Wrap(err, "create request")
	}
	result := RecommendedFeesDTO{}
	if err := c.sendRequest(req, &result); err != nil {
		return RecommendedFeesDTO{}, err
	}
	return result, nil
}

func (c *Client) sendRequest(req *http.Request, dest any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "execute request")
	}
	defer func() { _ = resp.Body.Close() }()
	if dest == nil {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		return errors.Wrap(err, "parse response")
	}
	return nil
}
