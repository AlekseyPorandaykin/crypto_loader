package bybit

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/pkg/errors"
	"io"
)

func (c *Client) PositionInfo(ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory) (any, error) {
	req, err := c.positionReq.GetPositionInfo(ctx, apiKey, apiSecret, category)
	if err != nil {
		return nil, errors.Wrap(err, "error create request")
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)

	return data, err
}
