package v5

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) GetUIDWalletType(ctx context.Context, apiKey, apiSecret string) (response.WalletTypeResponse, error) {
	var result response.WalletTypeResponse
	req, err := c.userRequest.GetUIDWalletType(ctx, apiKey, apiSecret)
	if err != nil {
		return response.WalletTypeResponse{}, err
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return response.WalletTypeResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.WalletTypeResponse{}, err
	}
	return result, nil
}

func (c *Client) GetApiKey(ctx context.Context, apiKey, apiSecret string) (response.GetApiKeyInformationResponse, error) {
	req, err := c.userRequest.GetApiKey(ctx, apiKey, apiSecret)
	if err != nil {
		return response.GetApiKeyInformationResponse{}, WrapErrCreateRequest(err)
	}
	res := response.GetApiKeyInformationResponse{}
	if err := c.sendRequest(req, &res); err != nil {
		return response.GetApiKeyInformationResponse{}, err
	}

	return res, err
}
