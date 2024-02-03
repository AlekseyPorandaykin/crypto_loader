package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"net/http"
	"net/url"
)

type Asset struct {
	host *url.URL
}

func NewAsset(host *url.URL) *Asset {
	return &Asset{host: host}
}

func (r *Asset) GetAssetInfo(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-asset-info").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "accountType", Value: string(accountType)}},
	}, apiKey, apiSecret)
}

func (r *Asset) GetAllCoinsBalance(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-account-coins-balance").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "accountType", Value: string(accountType)}},
	}, apiKey, apiSecret)
}
func (r *Asset) GetCoinExchangeRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/exchange/order-record").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}

func (r *Asset) GetInternalTransferRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-inter-transfer-list").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}

func (r *Asset) GetWithdrawalRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/withdraw/query-record").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}
func (r *Asset) GetUniversalTransferRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-universal-transfer-list").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}
