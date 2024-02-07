package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v3/request"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v3/response"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v3/sender"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
)

type Client struct {
	sender sender.Sender

	assetReq *request.Asset
}

func NewClient(host string) (*Client, error) {
	urlHost, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Client{sender: sender.NewBasic(), assetReq: request.NewAsset(urlHost)}, nil
}

func (c *Client) AssetWithdrawRecords(ctx context.Context, cred request.CredentialParam, param request.AssetWithdrawParam) (response.AssetWithdrawRecordsResponse, error) {
	req, err := c.assetReq.GetWithdrawRecords(ctx, cred, param)
	if err != nil {
		return response.AssetWithdrawRecordsResponse{}, err
	}
	var dest response.AssetWithdrawRecordsResponse
	if err := c.sendRequest(req, &dest); err != nil {
		return response.AssetWithdrawRecordsResponse{}, err
	}
	return dest, nil
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
