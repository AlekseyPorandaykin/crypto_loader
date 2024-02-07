package request

import (
	"context"
	"net/http"
	"net/url"
)

type Asset struct {
	host *url.URL
}

func NewAsset(host *url.URL) *Asset {
	return &Asset{host: host}
}

func (r *Asset) GetWithdrawRecords(ctx context.Context, cred CredentialParam, param AssetWithdrawParam) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/asset/v3/private/withdraw/record/query").String(),
		Method: http.MethodGet,
		Params: param.Params(),
	}, cred.ApiKey, cred.ApiSecret)
}
