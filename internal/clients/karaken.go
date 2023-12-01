package clients

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kraken"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Kraken struct {
	client *kraken.Client
}

func NewKraken(client *kraken.Client) *Kraken {
	return &Kraken{client: client}
}

func (c *Kraken) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 2500)
	var response kraken.TickerResponse
	err := backoff.Retry(func() error {
		var err error
		response, err = c.client.Ticker(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from kraken")
	}
	currentTime := time.Now()
	for symbolPair, tick := range response.Result {
		price, err := tick.AveragePrice()
		if err != nil {
			zap.L().Error("get average price from kraken", zap.Error(err), zap.Any("tick", tick))
			continue
		}
		result = append(result, domain.SymbolPrice{
			Exchange: "kraken",
			Symbol:   strings.Replace(string(symbolPair), "XBT", "BTC", 1),
			Price:    price,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *Kraken) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
