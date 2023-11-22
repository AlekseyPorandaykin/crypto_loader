package requests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const XFormContentType = "application/x-www-form-urlencoded"

const DefaultReceiveWindow = 5000

func longTimestamp(date time.Time) string {
	return strconv.FormatInt(date.UTC().UnixMilli(), 10)
}
func signature(params string, secretKey string) string {
	sig := hmac.New(sha256.New, []byte(secretKey))
	sig.Write([]byte(params))
	return hex.EncodeToString(sig.Sum(nil))
}

type param struct {
	key   string
	value string
}

func generateQuery(params []param) string {
	values := url.Values{}
	for _, p := range params {
		values.Add(p.key, p.value)
	}

	return values.Encode()
}

func createSecurityRequest(apiKey, secretKey, method, url string, params []param) (*http.Request, error) {
	var (
		body  io.Reader
		query string
	)
	query = generateQuery(params)
	if query != "" {
		query += "&"
	}
	query += fmt.Sprintf("recvWindow=%d", DefaultReceiveWindow)
	query += "&timestamp=" + longTimestamp(time.Now())
	query += "&signature=" + signature(query, secretKey)

	if method == http.MethodGet {
		url += "?" + query
	} else {
		body = strings.NewReader(query)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-MBX-APIKEY", apiKey)
	if method == http.MethodPost {
		req.Header.Set("Content-Type", XFormContentType)
	}

	return req, nil
}
