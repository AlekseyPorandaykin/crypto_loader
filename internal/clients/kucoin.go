package clients

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/response"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Kucoin struct {
	client *kucoin.Client
}

func NewKucoin(client *kucoin.Client) *Kucoin {
	return &Kucoin{
		client: client,
	}
}

func (c *Kucoin) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 1500)
	var allTickersResp response.AllTickersResponse
	err := backoff.Retry(func() error {
		var err error
		allTickersResp, err = c.client.GetAllTickers(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from kucoin")
	}
	if !allTickersResp.IsOk() {
		return nil, fmt.Errorf("incorrect response from kucoin: %s", allTickersResp.Code)
	}
	currentTime := time.UnixMilli(allTickersResp.Data.Time)
	if currentTime.Year() != time.Now().Year() {
		currentTime = time.Now()
	}
	for _, item := range allTickersResp.Data.Ticker {
		result = append(result, domain.SymbolPrice{
			Exchange: "kukoin",
			Symbol:   strings.Replace(item.Symbol, "-", "", 1),
			Price:    item.LastPrice,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *Kucoin) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
