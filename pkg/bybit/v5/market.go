package v5

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"strconv"
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

func (c *Client) MarketInstrumentsSpotInfo(ctx context.Context) (response.InstrumentsInfoResponse, error) {
	req, err := c.marketRequest.GetInstrumentsInfo(ctx, domain.SpotOrderCategory)
	if err != nil {
		return response.InstrumentsInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InstrumentsInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InstrumentsInfoResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketInstrumentsLinearInfo(ctx context.Context) (response.InstrumentsInfoResponse, error) {
	req, err := c.marketRequest.GetInstrumentsInfo(ctx, domain.LinearOrderCategory)
	if err != nil {
		return response.InstrumentsInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InstrumentsInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InstrumentsInfoResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketInstrumentsInverseInfo(ctx context.Context) (response.InstrumentsInfoResponse, error) {
	req, err := c.marketRequest.GetInstrumentsInfo(ctx, domain.InverseOrderCategory)
	if err != nil {
		return response.InstrumentsInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InstrumentsInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InstrumentsInfoResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketInstrumentsOptionInfo(ctx context.Context) (response.InstrumentsInfoResponse, error) {
	req, err := c.marketRequest.GetInstrumentsInfo(ctx, domain.OptionOrderCategory)
	if err != nil {
		return response.InstrumentsInfoResponse{}, WrapErrCreateRequest(err)
	}
	result := response.InstrumentsInfoResponse{}
	if err := c.sendRequest(req, &result); err != nil {
		return response.InstrumentsInfoResponse{}, err
	}

	return result, nil
}

func (c *Client) MarketGetLinearKlineMonth(ctx context.Context, symbol string) (any, error) {
	return c.MarketGetKline(ctx, request.MarketGetKlineParam{
		Category: domain.LinearOrderCategory,
		Symbol:   symbol,
		Interval: "M",
		Limit:    200,
	})
}

func (c *Client) MarketGetLinearKlineWeek(ctx context.Context, symbol string) (response.GetKlineResponse, error) {
	return c.MarketGetKline(ctx, request.MarketGetKlineParam{
		Category: domain.LinearOrderCategory,
		Symbol:   symbol,
		Interval: "W",
		Limit:    200,
	})
}

func (c *Client) MarketGetLinearKlineDay(ctx context.Context, symbol string) (response.GetKlineResponse, error) {
	return c.MarketGetKline(ctx, request.MarketGetKlineParam{
		Category: domain.LinearOrderCategory,
		Symbol:   symbol,
		Interval: "D",
		Limit:    200,
	})
}
func (c *Client) MarketGetLinearKlineMinute(ctx context.Context, symbol string, interval int) (response.GetKlineResponse, error) {
	return c.MarketGetKline(ctx, request.MarketGetKlineParam{
		Category: domain.LinearOrderCategory,
		Symbol:   symbol,
		Interval: strconv.Itoa(interval),
		Limit:    200,
	})
}

func (c *Client) MarketGetKline(ctx context.Context, param request.MarketGetKlineParam) (response.GetKlineResponse, error) {
	req, err := c.marketRequest.GetKline(ctx, param)
	if err != nil {
		return response.GetKlineResponse{}, WrapErrCreateRequest(err)
	}
	var res response.GetKlineResponse
	if err := c.sendRequest(req, &res); err != nil {
		return response.GetKlineResponse{}, err
	}
	return res, nil
}

func (c *Client) MarketGetOrderBook(ctx context.Context, param request.MarketGetOrderBookParam) (response.GetOrderBookResponse, error) {
	req, err := c.marketRequest.GetOrderBook(ctx, param)
	if err != nil {
		return response.GetOrderBookResponse{}, WrapErrCreateRequest(err)
	}
	var res response.GetOrderBookResponse
	if err := c.sendRequest(req, &res); err != nil {
		return response.GetOrderBookResponse{}, err
	}
	return res, nil
}
