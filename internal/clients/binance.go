package clients

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"time"
)

type Binance struct {
	client *binance.Client
}

func NewBinance(client *binance.Client) *Binance {
	return &Binance{client: client}
}

func (c *Binance) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 2500)
	var binancePrices []binance.PriceSymbol
	err := backoff.Retry(func() error {
		var err error
		binancePrices, err = c.client.GetPrice(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from binance")
	}
	currentTime := time.Now()
	for _, price := range binancePrices {
		result = append(result, domain.SymbolPrice{
			Exchange: "binance",
			Symbol:   price.Symbol,
			Price:    price.Price,
			Date:     currentTime,
		})
	}
	return result, nil
}
