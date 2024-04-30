package requests

import (
	"context"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type WalletRequest struct {
	host *url.URL
}

func NewWalletRequest(host string) (*WalletRequest, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &WalletRequest{
		host: urlHost,
	}, nil
}

func (r *WalletRequest) AllCoinsInformation(ctx context.Context) (*http.Request, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		r.host.JoinPath("/sapi/v1/capital/config/getall").String(),
		nil,
	)
	if err != nil {
		return nil, 0, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	return req, 10, nil
}
