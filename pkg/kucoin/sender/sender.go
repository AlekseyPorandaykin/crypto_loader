package sender

import (
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/response"
	"net/http"
)

type Sender interface {
	Send(r *http.Request) (*http.Response, response.Error)
}

type Basic struct {
	httpClient *http.Client
}

func New() *Basic {
	return &Basic{httpClient: http.DefaultClient}
}

func (s *Basic) WithHttpClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

func (s *Basic) Send(req *http.Request) (*http.Response, response.Error) {
	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, ExternalError{Err: err, Message: "http client do"}
	}
	if res.Body == nil {
		return nil, ExternalError{Message: "empty body response"}
	}
	if res.StatusCode >= 400 {
		var exErr response.ExchangeError
		defer func() { _ = res.Body.Close() }()
		if err := json.NewDecoder(res.Body).Decode(&exErr); err != nil {
			return nil, ExternalError{Err: err, Message: "parse error body"}
		}
		return nil, ExternalError{Message: "error from exchange", ExchangeError: exErr}
	}

	return res, nil
}
