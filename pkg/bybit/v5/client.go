package v5

import (
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/sender"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"sync"
	"time"
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
	logger *zap.Logger

	marketRequest  *request.Market
	accountRequest *request.Account
	assetRequest   *request.Asset
	userRequest    *request.User
	traderRequest  *request.Trade
	positionReq    *request.Position

	muAllowedRequests  sync.Mutex
	allowedNextExecute time.Time
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
		marketRequest:      request.NewMarket(urlHost),
		accountRequest:     request.NewAccount(urlHost),
		assetRequest:       request.NewAsset(urlHost),
		userRequest:        request.NewUser(urlHost),
		traderRequest:      request.NewTrad(urlHost),
		positionReq:        request.NewPosition(urlHost),
		sender:             sender,
		logger:             zap.NewNop(),
		allowedNextExecute: time.Now(),
	}, nil
}
func (c *Client) WithSender(s sender.Sender) {
	if s == nil {
		return
	}
	c.sender = s
}
func (c *Client) WithLogger(l *zap.Logger) {
	if l == nil {
		return
	}
	c.logger = l
}

// sendRequest - dest is pointer struct
func (c *Client) sendRequest(req *http.Request, dest any) error {
	c.wait()
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
		fields := []zap.Field{
			zap.Any("response", dest),
			zap.Any("headers", res.Header),
			zap.Int("status_code", res.StatusCode),
		}
		if req.URL != nil {
			fields = append(fields, zap.String("url", req.URL.String()))
		}
		if checker.StatusCode() == response.TooManyVisitsCode {
			c.addWaitInterval(1 * time.Minute)
		}
		c.logger.Error("error response from bybit", fields...)
		return fmt.Errorf("err message (%s)", checker.ErrMessage())
	}
	return nil
}

func (c *Client) wait() {
	c.muAllowedRequests.Lock()
	defer c.muAllowedRequests.Unlock()
	diff := c.allowedNextExecute.Sub(time.Now())
	if diff > 0 {
		time.Sleep(diff)
	}
	c.allowedNextExecute = time.Now()
}

func (c *Client) addWaitInterval(dur time.Duration) {
	c.muAllowedRequests.Lock()
	defer c.muAllowedRequests.Unlock()
	c.allowedNextExecute = time.Now().Add(dur)
}
