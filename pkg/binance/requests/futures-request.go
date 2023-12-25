package requests

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/domain"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	req, err := createSecurityRequest(
		apiKey, secretKey, http.MethodPost, r.host.JoinPath("/fapi/v1/order").String(), params,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	return req, 0, nil
}

func (r *FutureRequest) ExchangeInformation(ctx context.Context) (*http.Request, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		r.host.JoinPath("/fapi/v1/exchangeInfo").String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 1, nil
}

func (r *FutureRequest) QueryIndexPriceConstituents(ctx context.Context, symbol string) (*http.Request, int, error) {
	urlReq := r.host.JoinPath("/fapi/v1/constituents")
	q := urlReq.Query()
	q.Set("symbol", symbol)
	urlReq.RawQuery = q.Encode()
	req, err := http.NewRequest(
		http.MethodGet,
		urlReq.String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 2, nil
}

func (r *FutureRequest) PlaceMultipleOrders(apiKey, secretKey string, orders []domain.FutureOrder) (*http.Request, int, error) {
	ordersData := make([]map[string]string, 0, len(orders))
	for _, order := range orders {
		ordersData = append(ordersData, order.ToMap())
	}
	batchOrders, err := json.Marshal(ordersData)
	if err != nil {
		return nil, 0, errors.Wrap(err, "can't marshal orders")
	}
	params := make([]param, 0, 10)
	params = append(params, param{key: "batchOrders", value: string(batchOrders)})
	req, err := createSecurityRequest(
		apiKey, secretKey, http.MethodPost, r.host.JoinPath("/fapi/v1/batchOrders").String(), params,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	return req, 5, nil
}

func (r *FutureRequest) ChangeInitialLeverage(apiKey, secretKey string, symbol string, leverage int) (*http.Request, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	params := make([]param, 0, 2)
	params = append(params, param{key: "symbol", value: symbol}, param{key: "leverage", value: strconv.Itoa(leverage)})

	req, err := createSecurityRequest(
		apiKey,
		secretKey,
		http.MethodPost,
		r.host.JoinPath("/fapi/v1/leverage").String(),
		params,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 1, nil
}

func (r *FutureRequest) QueryOrder(
	ctx context.Context, apiKey, secretKey string, symbol string, orderID int,
) (*http.Request, int, error) {
	params := []param{{key: "symbol", value: symbol}, {key: "orderId", value: strconv.Itoa(orderID)}}
	req, err := createSecurityRequest(
		apiKey,
		secretKey,
		http.MethodGet,
		r.host.JoinPath("/fapi/v1/order").String(),
		params,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 2, nil
}

func (r *FutureRequest) CancelMultipleOrders(apiKey, secretKey string, symbol string, orderIDs []int) (*http.Request, int, error) {
	orderIDList, err := json.Marshal(orderIDs)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error decode order_id_list")
	}
	params := []param{{key: "symbol", value: symbol}, {key: "orderIdList", value: string(orderIDList)}}
	req, err := createSecurityRequest(
		apiKey,
		secretKey,
		http.MethodDelete,
		r.host.JoinPath("/fapi/v1/batchOrders").String(),
		params,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	return req, 1, nil
}

func (r *FutureRequest) CandlestickData(ctx context.Context, symbol string, interval domain.CandlestickInterval) (*http.Request, int, error) {
	urlReq := r.host.JoinPath("/fapi/v1/klines")
	q := urlReq.Query()
	q.Set("symbol", symbol)
	q.Set("interval", string(interval))
	q.Set("limit", "100")
	urlReq.RawQuery = q.Encode()
	req, err := http.NewRequest(
		http.MethodGet,
		urlReq.String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 2, nil
}
