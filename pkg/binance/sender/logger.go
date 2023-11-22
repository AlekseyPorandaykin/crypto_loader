package sender

import (
	"errors"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Logger struct {
	log    *zap.Logger
	sender Sender
}

func NewLogger(log *zap.Logger, sender Sender) *Logger {
	return &Logger{log: log, sender: sender}
}

func (l *Logger) Send(req *http.Request, weight int) (*http.Response, SenderError) {
	defer func(start time.Time) {
		l.log.Debug(
			"Send request",
			zap.String("duration", time.Since(start).String()),
			zap.Int("weight", weight),
			zap.String("url", req.URL.String()),
		)
	}(time.Now())
	resp, err := l.sender.Send(req, weight)
	if err != nil {
		if errors.Is(err, RateLimitError) || errors.Is(err, IPAutoBannedErr) {
			l.log.Error(
				"lot queries",
				zap.Error(err),
				zap.Any("IpUsedWeight", IpUsedWeight(resp.Header)),
				zap.String("err_detail", err.Detail()),
			)
		} else {
			l.log.Error(
				"execute request",
				zap.Error(err),
				zap.String("err_detail", err.Detail()),
			)
		}
		return nil, err
	}
	return resp, nil
}

func (l *Logger) Close() {
	_ = l.log.Sync()
	l.sender.Close()
}
