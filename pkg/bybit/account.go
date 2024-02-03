package bybit

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"
	"github.com/pkg/errors"
	"io"
)

func (c *Client) AccountWalletBalance(ctx context.Context, apiKey, apiSecret string, account domain.AccountType) (any, error) {
	req, err := c.accountRequest.GetWalletBalance(ctx, apiKey, apiSecret, account)
	if err != nil {
		return nil, WrapCreateRequestErr(err)
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)

	return data, err
}

func (c *Client) AccountTransactionLog(ctx context.Context, apiKey, apiSecret string) (any, error) {
	req, err := c.accountRequest.GetTransactionLog(ctx, apiKey, apiSecret)
	if err != nil {
		return nil, WrapCreateRequestErr(err)
	}
	res, err := c.sender.Send(req)
	if err != nil {
		return nil, errors.Wrap(err, "http client do")
	}
	if res.Body == nil {
		return nil, errors.New("empty body response")
	}
	defer func() { _ = res.Body.Close() }()
	data, err := io.ReadAll(res.Body)

	return data, err
}
