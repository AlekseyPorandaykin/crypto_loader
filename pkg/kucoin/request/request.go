package request

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ApiKeyVersionV1 is v1 api key version
const ApiKeyVersionV1 = "1"

// ApiKeyVersionV2 is v2 api key version
const ApiKeyVersionV2 = "2"

type ApiRequest struct {
	Method    string
	Endpoint  string
	Host      *url.URL
	Params    map[string]string
	Timestamp int64
}

func CreateApiRequest(method string, host *url.URL, endpoint string, params map[string]string) *ApiRequest {
	return &ApiRequest{
		Method:    method,
		Endpoint:  endpoint,
		Host:      host,
		Params:    params,
		Timestamp: time.Now().UnixNano() / 1000000,
	}
}

func (ar ApiRequest) Url() string {
	urlReq := ar.Host.JoinPath(ar.Endpoint)
	if ar.Params == nil || len(ar.Params) == 0 {
		return urlReq.String()
	}
	if ar.Method == http.MethodGet || ar.Method == http.MethodDelete {
		q := urlReq.Query()
		for key, value := range ar.Params {
			q.Set(key, value)
		}
		urlReq.RawQuery = q.Encode()
	}

	return urlReq.String()
}

func (ar ApiRequest) Body() ([]byte, error) {
	if ar.Method == http.MethodGet || ar.Method == http.MethodDelete {
		return nil, nil
	}
	data, err := json.Marshal(ar.Params)
	if err != nil {
		return nil, errors.Wrap(err, "error marshal params")
	}
	return data, nil
}

func (ar ApiRequest) TimestampStr() string {
	return strconv.FormatInt(ar.Timestamp, 10)
}

type Credential struct {
	ApiKey     string
	Secret     string
	PassPhrase string
	KeyVersion string
}

func (c Credential) sign(apiReq *ApiRequest) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(apiReq.TimestampStr())
	b.WriteString(apiReq.Method)
	b.WriteString(apiReq.Endpoint)

	body, err := apiReq.Body()
	if err != nil {
		return nil, err
	}
	if body != nil {
		b.Write(body)
	}

	hm := hmac.New(sha256.New, []byte(c.Secret))
	hm.Write(b.Bytes())
	return []byte(base64.StdEncoding.EncodeToString(hm.Sum(nil))), nil
}

func createRequest(ctx context.Context, apiReq *ApiRequest) (*http.Request, error) {
	body, err := apiReq.Body()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(apiReq.Method, apiReq.Url(), bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "error create request")
	}
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func signRequest(apiReq *ApiRequest, c *Credential, req *http.Request) (*http.Request, error) {
	s, err := c.sign(apiReq)
	if err != nil {
		return nil, err
	}
	if c.KeyVersion == "" {
		c.KeyVersion = ApiKeyVersionV2
	}
	req.Header.Set("KC-API-KEY", c.ApiKey)
	req.Header.Set("KC-API-PASSPHRASE", passPhraseEncrypt([]byte(c.Secret), []byte(c.PassPhrase)))
	req.Header.Set("KC-API-TIMESTAMP", apiReq.TimestampStr())
	req.Header.Set("KC-API-SIGN", string(s))
	req.Header.Set("KC-API-KEY-VERSION", c.KeyVersion)
	return req, nil
}

func passPhraseEncrypt(key, plain []byte) string {
	hm := hmac.New(sha256.New, key)
	hm.Write(plain)
	return base64.StdEncoding.EncodeToString(hm.Sum(nil))
}
