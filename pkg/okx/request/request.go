package request

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ApiRequest struct {
	Method    string
	Endpoint  string
	Host      *url.URL
	Params    map[string]string
	Timestamp string
}

type Credential struct {
	ApiKey     string
	Secret     string
	PassPhrase string
	KeyVersion string
}

func (c Credential) sign(apiReq ApiRequest) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(apiReq.Timestamp)
	b.WriteString(strings.ToUpper(apiReq.Method))
	b.WriteString(apiReq.Endpoint)

	hm := hmac.New(sha256.New, []byte(c.Secret))
	hm.Write(b.Bytes())
	return []byte(base64.StdEncoding.EncodeToString(hm.Sum(nil))), nil
}

func createRequest(ctx context.Context, apiReq ApiRequest) (*http.Request, error) {
	req, err := http.NewRequest(apiReq.Method, apiReq.Host.JoinPath(apiReq.Endpoint).String(), nil)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func signRequest(c Credential, apiReq ApiRequest, req *http.Request) (*http.Request, error) {
	if apiReq.Timestamp == "" {
		apiReq.Timestamp = createTimestamp()
	}
	sign, err := c.sign(apiReq)
	if err != nil {
		return nil, errors.Wrap(err, "error sign request")
	}
	req.Header.Set("OK-ACCESS-KEY", c.ApiKey)
	req.Header.Set("OK-ACCESS-SIGN", string(sign))
	req.Header.Set("OK-ACCESS-TIMESTAMP", createTimestamp())
	req.Header.Set("OK-ACCESS-PASSPHRASE", c.PassPhrase)
	return req, nil
}

func createTimestamp() string {
	now := time.Now().In(time.UTC)
	now.Minute()
	return fmt.Sprintf(
		"%s.000Z",
		now.Format("2006-01-02T15:04:05"),
	)
}
