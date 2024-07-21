package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"io"
)

func (c *Client) AccountWalletBalance(ctx context.Context, apiKey, apiSecret string, account domain.AccountType) (response.AccountWalletBalanceResponse, error) {
	req, err := c.accountRequest.GetWalletBalance(ctx, apiKey, apiSecret, account)
	if err != nil {
		return response.AccountWalletBalanceResponse{}, WrapErrCreateRequest(err)
	}
	res := response.AccountWalletBalanceResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return response.AccountWalletBalanceResponse{}, err
	}

	return res, err
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

func (c *Client) AccountGetAccountInfo(
	ctx context.Context, cred request.CredentialParam,
) (any, error) {
	req, err := c.accountRequest.GetAccountInfo(ctx, cred.ApiKey, cred.ApiSecret)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	return data, err
}
