package request

import (
	"context"
	"net/http"
	"net/url"
)

type User struct {
	host *url.URL
}

func NewUser(host string) (*User, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &User{host: urlHost}, nil
}

func (r *User) GetUIDWalletType(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{Url: r.host.JoinPath("/v5/user/get-member-type").String(), Method: http.MethodGet},
		apiKey,
		apiSecret,
	)
}
