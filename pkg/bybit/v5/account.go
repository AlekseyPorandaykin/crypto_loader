package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"github.com/pkg/errors"
	"io"
)

func (c *Client) AccountWalletBalance(ctx context.Context, apiKey, apiSecret string, account domain.AccountType) (any, error) {
	req, err := c.accountRequest.GetWalletBalance(ctx, apiKey, apiSecret, account)
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

func (c *Client) AccountTransactionLog(
	ctx context.Context, cred request.CredentialParam, param request.AccountTransactionLogParam,
) (response.TransactionLogsResponse, error) {
	req, err := c.accountRequest.GetTransactionLog(ctx, cred, param)
	if err != nil {
		return response.TransactionLogsResponse{}, WrapErrCreateRequest(err)
	}
	res := response.TransactionLogsResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return response.TransactionLogsResponse{}, err
	}

	return res, err
}
