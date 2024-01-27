package request

import (
	"context"
	"net/http"
	"net/url"
)

type Account struct {
	host *url.URL
}

func NewAccount(host string) (*Account, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Account{host: urlHost}, nil
}

func (r *Account) GetWalletBalance(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/account/wallet-balance").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "accountType", Value: "UNIFIED"}},
	}, apiKey, apiSecret)
}

func (r *Account) GetAccountInfo(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/account/info").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}
