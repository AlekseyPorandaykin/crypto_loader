package clients

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/gateio"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type GateIo struct {
	client *gateio.Client
}

func NewGateIo(client *gateio.Client) *GateIo {
	return &GateIo{client: client}
}

func (c *GateIo) LoadPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 5000)
	var prices []gateio.Tick
	err := backoff.Retry(func() error {
		var err error
		prices, err = c.client.Ticker(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from binance")
	}
	currentTime := time.Now()
	for _, price := range prices {
		result = append(result, domain.SymbolPrice{
			Exchange: "gate.io",
			Symbol:   strings.Replace(price.CurrencyPair, "_", "", 1),
			Price:    price.LastTradingPrice,
			Date:     currentTime,
		})
	}

	return result, nil
}

func (c *GateIo) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
