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
var intervalBetweenRequests = time.Second / 10

var limitRequests = 10
var criticalLimitRequests = 5

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

// Basic - проверяет лимиты по заголовкам и ждет, если лимит превышен. Ничего не знаем про тело ответа.
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
		s.addWaitInterval(11 * time.Minute)
		return nil, errors.New("status forbidden")
	}
	limitStatusStr := resp.Header.Get("X-Bapi-Limit-Status")
	limitStr := resp.Header.Get("X-Bapi-Limit")
	limitStatus, _ := strconv.Atoi(limitStatusStr) //your remaining requests for current endpoint
	limit, _ := strconv.Atoi(limitStr)             //your current limit for current endpoint
	if limitStatus < limitRequests && limitStatusStr != "" {
		now := time.Now().In(time.UTC)
		nextRequest := now.Add(1 * time.Second)
		resetTimestamp, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit-Reset-Timestamp")) //the timestamp indicating when your request limit resets if you have exceeded your rate_limit. Otherwise, this is just the current timestamp.
		if resetTimestamp > 0 {
			nextRequest = time.UnixMilli(int64(resetTimestamp)).In(time.UTC)
		}
		diff := nextRequest.Sub(now)
		waitDuration := time.Duration(limitRequests-limitStatus) * diff
		errorMessage := "allowed request to bybit limit is less than threshold"

		if limitStatus < criticalLimitRequests {
			errorMessage = "allowed request to bybit limit is less than critical threshold"
			waitDuration = 30 * time.Second
		}
		s.logger.Warn(
			errorMessage,
			zap.String("url", req.URL.String()),
			zap.Int("limitStatus", limitStatus),
			zap.Int("limit", limit),
			zap.Int("resetTimestamp", resetTimestamp),
			zap.Duration("diff", diff),
			zap.Duration("wait", waitDuration),
			zap.Any("headers", resp.Header),
		)
		s.addWaitInterval(waitDuration)
	}
	if limitStatusStr == "" {
		s.addWaitInterval(1 * time.Second)
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
	s.logger.Debug("add wait interval", zap.String("wait", dur.String()))
	s.allowedNextExecute = time.Now().Add(dur)
}
