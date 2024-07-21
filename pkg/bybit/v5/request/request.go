package request

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const defaultRecvWindow = "5000"

type Params []Param

func (p Params) ToMap() map[string]string {
	m := make(map[string]string)
	for _, item := range p {
		m[item.Key] = item.Value
	}
	return m
}

type Param struct {
	Key   string
	Value string
}

type Request struct {
	Url        string
	Method     string
	Params     Params
	RecvWindow int
}

func createRequest(ctx context.Context, req Request) (*http.Request, error) {
	url := req.Url
	var (
		body io.Reader
	)
	switch req.Method {
	case http.MethodGet:
		if len(req.Params) > 0 {
			url = fmt.Sprintf("%s?%s", req.Url, getParams(req.Params))
		}
	case http.MethodPost:
		url = req.Url
	}
	r, err := http.NewRequest(req.Method, url, body)
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)
	return r, nil
}

func personalRequest(ctx context.Context, req Request, apiKey, apiSecret string) (*http.Request, error) {
	t := timestamp()
	r, err := createRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	var (
		recvWindow = defaultRecvWindow
	)
	if req.RecvWindow > 0 {
		recvWindow = strconv.Itoa(req.RecvWindow)
	}
	hmac256 := hmac.New(sha256.New, []byte(apiSecret))
	hmac256.Write([]byte(strconv.FormatInt(t, 10) + apiKey + recvWindow + getParams(req.Params)))
	signature := hex.EncodeToString(hmac256.Sum(nil))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-BAPI-API-KEY", apiKey)
	r.Header.Set("X-BAPI-SIGN", signature)
	r.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(t, 10))
	r.Header.Set("X-BAPI-SIGN-TYPE", "2")
	r.Header.Set("X-BAPI-RECV-WINDOW", recvWindow)

	return r, nil
}

func postPersonalRequest(ctx context.Context, req Request, apiKey, apiSecret string) (*http.Request, error) {
	t := timestamp()
	body, err := json.Marshal(req.Params.ToMap())
	if err != nil {
		return nil, err
	}
	hmac256 := hmac.New(sha256.New, []byte(apiSecret))
	hmac256.Write([]byte(strconv.FormatInt(t, 10) + apiKey + string(body)))
	signature := hex.EncodeToString(hmac256.Sum(nil))

	r, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("X-BAPI-API-KEY", apiKey)
	r.Header.Set("X-BAPI-SIGN", signature)
	r.Header.Set("X-BAPI-TIMESTAMP", strconv.FormatInt(t, 10))

	return r, nil
}

func getParams(data []Param) string {
	params := make([]string, 0, len(data))
	for _, item := range data {
		params = append(params, fmt.Sprintf("%s=%s", item.Key, item.Value))
	}
	return strings.Join(params, "&")
}

func timestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
