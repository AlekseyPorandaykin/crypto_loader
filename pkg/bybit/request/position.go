package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"net/http"
	"net/url"
)

type Position struct {
	host *url.URL
}

func NewPosition(host *url.URL) *Position {
	return &Position{host: host}
}

func (p *Position) GetPositionInfo(
	ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory,
) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    p.host.JoinPath("/v5/position/list").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "category", Value: string(category)}},
	}, apiKey, apiSecret)
}
