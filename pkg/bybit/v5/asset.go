package v5

import (
	"context"
	"errors"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"strings"
)

func (c *Client) SpotAssetInfo(ctx context.Context, apiKey, apiSecret string) ([]response.AccountAssets, error) {
	return c.assetInfo(ctx, apiKey, apiSecret, domain.SpotAccountType)
}

func (c *Client) FundAssetInfo(ctx context.Context, apiKey, apiSecret string) ([]response.AccountAssets, error) {
	return c.assetInfo(ctx, apiKey, apiSecret, domain.FundAccountType)
}

func (c *Client) OptionAssetInfo(ctx context.Context, apiKey, apiSecret string) ([]response.AccountAssets, error) {
	return c.assetInfo(ctx, apiKey, apiSecret, domain.OptionAccountType)
}

func (c *Client) ContractAssetInfo(ctx context.Context, apiKey, apiSecret string) ([]response.AccountAssets, error) {
	return c.assetInfo(ctx, apiKey, apiSecret, domain.ContractAccountType)
}

func (c *Client) UnifiedAssetInfo(ctx context.Context, apiKey, apiSecret string) ([]response.AccountAssets, error) {
	return c.assetInfo(ctx, apiKey, apiSecret, domain.UnifiedAccountType)
}

func (c *Client) assetInfo(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType) ([]response.AccountAssets, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	var result response.AssetResponse
	req, err := c.assetRequest.GetAssetInfo(ctx, apiKey, apiSecret, accountType)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}
	return result.Result[strings.ToLower(string(accountType))].Assets, nil
}

func (c *Client) AssetContractCoinsBalance(
	ctx context.Context, apiKey, apiSecret string,
) (response.CoinBalanceResponse, error) {
	return c.assetCoinsBalance(ctx, apiKey, apiSecret, domain.ContractAccountType, nil)
}

func (c *Client) AssetUnifiedCoinsBalance(
	ctx context.Context, apiKey, apiSecret string, coins []string,
) (response.CoinBalanceResponse, error) {
	if len(coins) == 0 {
		return response.CoinBalanceResponse{}, errors.New("coins required for unified account balance")
	}
	if len(coins) > 10 {
		return response.CoinBalanceResponse{}, errors.New("maximum 10 coins can be requested for unified account balance")
	}
	return c.assetCoinsBalance(ctx, apiKey, apiSecret, domain.UnifiedAccountType, coins)
}

func (c *Client) AssetFundCoinsBalance(
	ctx context.Context, apiKey, apiSecret string,
) (response.CoinBalanceResponse, error) {
	return c.assetCoinsBalance(ctx, apiKey, apiSecret, domain.FundAccountType, nil)
}

func (c *Client) assetCoinsBalance(
	ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType, coins []string,
) (response.CoinBalanceResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	var result response.CoinBalanceResponse
	req, err := c.assetRequest.GetAllCoinsBalance(ctx, apiKey, apiSecret, accountType, coins)
	if err != nil {
		return response.CoinBalanceResponse{}, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return response.CoinBalanceResponse{}, err
	}
	return result, nil
}

func (c *Client) AssetCoinExchangeRecords(ctx context.Context, apiKey, apiSecret string) (any, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetCoinExchangeRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	result := make(map[string]any)
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *Client) AssetInternalTransferRecords(
	ctx context.Context, apiKey, apiSecret string,
) (response.InternalTransferRecordsResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetInternalTransferRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return response.InternalTransferRecordsResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InternalTransferRecordsResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InternalTransferRecordsResponse{}, err
	}

	return result, nil
}
func (c *Client) AssetWithdrawalRecords(
	ctx context.Context, cred request.CredentialParam, param request.AssetWithdrawalRecordsParam,
) (response.WithdrawalRecordsResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetWithdrawalRecords(ctx, cred, param)
	if err != nil {
		return response.WithdrawalRecordsResponse{}, WrapErrCreateRequest(err)
	}
	result := response.WithdrawalRecordsResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.WithdrawalRecordsResponse{}, err
	}

	return result, nil
}
func (c *Client) AssetDepositRecords(ctx context.Context, cred request.CredentialParam, param request.GetDepositRecordParam) (CommonResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetDepositRecords(ctx, cred, param)
	if err != nil {
		return CommonResponse{}, WrapErrCreateRequest(err)
	}
	result := CommonResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return CommonResponse{}, err
	}
	return result, nil
}

func (c *Client) AssetUniversalTransferRecords(ctx context.Context, apiKey, apiSecret string) (any, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetUniversalTransferRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return nil, WrapErrCreateRequest(err)
	}
	result := make(map[string]interface{})
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) AssetCoinInfo(ctx context.Context, apiKey, apiSecret, coin string) (response.CoinInfoResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetCoinInfo(ctx, apiKey, apiSecret, strings.ToUpper(coin))
	if err != nil {
		return response.CoinInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.CoinInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.CoinInfoResponse{}, err
	}

	return result, nil
}
func (c *Client) AssetCoinsInfo(ctx context.Context, apiKey, apiSecret string) (response.CoinInfoResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.assetRequest.GetCoinInfo(ctx, apiKey, apiSecret, "")
	if err != nil {
		return response.CoinInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.CoinInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.CoinInfoResponse{}, err
	}

	return result, nil
}
