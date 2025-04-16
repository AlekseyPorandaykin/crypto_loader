package sender

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type RequestLogger struct {
	sender Sender
	logger *zap.Logger
}

func NewRequestLogger(sender Sender, logger *zap.Logger) Sender {
	return &RequestLogger{sender: sender, logger: logger}
}

func (r *RequestLogger) Send(req *http.Request) (Response, error) {
	fields := []zap.Field{
		zap.String("method", req.Method),
		zap.String("url", req.URL.String()),
	}
	start := time.Now()
	resp, err := r.sender.Send(req)
	fields = append(fields, zap.String("duration", time.Since(start).String()))
	if resp.HttpResp != nil {
		fields = append(fields, zap.Int("status_code", resp.HttpResp.StatusCode))
		fields = append(fields, zap.String("status", resp.HttpResp.Status))
		fields = append(fields, zap.Any("resp_headers", resp.HttpResp.Header))
	}
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	r.logger.Debug("send request", fields...)
	return resp, err
}
