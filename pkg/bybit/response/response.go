package response

import "github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/domain"

type CommonResponse struct {
	//https://bybit-exchange.github.io/docs/v5/error#uma--uta--futures-of-classic-account
	Code       int         `json:"retCode"`    //Success/Error code
	Message    string      `json:"retMsg"`     //Success/Error msg. OK, success, SUCCESS indicate a successful response
	Result     interface{} `json:"result"`     //Business data result
	ExtendInfo interface{} `json:"retExtInfo"` //Extend info. Most of the time, it is {}
	Time       int         `json:"time"`       //Current timestamp (ms)
}

type AccountAssets struct {
	Coin     string `json:"coin"`
	Frozen   string `json:"frozen"`
	Free     string `json:"free"`
	Withdraw string `json:"withdraw"`
}

type WalletTypeResult struct {
	UID         string               `json:"uid"`
	AccountType []domain.AccountType `json:"accountType"`
}

type WalletTypeResponse struct {
	CommonResponse
	Result struct {
		Accounts []WalletTypeResult `json:"accounts"`
	} `json:"result"`
}

type AccountTypeAssets struct {
	Status domain.AccountStatus `json:"status"`
	Assets []AccountAssets      `json:"assets"`
}

type AssetResponse struct {
	CommonResponse
	Result map[string]AccountTypeAssets `json:"result"`
}

type CoinBalance struct {
	Coin            string `json:"coin"`
	TransferBalance string `json:"transferBalance"`
	WalletBalance   string `json:"walletBalance"`
	Bonus           string `json:"bonus"`
}

type AccountCoinBalance struct {
	MemberID    string             `json:"memberId"`
	AccountType domain.AccountType `json:"accountType"`
	Balance     []CoinBalance      `json:"balance"`
}

type CoinBalanceResponse struct {
	CommonResponse
	Result AccountCoinBalance `json:"result"`
}

type TradeOrder struct {
	OrderLinkID        string `json:"orderLinkId"` //User customised order ID
	OrderID            string `json:"orderId"`
	BlockTradeID       string `json:"blockTradeId"`
	Symbol             string `json:"symbol"`
	Price              string `json:"price"`
	IsLeverage         string `json:"isLeverage"`  //Whether to borrow. Unified spot only. 0: false, 1: true. . Classic spot is not supported, always 0
	PositionIDX        int    `json:"positionIdx"` //Position index. Used to identify positions in different position modes
	Qty                string `json:"qty"`
	Side               string `json:"side"` //	Side. Buy,Sell
	OrderStatus        string `json:"orderStatus"`
	CreateType         string `json:"createType"` //Order create type .Only for category=linear or inverse .Spot, Option do not have this key
	CancelType         string `json:"cancelType"`
	RejectReason       string `json:"rejectReason"` //Reject reason. Classic spot is not supported
	AvgPrice           string `json:"avgPrice"`     //Average filled price. UTA: returns "" for those orders without avg price. Classic account: returns "0" for those orders without avg price
	LeavesQty          string `json:"leavesQty"`    //	The estimated value not executed. Classic spot is not supported
	LeavesValue        string `json:"leavesValue"`
	CumExecQty         string `json:"cumExecQty"`   //Cumulative executed order qty
	CumExecValue       string `json:"cumExecValue"` //Cumulative executed order value. Classic spot is not supported
	CumExecFee         string `json:"cumExecFee"`   //Cumulative executed trading fee. Classic spot is not supported
	TimeInForce        string `json:"timeInForce"`
	OrderType          string `json:"orderType"` //Order type. Market,Limit. For TP/SL order, it means the order type after triggered. Block trade Roll Back, Block trade-Limit: Unique enum values for Unified account block trades
	StopOrderType      string `json:"stopOrderType"`
	OrderIv            string `json:"orderIv"`
	TriggerPrice       string `json:"triggerPrice"` //Trigger price. If stopOrderType=TrailingStop, it is activate price. Otherwise, it is trigger price
	TakeProfit         string `json:"takeProfit"`
	StopLoss           string `json:"stopLoss"`
	TpTriggerBy        string `json:"tpTriggerBy"`
	SlTriggerBy        string `json:"slTriggerBy"`
	TriggerDirection   int    `json:"triggerDirection"`
	TriggerBy          string `json:"triggerBy"`
	LastPriceOnCreated string `json:"lastPriceOnCreated"`
	ReduceOnly         bool   `json:"reduceOnly"`
	CloseOnTrigger     bool   `json:"closeOnTrigger"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
	SmpType            string `json:"smpType"`
	SmpGroup           int    `json:"smpGroup"`
	SmpOrderId         string `json:"smpOrderId"`
}

type TradeOrderHistory struct {
	NextPageCursor string       `json:"nextPageCursor"`
	Category       string       `json:"category"`
	List           []TradeOrder `json:"list"`
}

type TradeOrderHistoryResponse struct {
	CommonResponse
	Result TradeOrderHistory `json:"result"`
}

type Trade struct {
	Symbol          string `json:"symbol"`
	OrderID         string `json:"orderId"`
	OrderLinkID     string `json:"orderLinkId"` //User customized order ID. Classic spot is not supported
	Side            string `json:"side"`        //Side. Buy,Sell
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	OrderType       string `json:"orderType"`     //Order type. Market,Limit
	CreateType      string `json:"createType"`    //Order create type .Only for category=linear or inverse .Spot, Option do not have this key
	StopOrderType   string `json:"stopOrderType"` //Stop order type. If the order is not stop order, it either returns UNKNOWN or "". Classic spot is not supported
	ExecFee         string `json:"execFee"`       //Executed trading fee. You can get spot fee currency instruction (https://bybit-exchange.github.io/docs/v5/enum#spot-fee-currency-instruction)
	ExecID          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	FeeCurrency     string `json:"feeCurrency"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`         //Implied volatility. Valid for option
	MarkIv          string `json:"markIv"`          //Implied volatility of mark price. Valid for option
	MarkPrice       string `json:"markPrice"`       //The mark price of the symbol when executing. Classic spot is not supported
	IndexPrice      string `json:"indexPrice"`      //The index price of the symbol when executing. Valid for option only
	UnderlyingPrice string `json:"underlyingPrice"` //The underlying price of the symbol when executing. Valid for option
	BlockTradeID    string `json:"blockTradeId"`    //Paradigm block trade ID
	ClosedSize      string `json:"closedSize"`
	//Seq             string `json:"seq"`
}

type TradeHistory struct {
	NextPageCursor string  `json:"nextPageCursor"`
	Category       string  `json:"category"`
	List           []Trade `json:"list"`
}

type TradeHistoryResponse struct {
	CommonResponse
	Result TradeHistory `json:"result"`
}
