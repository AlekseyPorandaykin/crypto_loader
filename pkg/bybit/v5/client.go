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

var (
	ErrCreateRequest = errors.New("error create request")
	ErrApiKeyExpired = errors.New("api key expired")
)

func WrapErrCreateRequest(err error) error {
	return errors.Wrap(err, ErrCreateRequest.Error())
}

var ErrHttpClientDo = errors.New("http client do")

func WrapErrHttpClientDo(err error) error {
	return errors.Wrap(err, ErrCreateRequest.Error())
}

const durationWaitManyVisits = 1 * time.Minute

type Client struct {
	sender sender.Sender
	logger *zap.Logger

	marketRequest  *request.Market
	accountRequest *request.Account
	assetRequest   *request.Asset
	userRequest    *request.User
	traderRequest  *request.Trade
	positionReq    *request.Position

	muExecRequests sync.Mutex

	muAllowedRequests  sync.Mutex
	allowedNextExecute time.Time

	muCreateRequest sync.Mutex
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

func (c *Client) SetPauseDuration(dur time.Duration) {
	c.addWaitInterval(dur)
	c.logger.Debug("set pause duration send requests", zap.String("duration", dur.String()))
}

// sendRequest - dest is pointer struct
func (c *Client) sendRequest(req *http.Request, dest any) error {
	c.executeTimeLimitRequestSafely()
	c.muExecRequests.Lock()
	defer c.muExecRequests.Unlock()
	res, err := c.sender.Send(req)
	if err != nil {
		return WrapErrHttpClientDo(err)
	}
	if res.HttpResp.Body == nil {
		return errors.New("empty body response")
	}
	defer func() { _ = res.HttpResp.Body.Close() }()
	logger := c.logger.With(
		zap.Int("status_code", res.HttpResp.StatusCode),
		zap.Any("wait_duration", res.WaitDuration.String()),
	)
	if req.URL != nil {
		logger = logger.With(zap.String("url", req.URL.String()))
	}
	if err := json.NewDecoder(res.HttpResp.Body).Decode(dest); err != nil {
		return errors.Wrap(err, "error decode response")
	}

	if checker, ok := dest.(response.CheckerResponse); ok && !checker.IsOk() {
		logger = logger.With(zap.Any("response", dest))
		logger = logger.With(zap.Any("headers", res.HttpResp.Header))
		res.AddAction("Error response from bybit", checker.ErrMessage())
		if checker.StatusCode() == response.TooManyVisitsCode {
			res.AddActionWithWait("Too many visits", "", durationWaitManyVisits)
		}
		c.addWaitInterval(res.WaitDuration)
		if checker.StatusCode() == response.ApiKeyHasExpired {
			res.AddAction("(Derivatives) Your api key has expired", "")
			logger.Error("api key expired", zap.Any("actions", res.Actions))
			return ErrApiKeyExpired
		}
		logger.Error("error response from bybit", zap.Any("actions", res.Actions))
		return fmt.Errorf("err message (%s)", checker.ErrMessage())
	}
	logger.Debug("success response from bybit", zap.Any("actions", res.Actions))
	c.addWaitInterval(res.WaitDuration)
	return nil
}

// createRequestSafely - проверки перед созданием запроса
func (c *Client) createRequestSafely() {
	c.executeTimeLimitRequestSafely()
}

func (c *Client) executeTimeLimitRequestSafely() {
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
