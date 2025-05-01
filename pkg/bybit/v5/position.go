package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) PositionInfo(ctx context.Context, cred request.CredentialParam, param request.PositionInfoParam) (response.PositionInfoResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.positionReq.GetPositionInfo(ctx, cred, param)
	if err != nil {
		return response.PositionInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.PositionInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.PositionInfoResponse{}, err
	}

	return result, err
}
func (c *Client) PositionMoveHistory(ctx context.Context, cred request.CredentialParam, param request.MovePositionHistoryParam) (any, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.positionReq.GetMovePositionHistory(ctx, cred, param)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	result := make(map[string]any)
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) PositionClosedPnL(ctx context.Context, cred request.CredentialParam, param request.ClosedPnlParam) (response.ClosedPnlResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
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
