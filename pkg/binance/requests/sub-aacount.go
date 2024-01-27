package requests

import (
	"net/http"
	"net/url"
)

type SubAccount struct {
	host *url.URL
}

func NewSubAccount(host string) (*SubAccount, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &SubAccount{
		host: urlHost,
	}, nil
}

func (r *SubAccount) GetDetailOnSubAccount(apiKey, secretKey string) (*http.Request, int, error) {
	req, err := createSecurityRequest(
		apiKey, secretKey, http.MethodGet, r.host.JoinPath("/sapi/v1/sub-account/margin/account").String(), nil,
	)
	return req, 10, err
}

func (r *SubAccount) QuerySubAccountList(apiKey, secretKey string) (*http.Request, int, error) {
	req, err := createSecurityRequest(
		apiKey, secretKey, http.MethodGet, r.host.JoinPath("/sapi/v1/sub-account/list").String(), nil,
	)
	return req, 1, err
}
