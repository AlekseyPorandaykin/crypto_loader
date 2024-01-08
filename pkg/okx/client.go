package okx

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx/sender"
	"github.com/pkg/errors"
)

type Client struct {
	market         *request.Market
	funding        *request.Funding
	tradingAccount *request.TradingAccount

	sender sender.Sender
}

func NewClient(host string) (*Client, error) {
	market, err := request.NewMarket(host)
	if err != nil {
		return nil, err
	}
	funding, err := request.NewFunding(host)
	if err != nil {
		return nil, err
	}
	tradingAccount, err := request.NewTradingAccount(host)
	if err != nil {
		return nil, err
	}

	return &Client{
		market:         market,
		funding:        funding,
		tradingAccount: tradingAccount,
		sender:         sender.New(),
	}, nil
}

func (c *Client) Tickers(ctx context.Context) (response.TickersResponse, error) {
	req, err := c.market.Tickers(ctx)
	if err != nil {
		return response.TickersResponse{}, err
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return response.TickersResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return response.TickersResponse{}, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	result := response.TickersResponse{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return response.TickersResponse{}, errors.Wrap(err, "decode response")
	}
	return result, nil
}

func (c *Client) TradingAccountBalance(
	ctx context.Context, cred request.Credential,
) (response.TradingAccountBalanceResponse, error) {
	req, err := c.tradingAccount.GetBalance(ctx, cred)
	if err != nil {
		return response.TradingAccountBalanceResponse{}, err
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return response.TradingAccountBalanceResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	balance := response.TradingAccountBalanceResponse{}
	if errDecode := json.NewDecoder(resp.Body).Decode(&balance); errDecode != nil {
		return response.TradingAccountBalanceResponse{}, errDecode
	}
	return balance, nil
}

func (c *Client) FundingBalance(ctx context.Context, cred request.Credential) (response.FundingBalanceResponse, error) {
	req, err := c.funding.GetBalance(ctx, cred)
	if err != nil {
		return response.FundingBalanceResponse{}, err
	}
	resp, err := c.sender.Send(req)
	if err != nil {
		return response.FundingBalanceResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	var balance = response.FundingBalanceResponse{}
	if errDecode := json.NewDecoder(resp.Body).Decode(&balance); errDecode != nil {
		return response.FundingBalanceResponse{}, errDecode
	}
	return balance, nil
}
