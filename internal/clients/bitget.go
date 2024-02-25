package clients

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bitget"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"time"
)

type Bitget struct {
	client *bitget.Client
}

func NewBitget(client *bitget.Client) *Bitget {
	return &Bitget{client: client}
}

func (c *Bitget) LoadPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 500)
	var tickerResp bitget.TickersResponse
	err := backoff.Retry(func() error {
		var err error
		tickerResp, err = c.client.GetTicker(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from bitget")
	}
	if !tickerResp.IsOk() {
		return nil, fmt.Errorf("incorrect response from bitget: %s", tickerResp.Message)
	}
	currentTime := time.UnixMilli(tickerResp.RequestTime)
	if currentTime.Year() != time.Now().Year() {
		currentTime = time.Now()
	}
	for _, tick := range tickerResp.Data {
		result = append(result, domain.SymbolPrice{
			Exchange: "bitget",
			Symbol:   tick.Symbol,
			Price:    tick.LastPrice,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *Bitget) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
