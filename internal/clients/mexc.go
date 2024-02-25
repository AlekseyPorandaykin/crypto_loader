package clients

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/mexc"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"time"
)

type Mexc struct {
	client *mexc.Client
}

func NewMexc(client *mexc.Client) *Mexc {
	return &Mexc{client: client}
}

func (c *Mexc) LoadPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 2000)
	var priceSymbols []mexc.PriceSymbol
	err := backoff.Retry(func() error {
		var err error
		priceSymbols, err = c.client.SymbolPriceTicker(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "error get price from mexc")
	}
	currentTime := time.Now()
	for _, price := range priceSymbols {
		result = append(result, domain.SymbolPrice{
			Exchange: "mexc",
			Symbol:   price.Symbol,
			Price:    price.Price,
			Date:     currentTime,
		})
	}
	return result, nil
}

func (c *Mexc) CreateFutureOrder(cred domain.ExchangeCredential, order domain.FutureOrder) ([]dto.OrderDTO, error) {
	return nil, nil
}
