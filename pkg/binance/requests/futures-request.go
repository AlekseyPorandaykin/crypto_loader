package requests

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/domain"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type FutureRequest struct {
	host *url.URL
}

func NewFutureRequest(host string) (*FutureRequest, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &FutureRequest{
		host: urlHost,
	}, nil
}

func (r *FutureRequest) SymbolPriceTicker(ctx context.Context) (*http.Request, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		r.host.JoinPath("/fapi/v1/ticker/price").String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 2, nil
}

func (r *FutureRequest) NewOrder(apiKey, secretKey string, order domain.FutureOrder) (*http.Request, int, error) {
	params := make([]param, 0, 10)
	for key, val := range order.ToMap() {
		params = append(params, param{key: key, value: val})
	}
	req, err := createSecurityRequest(apiKey, secretKey, http.MethodPost, r.host.JoinPath("/fapi/v1/order").String(), params)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	return req, 2, nil
}
