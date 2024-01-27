package bybit

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/response"
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
