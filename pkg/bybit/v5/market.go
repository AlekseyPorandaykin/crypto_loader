package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
)

func (c *Client) MarketSpotTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	req, err := c.marketRequest.GetTickers(ctx, domain.SpotOrderCategory)
	if err != nil {
		return result, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return TickerResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketLinearTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	req, err := c.marketRequest.GetTickers(ctx, domain.LinearOrderCategory)
	if err != nil {
		return result, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return TickerResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketInverseTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	req, err := c.marketRequest.GetTickers(ctx, domain.InverseOrderCategory)
	if err != nil {
		return result, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return TickerResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketOptionTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	req, err := c.marketRequest.GetTickers(ctx, domain.OptionOrderCategory)
	if err != nil {
		return result, WrapErrCreateRequest(err)
	}
	if err := c.sendRequest(req, &result); err != nil {
		return TickerResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketInstrumentsInfo(ctx context.Context) (response.InstrumentsInfoResponse, error) {
	req, err := c.marketRequest.GetInstrumentsInfo(ctx, domain.SpotOrderCategory)
	if err != nil {
		return response.InstrumentsInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InstrumentsInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InstrumentsInfoResponse{}, err
	}

	return result, err
}
