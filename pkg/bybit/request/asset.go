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

func NewAsset(host string) (*Asset, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Asset{host: urlHost}, nil
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
