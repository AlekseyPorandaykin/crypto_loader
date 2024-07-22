package v5

import (
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/sender"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

var ErrCreateRequest = errors.New("error create request")

func WrapErrCreateRequest(err error) error {
	return errors.Wrap(err, ErrCreateRequest.Error())
}

var ErrHttpClientDo = errors.New("http client do")

func WrapErrHttpClientDo(err error) error {
	return errors.Wrap(err, ErrCreateRequest.Error())
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

// sendRequest - dest is pointer struct
func (c *Client) sendRequest(req *http.Request, dest any) error {
	res, err := c.sender.Send(req)
	if err != nil {
		return WrapErrHttpClientDo(err)
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
