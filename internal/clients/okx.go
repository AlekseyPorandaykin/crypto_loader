package clients

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Okx struct {
	client *okx.Client
}

func NewOkx(client *okx.Client) *Okx {
	return &Okx{client: client}
}

func (c *Okx) Load(ctx context.Context) ([]domain.SymbolPrice, error) {
	result := make([]domain.SymbolPrice, 0, 500)
	var tickerResp okx.TickersResponse
	err := backoff.Retry(func() error {
		var err error
		tickerResp, err = c.client.Tickers(ctx)
		if err != nil {
			return err
		}
		return nil
	}, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, errors.Wrap(err, "get price from okx")
	}
	if !tickerResp.IsOk() {
		return nil, fmt.Errorf("incorrect response from okx: %s", tickerResp.Message)
	}
	currentTime := time.Now()
	for _, item := range tickerResp.Data {
		result = append(result, domain.SymbolPrice{
			Exchange: "okx",
			Symbol:   strings.Replace(item.InstrumentID, "-", "", 1),
			Price:    item.LastTradedPrice,
			Date:     currentTime,
		})
	}
	return result, nil
}
