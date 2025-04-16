package sender

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Basic struct {
	httpClient HTTPDoer
	logger     *zap.Logger
}

func NewBasic() *Basic {
	return &Basic{httpClient: http.DefaultClient}
}

func (s *Basic) WithHttpClient(httpClient HTTPDoer) {
	if httpClient == nil {
		return
	}
	s.httpClient = httpClient
}

func (s *Basic) WithLogger(logger *zap.Logger) {
	if logger == nil {
		return
	}
	s.logger = logger
}

func (s *Basic) Send(req *http.Request) (Response, error) {
	start := time.Now()
	resp, err := s.httpClient.Do(req)
	result := NewResponse(resp)
	result.AddAction("Duration http transport", time.Since(start).String())
	return result, err
}
