package v5

import (
	"context"
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) AccountWalletBalance(ctx context.Context, apiKey, apiSecret string, account domain.AccountType) (response.AccountWalletBalanceResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
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
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
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
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.accountRequest.GetAccountInfo(ctx, cred.ApiKey, cred.ApiSecret)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	result := make(map[string]any)
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) AccountFreeRate(ctx context.Context, cred request.CredentialParam, params request.AccountFeeRateParam) (response.AccountFeeRateResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	if params.Category == "" {
		return response.AccountFeeRateResponse{}, errors.New("category required")
	}
	if params.BaseCoin != "" && params.Category != domain.OptionOrderCategory {
		return response.AccountFeeRateResponse{}, errors.New("base coin valid only for option category")
	}
	req, err := c.accountRequest.GetFeeRate(ctx, cred, params)
	if err != nil {
		return response.AccountFeeRateResponse{}, WrapErrCreateRequest(err)
	}
	result := response.AccountFeeRateResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.AccountFeeRateResponse{}, err
	}
	return result, nil
}
