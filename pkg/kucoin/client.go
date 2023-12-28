package kucoin

import (
	"context"
	"encoding/json"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/sender"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Client struct {
	sender  sender.Sender
	market  *request.Market
	account *request.Account
}

func NewClient(host string, sender sender.Sender) (*Client, error) {
	market, err := request.NewMarket(host)
	if err != nil {
		return nil, err
	}

	account, err := request.NewAccount(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		sender:  sender,
		market:  market,
		account: account,
	}, nil
}

func (c *Client) GetAllTickers(ctx context.Context) (response.AllTickersResponse, error) {
	data, err := c.do(c.market.AllTickers(ctx))
	if err != nil {
		return response.AllTickersResponse{}, err
	}
	result := response.AllTickersResponse{}
	if err := json.Unmarshal(data, &result); err != nil {
		return response.AllTickersResponse{}, errors.Wrap(err, "error decode response")
	}

	return result, nil
}

func (c *Client) GetAccountSummaryInfo(ctx context.Context, cred request.Credential) (response.UserInfoResponse, error) {
	data, err := c.do(c.account.GetAccountSummaryInfo(ctx, &cred))
	if err != nil {
		return response.UserInfoResponse{}, err
	}
	result := response.UserInfoResponse{}
	if err := json.Unmarshal(data, &result); err != nil {
		return response.UserInfoResponse{}, errors.Wrap(err, "error decode response")
	}

	return result, nil
}

func (c *Client) GetAccountList(ctx context.Context, cred request.Credential) (response.AccountListResponse, error) {
	data, err := c.do(c.account.GetAccountList(ctx, &cred))
	if err != nil {
		return response.AccountListResponse{}, err
	}
	result := response.AccountListResponse{}
	if err := json.Unmarshal(data, &result); err != nil {
		return response.AccountListResponse{}, errors.Wrap(err, "error decode response")
	}

	return result, nil
}

func (c *Client) do(req *http.Request, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, err
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
