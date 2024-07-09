package request

import (
	"context"
	"net/http"
	"net/url"
)

type User struct {
	host *url.URL
}

func NewUser(host *url.URL) *User {
	return &User{host: host}
}

func (r *User) GetUIDWalletType(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{Url: r.host.JoinPath("/v5/user/get-member-type").String(), Method: http.MethodGet},
		apiKey,
		apiSecret,
	)
}
func (r *User) GetApiKey(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(
		ctx,
		Request{Url: r.host.JoinPath("/v5/user/query-api").String(), Method: http.MethodGet},
		apiKey,
		apiSecret,
	)
}
