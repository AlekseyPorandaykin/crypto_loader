package request

import (
	"context"
	"net/http"
	"net/url"
)

type TradingAccount struct {
	host *url.URL
}

func NewTradingAccount(host string) (*TradingAccount, error) {
	h, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &TradingAccount{host: h}, nil
}

func (t *TradingAccount) GetBalance(ctx context.Context, c Credential) (*http.Request, error) {
	apiReq := ApiRequest{
		Method:   http.MethodGet,
		Host:     t.host,
		Endpoint: "/api/v5/account/balance",
	}
	req, err := createRequest(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	return signRequest(c, apiReq, req)
}
