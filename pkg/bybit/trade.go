package bybit

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/response"
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
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeOrderHistoryResponse, error) {
	return c.tradeOrderHistory(ctx, apiKey, apiSecret, domain.SpotOrderCategory)
}

func (c *Client) TradeLinearOrderHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeOrderHistoryResponse, error) {
	return c.tradeOrderHistory(ctx, apiKey, apiSecret, domain.LinearOrderCategory)
}

func (c *Client) TradeInverseOrderHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeOrderHistoryResponse, error) {
	return c.tradeOrderHistory(ctx, apiKey, apiSecret, domain.InverseOrderCategory)
}
func (c *Client) TradeOptionOrderHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeOrderHistoryResponse, error) {
	return c.tradeOrderHistory(ctx, apiKey, apiSecret, domain.OptionOrderCategory)
}

func (c *Client) tradeOrderHistory(
	ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory,
) (response.TradeOrderHistoryResponse, error) {
	req, err := c.traderRequest.GetOrderHistory(ctx, apiKey, apiSecret, category)
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
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeHistoryResponse, error) {
	return c.tradeHistory(ctx, apiKey, apiSecret, domain.SpotOrderCategory)
}

func (c *Client) TradeLinearHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeHistoryResponse, error) {
	return c.tradeHistory(ctx, apiKey, apiSecret, domain.LinearOrderCategory)
}

func (c *Client) TradeInverseHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeHistoryResponse, error) {
	return c.tradeHistory(ctx, apiKey, apiSecret, domain.InverseOrderCategory)
}

// TradeOptionHistory - Error execute (404)
func (c *Client) TradeOptionHistory(
	ctx context.Context, apiKey, apiSecret string,
) (response.TradeHistoryResponse, error) {
	return c.tradeHistory(ctx, apiKey, apiSecret, domain.OptionOrderCategory)
}

func (c *Client) tradeHistory(
	ctx context.Context, apiKey, apiSecret string, category domain.OrderCategory,
) (response.TradeHistoryResponse, error) {
	req, err := c.traderRequest.GetTradeHistory(ctx, apiKey, apiSecret, category)
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
