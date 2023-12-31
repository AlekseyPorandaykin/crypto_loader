package clients

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"time"
)

type ByBit struct {
	client *bybit.Client
}

func NewByBit(client *bybit.Client) *ByBit {
	return &ByBit{client: client}
}

func (c *ByBit) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 500)
	var tickerResp bybit.TickerResponse
	err := backoff.Retry(func() error {
		var err error
		tickerResp, err = c.client.SpotTicker(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from bybit")
	}
	if !tickerResp.IsOk() {
		return nil, fmt.Errorf("incorrect response from bybit: %s", tickerResp.Message)
	}
	currentTime := time.UnixMilli(tickerResp.Time)
	if currentTime.Year() != time.Now().Year() {
		currentTime = time.Now()
	}
	for _, price := range tickerResp.Result.List {
		result = append(result, domain.SymbolPrice{
			Exchange: "bybit",
			Symbol:   price.Symbol,
			Price:    price.LastPrice,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *ByBit) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
