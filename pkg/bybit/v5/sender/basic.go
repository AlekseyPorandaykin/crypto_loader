package sender

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Basic struct {
	httpClient HTTPDoer
	logger     *zap.Logger

	mu              sync.Mutex
	counterRequests int
	lastWindow      time.Time
}

func NewBasic() *Basic {
	return &Basic{httpClient: http.DefaultClient, lastWindow: time.Now().Truncate(time.Minute)}
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
	result.AddAction("Current request counter", strconv.Itoa(s.increaseCounter()))
	return result, err
}

func (s *Basic) increaseCounter() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().Truncate(time.Minute)
	if s.lastWindow.Before(now) {
		s.lastWindow = now
		s.counterRequests = 1
	} else {
		s.counterRequests++
	}
	return s.counterRequests
}
