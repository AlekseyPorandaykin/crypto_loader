package sender

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// 600 запросов разрешено в 5 секундный интервал, но у некоторых запросов лимит маленький, и можно упереться в него раньше блокера.
// Поэтому ставим лимит по самому крайнему ограничению = 5 запросов в секунду
var intervalBetweenRequests = time.Second / 100

var limitRequests = 10
var criticalLimitRequests = 5

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Basic struct {
	httpClient HTTPDoer
	logger     *zap.Logger

	muAllowedRequests  sync.Mutex
	allowedNextExecute time.Time

	execMu sync.Mutex
}

func NewBasic() *Basic {
	return &Basic{httpClient: http.DefaultClient, logger: zap.NewNop(), allowedNextExecute: time.Now()}
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

func (s *Basic) Send(req *http.Request) (*http.Response, error) {
	s.wait()
	s.execMu.Lock()
	defer s.execMu.Unlock()
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error execute http request")
	}
	s.logger.Debug(
		"sent request",
		zap.String("url", req.URL.String()),
		zap.String("method", req.Method),
		zap.String("status", resp.Status),
		zap.Int("status_code", resp.StatusCode),
		zap.Any("headers", resp.Header),
	)
	if resp.StatusCode == http.StatusForbidden {
		s.logger.Warn("status forbidden", zap.String("url", req.URL.String()))
		s.addWaitInterval(5 * time.Minute)
		return nil, errors.New("status forbidden")
	}
	limitStatus, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit-Status")) //your remaining requests for current endpoint
	limit, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit"))              //your current limit for current endpoint
	if limitStatus < limitRequests {
		now := time.Now().In(time.UTC)
		nextRequest := now.Add(1 * time.Second)
		resetTimestamp, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit-Reset-Timestamp")) //the timestamp indicating when your request limit resets if you have exceeded your rate_limit. Otherwise, this is just the current timestamp.
		if resetTimestamp > 0 {
			nextRequest = time.UnixMilli(int64(resetTimestamp)).In(time.UTC)
		}
		diff := nextRequest.Sub(now)
		waitDuration := time.Duration(limitRequests-limitStatus) * diff
		s.logger.Warn(
			"many requests to bybit",
			zap.String("url", req.URL.String()),
			zap.Int("limitStatus", limitStatus),
			zap.Int("limit", limit),
			zap.Int("resetTimestamp", resetTimestamp),
			zap.Duration("diff", diff),
			zap.Duration("wait", waitDuration),
		)
		if limitStatus < criticalLimitRequests {
			waitDuration = 30 * time.Second
		}
		s.addWaitInterval(waitDuration)
	}
	if resp.StatusCode != http.StatusOK {
		s.logger.Error(
			"incorrect status code",
			zap.String("url", req.URL.String()),
			zap.Int("status", resp.StatusCode),
			zap.String("status_text", resp.Status))
		s.addWaitInterval(30 * time.Second)
		return nil, fmt.Errorf("incorrect status code=%d (%s)", resp.StatusCode, resp.Status)
	}

	return resp, nil
}

func (s *Basic) wait() {
	s.muAllowedRequests.Lock()
	defer s.muAllowedRequests.Unlock()
	diff := s.allowedNextExecute.Sub(time.Now())
	if diff > 0 {
		s.logger.Debug("wait for next execute", zap.Duration("wait", diff))
	}
	if diff <= 0 {
		diff = intervalBetweenRequests
	}
	time.Sleep(diff)
	s.allowedNextExecute = time.Now()
}

func (s *Basic) addWaitInterval(dur time.Duration) {
	s.muAllowedRequests.Lock()
	defer s.muAllowedRequests.Unlock()
	s.logger.Debug("add wait interval", zap.Duration("wait", dur))
	s.allowedNextExecute = time.Now().Add(dur)
}
