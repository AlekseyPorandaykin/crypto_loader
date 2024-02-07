package v5

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/sender"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

var CreateRequestErr = errors.New("error create request")

func WrapCreateRequestErr(err error) error {
	return errors.Wrap(err, CreateRequestErr.Error())
}

type Client struct {
	sender sender.Sender

	marketRequest  *request.Market
	accountRequest *request.Account
	assetRequest   *request.Asset
	userRequest    *request.User
	traderRequest  *request.Trade
	positionReq    *request.Position
}

func DefaultClient(host string) (*Client, error) {
	return NewClient(host, sender.NewBasic())
}

func NewClient(host string, sender sender.Sender) (*Client, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		marketRequest:  request.NewMarket(urlHost),
		accountRequest: request.NewAccount(urlHost),
		assetRequest:   request.NewAsset(urlHost),
		userRequest:    request.NewUser(urlHost),
		traderRequest:  request.NewTrad(urlHost),
		positionReq:    request.NewPosition(urlHost),
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

// sendRequest - dest is pointer struct
func (c *Client) sendRequest(req *http.Request, dest any) error {
	res, err := c.sender.Send(req)
	if err != nil {
		return errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	if err := json.NewDecoder(res.Body).Decode(dest); err != nil {
		return errors.Wrap(err, "error decode response")
	}

	if checker, ok := dest.(response.CheckerResponse); ok && !checker.IsOk() {
		return fmt.Errorf("err message (%s)", checker.ErrMessage())
	}
	return nil
}
