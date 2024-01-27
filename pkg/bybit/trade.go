package bybit

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/response"
	"io"
)

func (c *Client) TradeOpenOrders(ctx context.Context, apiKey, apiSecret string) (any, error) {
	req, err := c.traderRequest.GetOpenOrders(ctx, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	resp, err := c.sender.Send(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	data, err := io.ReadAll(resp.Body)
	return data, err
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
		return response.TradeHistoryResponse{}, err
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
