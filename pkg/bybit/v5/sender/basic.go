package sender

import (
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Basic struct {
	httpClient *http.Client
	logger     *zap.Logger
}

func NewBasic() *Basic {
	return &Basic{httpClient: http.DefaultClient, logger: zap.NewNop()}
}

func (s *Basic) WithHttpClient(httpClient *http.Client) {
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
	now := time.Now().In(time.UTC)
	if resp.StatusCode == http.StatusForbidden {
		s.logger.Warn("status forbidden", zap.String("url", req.URL.String()))
		time.Sleep(5 * time.Minute)
		return nil, errors.New("status forbidden")
	}
	nextRequest := now.Add(1 * time.Minute)
	limitStatus, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit-Status"))                    //your remaining requests for current endpoint
	limit, _ := strconv.Atoi(resp.Header.Get("X-Bapi-Limit"))                                 //your current limit for current endpoint
	resetTimestamp, errParse := strconv.Atoi(resp.Header.Get("X-Bapi-Limit-Reset-Timestamp")) //the timestamp indicating when your request limit resets if you have exceeded your rate_limit. Otherwise, this is just the current timestamp.
	if errParse == nil {
		nextRequest = time.UnixMilli(int64(resetTimestamp)).In(time.UTC)
	}
	diff := nextRequest.Sub(now)
	if limitStatus < 10 && limit != limitStatus && diff > 0 {
		s.logger.Warn(
			"many requests to bybit",
			zap.String("url", req.URL.String()),
			zap.Int("limitStatus", limitStatus),
			zap.Int("limit", limit),
			zap.Int("resetTimestamp", resetTimestamp),
			zap.Duration("sleep", diff),
		)
		time.Sleep(diff)
	}
	if resp.StatusCode != http.StatusOK {
		s.logger.Error(
			"incorrect status code",
			zap.String("url", req.URL.String()),
			zap.Int("status", resp.StatusCode),
			zap.String("status_text", resp.Status))
		return nil, fmt.Errorf("incorrect status code=%d (%s)", resp.StatusCode, resp.Status)
	}
	return resp, nil
}
