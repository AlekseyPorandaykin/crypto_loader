package request

import (
	"context"
	"net/http"
	"net/url"
)

type Funding struct {
	host *url.URL
}

func NewFunding(host string) (*Funding, error) {
	h, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Funding{host: h}, nil
}

func (f *Funding) GetBalance(ctx context.Context, cred Credential) (*http.Request, error) {
	apiReq := ApiRequest{
		Method:   http.MethodGet,
		Host:     f.host,
		Endpoint: "/api/v5/asset/balances",
	}
	req, err := createRequest(ctx, apiReq)
	if err != nil {
		return nil, err
	}
	return signRequest(cred, apiReq, req)
}
