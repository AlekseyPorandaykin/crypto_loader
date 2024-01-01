package sender

import (
	"net/http"
)

type Sender interface {
	Send(r *http.Request) (*http.Response, error)
}

type Basic struct {
	httpClient *http.Client
}

func New() *Basic {
	return &Basic{httpClient: http.DefaultClient}
}

func (b *Basic) WithHttpClient(httpClient *http.Client) {
	b.httpClient = httpClient
}

func (b *Basic) Send(req *http.Request) (*http.Response, error) {
	res, err := b.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
