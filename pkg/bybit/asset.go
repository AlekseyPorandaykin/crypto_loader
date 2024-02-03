package bybit

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/response"
	"github.com/pkg/errors"
	"io"
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
	var result response.AssetResponse
	req, err := c.assetRequest.GetAssetInfo(ctx, apiKey, apiSecret, accountType)
	if err != nil {
		return nil, err
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Result[strings.ToLower(string(accountType))].Assets, nil
}

func (c *Client) coinsBalance(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType) (response.CoinBalanceResponse, error) {
	var result response.CoinBalanceResponse
	req, err := c.assetRequest.GetAllCoinsBalance(ctx, apiKey, apiSecret, accountType)
	if err != nil {
		return response.CoinBalanceResponse{}, err
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return response.CoinBalanceResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.CoinBalanceResponse{}, err
	}
	return result, nil
}

func (c *Client) AssetCoinExchangeRecords(ctx context.Context, apiKey, apiSecret string) (any, error) {
	req, err := c.assetRequest.GetCoinExchangeRecords(ctx, apiKey, apiSecret)
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
func (c *Client) AssetInternalTransferRecords(
	ctx context.Context, apiKey, apiSecret string,
) (response.InternalTransferRecordsResponse, error) {
	req, err := c.assetRequest.GetInternalTransferRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return response.InternalTransferRecordsResponse{}, WrapCreateRequestErr(err)
	}
	result := response.InternalTransferRecordsResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InternalTransferRecordsResponse{}, err
	}

	return result, nil
}
func (c *Client) AssetWithdrawalRecords(ctx context.Context, apiKey, apiSecret string) (response.WithdrawalRecordsResponse, error) {
	req, err := c.assetRequest.GetWithdrawalRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return response.WithdrawalRecordsResponse{}, errors.Wrap(err, "error create request")
	}
	result := response.WithdrawalRecordsResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.WithdrawalRecordsResponse{}, err
	}

	return result, nil
}

func (c *Client) AssetUniversalTransferRecords(ctx context.Context, apiKey, apiSecret string) (any, error) {
	req, err := c.assetRequest.GetUniversalTransferRecords(ctx, apiKey, apiSecret)
	if err != nil {
		return nil, errors.Wrap(err, "error create request")
	}
	result := make(map[string]interface{})
	if err := c.sendRequest(req, &result); err != nil {
		return nil, err
	}

	return result, nil
}
