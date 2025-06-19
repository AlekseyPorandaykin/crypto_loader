package request

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Asset struct {
	host *url.URL
}

func NewAsset(host *url.URL) *Asset {
	return &Asset{host: host}
}

func (r *Asset) GetAssetInfo(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-asset-info").String(),
		Method: http.MethodGet,
		Params: []Param{{Key: "accountType", Value: string(accountType)}},
	}, apiKey, apiSecret)
}

func (r *Asset) GetAllCoinsBalance(ctx context.Context, apiKey, apiSecret string, accountType domain.AccountType, coins []string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-account-coins-balance").String(),
		Method: http.MethodGet,
		Params: []Param{
			{Key: "accountType", Value: string(accountType)},
			{Key: "coin", Value: strings.Join(coins, ",")},
		},
	}, apiKey, apiSecret)
}
func (r *Asset) GetCoinExchangeRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/exchange/order-record").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}

func (r *Asset) GetInternalTransferRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-inter-transfer-list").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}

func (r *Asset) GetWithdrawalRecords(ctx context.Context, cred CredentialParam, param AssetWithdrawalRecordsParam) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/withdraw/query-record").String(),
		Method: http.MethodGet,
		Params: param.Params(),
	}, cred.ApiKey, cred.ApiSecret)
}

func (r *Asset) GetDepositRecords(ctx context.Context, cred CredentialParam, param GetDepositRecordParam) (*http.Request, error) {
	var params []Param
	if param.Coin != "" {
		params = append(params, Param{Key: "coin", Value: param.Coin})
	}
	if !param.StartTime.IsZero() {
		params = append(params, Param{Key: "startTime", Value: strconv.Itoa(int(param.StartTime.UnixMilli()))})
	}
	if !param.EndTime.IsZero() {
		params = append(params, Param{Key: "endTime", Value: strconv.Itoa(int(param.EndTime.UnixMilli()))})
	}
	if param.Limit > 0 {
		params = append(params, Param{Key: "limit", Value: strconv.Itoa(param.Limit)})
	}
	if param.Cursor != "" {
		params = append(params, Param{Key: "cursor", Value: param.Cursor})
	}
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/deposit/query-record").String(),
		Method: http.MethodGet,
		Params: params,
	}, cred.ApiKey, cred.ApiSecret)
}
func (r *Asset) GetUniversalTransferRecords(ctx context.Context, apiKey, apiSecret string) (*http.Request, error) {
	return personalRequest(ctx, Request{
		Url:    r.host.JoinPath("/v5/asset/transfer/query-universal-transfer-list").String(),
		Method: http.MethodGet,
	}, apiKey, apiSecret)
}
