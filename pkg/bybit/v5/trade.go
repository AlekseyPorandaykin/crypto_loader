package v5

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) TradeSpotOpenOrders(ctx context.Context, apiKey, apiSecret string) (response.TradeOpenOrdersResponse, error) {
	return c.tradeOpenOrders(ctx, apiKey, apiSecret, domain.SpotOrderCategory)
}

func (c *Client) TradeLinearOpenOrders(ctx context.Context, apiKey, apiSecret string) (response.TradeOpenOrdersResponse, error) {
	return c.tradeOpenOrders(ctx, apiKey, apiSecret, domain.LinearOrderCategory)
}

func (c *Client) TradeInverseOpenOrders(ctx context.Context, apiKey, apiSecret string) (response.TradeOpenOrdersResponse, error) {
	return c.tradeOpenOrders(ctx, apiKey, apiSecret, domain.InverseOrderCategory)
}

func (c *Client) TradeOptionOpenOrders(ctx context.Context, apiKey, apiSecret string) (response.TradeOpenOrdersResponse, error) {
	return c.tradeOpenOrders(ctx, apiKey, apiSecret, domain.OptionOrderCategory)
}

func (c *Client) tradeOpenOrders(ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory) (response.TradeOpenOrdersResponse, error) {
	req, err := c.traderRequest.GetOpenOrders(ctx, apiKey, apiSecret, category)
	if err != nil {
		return response.TradeOpenOrdersResponse{}, err
	}

	resp, err := c.sender.Send(req)
	if err != nil {
		return response.TradeOpenOrdersResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	result := response.TradeOpenOrdersResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.TradeOpenOrdersResponse{}, err
	}
	return result, err
}
func (c *Client) TradeSpotOrderHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeOrderHistoryResponse, error) {
	return c.TradeOrderHistory(ctx, cred, request.TradeOrderHistoryParam{Category: domain.SpotOrderCategory})
}

func (c *Client) TradeLinearOrderHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeOrderHistoryResponse, error) {
	return c.TradeOrderHistory(ctx, cred, request.TradeOrderHistoryParam{Category: domain.LinearOrderCategory})
}

func (c *Client) TradeInverseOrderHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeOrderHistoryResponse, error) {
	return c.TradeOrderHistory(ctx, cred, request.TradeOrderHistoryParam{Category: domain.InverseOrderCategory})
}
func (c *Client) TradeOptionOrderHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeOrderHistoryResponse, error) {
	return c.TradeOrderHistory(ctx, cred, request.TradeOrderHistoryParam{Category: domain.OptionOrderCategory})
}

func (c *Client) TradeOrderHistory(
	ctx context.Context, cred request.CredentialParam, param request.TradeOrderHistoryParam,
) (response.TradeOrderHistoryResponse, error) {
	req, err := c.traderRequest.GetOrderHistory(ctx, cred, param)
	if err != nil {
		return response.TradeOrderHistoryResponse{}, err
	}

	resp, err := c.sender.Send(req)
	if err != nil {
		return response.TradeOrderHistoryResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	result := response.TradeOrderHistoryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.TradeOrderHistoryResponse{}, err
	}
	return result, nil
}

func (c *Client) TradeSpotHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeHistoryResponse, error) {

	return c.TradeHistory(ctx, cred, request.TradeHistoryParam{Category: domain.SpotOrderCategory})
}

func (c *Client) TradeLinearHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeHistoryResponse, error) {
	return c.TradeHistory(ctx, cred, request.TradeHistoryParam{Category: domain.LinearOrderCategory})
}

func (c *Client) TradeInverseHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeHistoryResponse, error) {
	return c.TradeHistory(ctx, cred, request.TradeHistoryParam{Category: domain.InverseOrderCategory})
}

// TradeOptionHistory - Error execute (404)
func (c *Client) TradeOptionHistory(
	ctx context.Context, cred request.CredentialParam,
) (response.TradeHistoryResponse, error) {
	return c.TradeHistory(ctx, cred, request.TradeHistoryParam{Category: domain.OptionOrderCategory})
}

func (c *Client) TradeHistory(
	ctx context.Context, cred request.CredentialParam, param request.TradeHistoryParam,
) (response.TradeHistoryResponse, error) {
	req, err := c.traderRequest.GetTradeHistory(ctx, cred, param)
	if err != nil {
		return response.TradeHistoryResponse{}, WrapCreateRequestErr(err)
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return response.TradeHistoryResponse{}, err
	}

	defer func() { _ = resp.Body.Close() }()
	result := response.TradeHistoryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.TradeHistoryResponse{}, err
	}
	return result, err
}
