package bybit

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/sender"
	"github.com/pkg/errors"
)

type Client struct {
	sender sender.Sender

	marketRequest  *request.Market
	accountRequest *request.Account
	assetRequest   *request.Asset
	userRequest    *request.User
	traderRequest  *request.Trade
}

func DefaultClient(host string) (*Client, error) {
	return NewClient(host, sender.NewBasic())
}

func NewClient(host string, sender sender.Sender) (*Client, error) {
	marketReq, err := request.NewMarket(host)
	if err != nil {
		return nil, err
	}
	accountReq, err := request.NewAccount(host)
	if err != nil {
		return nil, err
	}
	assetReq, err := request.NewAsset(host)
	if err != nil {
		return nil, err
	}
	userReq, err := request.NewUser(host)
	if err != nil {
		return nil, err
	}
	tradeReq, err := request.NewTrad(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		marketRequest:  marketReq,
		accountRequest: accountReq,
		assetRequest:   assetReq,
		userRequest:    userReq,
		traderRequest:  tradeReq,
		sender:         sender,
	}, nil
}
func (c *Client) WithSender(s sender.Sender) {
	if s == nil {
		return
	}
	c.sender = s
}
func (c *Client) SpotTicker(ctx context.Context) (TickerResponse, error) {
	result := TickerResponse{}
	req, err := c.marketRequest.GetTickers(ctx, "spot")
	if err != nil {
		return result, errors.Wrap(err, "error create request")
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return TickerResponse{}, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return result, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, errors.Wrap(err, "error decode response")
	}

	return result, nil
}

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
