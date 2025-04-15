package response

import (
	"strings"
)

type CheckerResponse interface {
	IsOk() bool
	ErrMessage() string
}

type CommonResponse struct {
	//https://bybit-exchange.github.io/docs/v5/error#uma--uta--futures-of-classic-account
	Code       int         `json:"retCode"`    //Success/Error code
	Message    string      `json:"retMsg"`     //Success/Error msg. OK, success, SUCCESS indicate a successful response
	Result     interface{} `json:"result"`     //Business data result
	ExtendInfo interface{} `json:"retExtInfo"` //Extend info. Most of the time, it is {}
	Time       int         `json:"time"`       //Current timestamp (ms)
}

func (resp CommonResponse) IsOk() bool {
	return resp.Code == 0
}
func (resp CommonResponse) ErrMessage() string {
	if strings.ToLower(resp.Message) == "success" || strings.ToLower(resp.Message) == "ok" {
		return ""
	}
	return resp.Message
}
func (resp CommonResponse) StatusCode() int {
	return resp.Code
}

type AssetWithdrawRecord struct {
	Coin         string `json:"coin"`
	Chan         string `json:"chan"`
	Amount       string `json:"amount"`
	TxID         string `json:"txID"`
	Status       string `json:"status"`
	ToAddress    string `json:"toAddress"`
	Tag          string `json:"tag"`
	WithdrawFee  string `json:"withdrawFee"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
	WithdrawID   string `json:"withdrawId"`
	WithdrawType int    `json:"withdrawType"`
}

type AssetWithdrawRecords struct {
	Rows           []AssetWithdrawRecord `json:"rows"`
	NextPageCursor string                `json:"nextPageCursor"`
}

type AssetWithdrawRecordsResponse struct {
	CommonResponse
	Result AssetWithdrawRecords `json:"result"`
}
