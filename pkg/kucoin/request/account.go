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

func (a *Account) GetAccountSummaryInfo(ctx context.Context, c *Credential) (*http.Request, error) {
	apiReq := CreateApiRequest(http.MethodGet, a.host, "/api/v2/user-info", nil)
	req, err := createRequest(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	req, errSign := signRequest(apiReq, c, req)
	if errSign != nil {
		return nil, errSign
	}
	return req, nil
}
func (a *Account) GetAccountList(ctx context.Context, c *Credential) (*http.Request, error) {
	apiReq := CreateApiRequest(http.MethodGet, a.host, "/api/v1/accounts", nil)
	req, err := createRequest(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	req, errSign := signRequest(apiReq, c, req)
	if errSign != nil {
		return nil, errSign
	}
	return req, nil
}
