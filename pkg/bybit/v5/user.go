package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) GetUIDWalletType(ctx context.Context, apiKey, apiSecret string) (response.WalletTypeResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
	req, err := c.userRequest.GetUIDWalletType(ctx, apiKey, apiSecret)
	if err != nil {
		return response.WalletTypeResponse{}, err
	}
	var result response.WalletTypeResponse
	if err := c.sendRequest(req, &result); err != nil {
		return response.WalletTypeResponse{}, err
	}
	return result, nil
}

func (c *Client) GetApiKey(ctx context.Context, apiKey, apiSecret string) (response.GetApiKeyInformationResponse, error) {
	c.muCreateRequest.Lock()
	defer c.muCreateRequest.Unlock()
	c.createRequestSafely()
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
