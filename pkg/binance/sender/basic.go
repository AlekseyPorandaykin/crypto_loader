package sender

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const IpWeightLimit = 12000

type Basic struct {
	httpClient *http.Client
	locker     *Locker

	once sync.Once
}

func NewBasic() *Basic {
	return &Basic{
		httpClient: http.DefaultClient,
		locker:     NewLocker(),
	}
}

func (b *Basic) WithHttpClient(httpClient *http.Client) {
	if httpClient == nil {
		return
	}
	b.httpClient = httpClient
}

func (b *Basic) Send(req *http.Request, weight int) (*http.Response, SenderError) {
	b.locker.Lock()
	resp, err := b.httpClient.Do(req)
	b.locker.Unlock()
	if err != nil {
		return nil, WrapErr(err, "error http client do")
	}
	errResp := checkResponse(resp.StatusCode)
	if errResp != nil {
		b.handleResponseError(resp, errResp)
		return resp, errResp
	}
	ipWeight := IpUsedWeight(resp.Header)
	if ipWeight.LastIpWeight >= IpWeightLimit {
		b.locker.AsyncDelay(1 * time.Minute)
	}
	return resp, nil
}

func (b *Basic) handleResponseError(resp *http.Response, err SenderError) {
	if errors.Is(err, RateLimitError) || errors.Is(err, RateLimitError) {
		retryAfter, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		if retryAfter > 0 {
			b.locker.SyncDelay(time.Duration(retryAfter) * time.Second)
		}
		b.locker.SyncDelay(1 * time.Minute)
	}
	var e ExternalError
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		return
	}
	err.WithExternalError(e)
}

func (b *Basic) Close() {
	b.once.Do(func() {
		b.httpClient.CloseIdleConnections()
	})
}

func checkResponse(code int) SenderError {
	if code >= 500 && code < 600 {
		return InternalError
	}
	if code >= 400 && code < 500 {
		switch code {
		case 401:
			return UnauthorizedError
		case 403:
			return WAFLimitErr
		case 409:
			return CancelReplaceErr
		case 429:
			// Нужно остановить запросы или заблокируют
			// A Retry-After header is sent with a 418 or 429 responses and will give the number of seconds required to wait,
			// in the case of a 429, to prevent a ban, or, in the case of a 418, until the ban is over.
			return RateLimitError
		case 418:
			return IPAutoBannedErr
		default:
			return MalformedRequestErr.WithHttpCode(code)
		}
	}

	return nil
}
