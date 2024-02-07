package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"net/http"
	"net/url"
)

type Account struct {
	host *url.URL
}

func NewAccount(host *url.URL) *Account {
	return &Account{host: host}
}

func (r *Account) GetWalletBalance(
	ctx context.Context, apiKey, apiSecret string, account domain.AccountType,
) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/account/wallet-balance").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "accountType", Value: string(account)}},
	}, apiKey, apiSecret)
}

func (r *Account) GetAccountInfo(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/account/info").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}

func (r *Account) GetTransactionLog(
	ctx context.Context, cred CredentialParam, param AccountTransactionLogParam,
) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/account/transaction-log").String(),
		Method: http.MethodGet,
		Params: param.Params(),
	}, cred.ApiKey, cred.ApiSecret)
}
