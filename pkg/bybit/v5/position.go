package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/pkg/errors"
	"io"
)

func (c *Client) PositionInfo(ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory) (any, error) {
	req, err := c.positionReq.GetPositionInfo(ctx, apiKey, apiSecret, category)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, WrapErrHttpClientDo(err)
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)

	return data, err
}
func (c *Client) PositionMoveHistory(ctx context.Context, cred request.CredentialParam, param request.MovePositionHistoryParam) ([]byte, error) {
	req, err := c.positionReq.GetMovePositionHistory(ctx, cred, param)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, WrapErrHttpClientDo(err)
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)

	return data, err
}
