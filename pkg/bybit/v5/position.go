package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"github.com/pkg/errors"
	"io"
)

func (c *Client) PositionInfo(ctx context.Context, cred request.CredentialParam, param request.PositionInfoParam) (response.PositionInfoResponse, error) {
	req, err := c.positionReq.GetPositionInfo(ctx, cred, param)
	if err != nil {
		return response.PositionInfoResponse{}, WrapErrCreateRequest(err)
	}
	if err != nil {
		return response.PositionInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.PositionInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.PositionInfoResponse{}, err
	}

	return result, err
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

func (c *Client) PositionClosedPnL(ctx context.Context, cred request.CredentialParam, param request.ClosedPnlParam) (response.ClosedPnlResponse, error) {
	req, err := c.positionReq.GetClosedPnL(ctx, cred, param)
	if err != nil {
		return response.ClosedPnlResponse{}, WrapErrCreateRequest(err)
	}
	result := response.ClosedPnlResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.ClosedPnlResponse{}, err
	}

	return result, err
}
