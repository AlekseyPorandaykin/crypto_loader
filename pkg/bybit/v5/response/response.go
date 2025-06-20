package response

import (
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/domain"
	"strings"
	"time"
)

type CheckerResponse interface {
	IsOk() bool
	ErrMessage() string
	StatusCode() int
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

type OpenOrder struct {
	OrderLinkID        string           `json:"orderLinkId"`
	OrderID            string           `json:"orderId"`
	BlockTradeID       string           `json:"blockTradeId"`
	Symbol             string           `json:"symbol"`
	Price              string           `json:"price"`
	PositionIdx        PositionIdx      `json:"positionIdx"`
	Qty                string           `json:"qty"`
	Side               Side             `json:"side"`
	IsLeverage         string           `json:"isLeverage"`
	OrderStatus        OrderStatus      `json:"orderStatus"`
	CreateType         CreateType       `json:"createType"`
	CancelType         CancelType       `json:"cancelType"`
	RejectReason       RejectReason     `json:"rejectReason"`
	AvgPrice           string           `json:"avgPrice"`
	LeavesQty          string           `json:"leavesQty"`
	LeavesValue        string           `json:"leavesValue"`
	CumExecQty         string           `json:"cumExecQty"`
	CumExecValue       string           `json:"cumExecValue"`
	CumExecFee         string           `json:"cumExecFee"`
	TimeInForce        TimeInForce      `json:"timeInForce"`
	OrderType          OrderType        `json:"orderType"` //	Order type. Market,Limit. For TP/SL order, it means the order type after triggered
	StopOrderType      StopOrderType    `json:"stopOrderType"`
	OrderIv            string           `json:"orderIv"`
	MarketUnit         string           `json:"marketUnit"`
	TriggerPrice       string           `json:"triggerPrice"`
	TakeProfit         string           `json:"takeProfit"`
	TpslMode           TpslMode         `json:"tpslMode"` //TP/SL mode, Spot does not have this field, and Option returns always ""
	OcoTriggerType     string           `json:"ocoTriggerType"`
	TpLimitPrice       string           `json:"tpLimitPrice"`
	SlLimitPrice       string           `json:"slLimitPrice"`
	StopLoss           string           `json:"stopLoss"`
	TpTriggerBy        string           `json:"tpTriggerBy"`
	SlTriggerBy        string           `json:"slTriggerBy"`
	TriggerDirection   TriggerDirection `json:"triggerDirection"`
	TriggerBy          TriggerBy        `json:"triggerBy"`
	LastPriceOnCreated string           `json:"lastPriceOnCreated"` //Last price when place the order
	ReduceOnly         bool             `json:"reduceOnly"`
	CloseOnTrigger     bool             `json:"closeOnTrigger"` // Close On Trigger function seeks to execute as a closing order that guarantees its execution regardless of margin requirements.
	PlaceType          string           `json:"placeType"`
	CreatedTime        string           `json:"createdTime"` //Order created timestamp (ms)
	UpdatedTime        string           `json:"updatedTime"` //Order updated timestamp (ms)
	SmpType            SmpType          `json:"smpType"`
	SmpGroup           int              `json:"smpGroup"`
	SmpOrderId         string           `json:"smpOrderId"`
}

type TradeOpenOrders struct {
	NextPageCursor string      `json:"nextPageCursor"`
	Category       string      `json:"category"`
	List           []OpenOrder `json:"list"`
}

type TradeOpenOrdersResponse struct {
	CommonResponse
	Result TradeOpenOrders `json:"result"`
}

type InternalTransferRecord struct {
	TransferID      string `json:"transferId"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
}

type InternalTransferRecords struct {
	List           []InternalTransferRecord `json:"list"`
	NextPageCursor string                   `json:"nextPageCursor"`
}

type InternalTransferRecordsResponse struct {
	CommonResponse
	Result InternalTransferRecords `json:"result"`
}

type WithdrawalRecord struct {
	WithdrawId   string `json:"withdrawId"`
	TxID         string `json:"txID"`
	WithdrawType int    `json:"withdrawType"`
	Coin         string `json:"coin"`
	Chain        string `json:"chain"`
	Amount       string `json:"amount"`
	WithdrawFee  string `json:"withdrawFee"`
	Status       string `json:"status"`
	ToAddress    string `json:"toAddress"`
	Tag          string `json:"tag"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type WithdrawalRecords struct {
	Rows           []WithdrawalRecord `json:"rows"`
	NextPageCursor string             `json:"nextPageCursor"`
}

type WithdrawalRecordsResponse struct {
	CommonResponse
	Result WithdrawalRecords `json:"result"`
}

type TransactionLog struct {
	ID              string `json:"id"`
	Symbol          string `json:"symbol"`
	Category        string `json:"category"`
	Side            string `json:"side"`
	TransactionTime string `json:"transactionTime"`
	Qty             string `json:"qty"`
	Size            string `json:"size"`
	Currency        string `json:"currency"`
	TradePrice      string `json:"tradePrice"`
	Funding         string `json:"funding"`
	Fee             string `json:"fee"`
	CashFlow        string `json:"cashFlow"`
	Change          string `json:"change"`
	CashBalance     string `json:"cashBalance"`
	FeeRate         string `json:"feeRate"`
	BonusChange     string `json:"bonusChange"`
	TradeID         string `json:"tradeId"`
	OrderID         string `json:"orderId"`
	OrderLinkID     string `json:"orderLinkId"`
	Type            string `json:"type"`
}

type TransactionLogs struct {
	List           []TransactionLog `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type TransactionLogsResponse struct {
	CommonResponse
	Result TransactionLogs `json:"result"`
}

type InstrumentInfo struct {
	Symbol             string      `json:"symbol"`          // Название торговой пары
	ContractType       string      `json:"contractType"`    // Тип контракта (например, LinearPerpetual для USDT бессрочных, InversePerpetual для инверсных бессрочных, InverseFutures для инверсных фьючерсов).
	Status             string      `json:"status"`          // Статус инструмента (например, Trading, PreLaunch, Settling, Delivering, Closed).
	BaseCoin           string      `json:"baseCoin"`        // Базовая валюта (например, BTC).
	QuoteCoin          string      `json:"quoteCoin"`       // Валюта котировки (например, USDT).
	LaunchTime         string      `json:"launchTime"`      // Время запуска инструмента в Unix timestamp (мс).
	DeliveryTime       string      `json:"deliveryTime"`    // Время доставки (для фьючерсов) или время делистинга (для бессрочных) в Unix timestamp (мс). 0 для бессрочных, если нет планов на делистинг.
	DeliveryFeeRate    string      `json:"deliveryFeeRate"` // Ставка комиссии за доставку (для фьючерсов).
	PriceScale         string      `json:"priceScale"`      // Масштаб цены, количество знаков после запятой для цены.
	UnifiedMarginTrade bool        `json:"unifiedMarginTrade"`
	FundingInterval    int         `json:"fundingInterval"`
	SettleCoin         string      `json:"settleCoin"`
	CopyTrading        string      `json:"copyTrading"`
	UpperFundingRate   string      `json:"upperFundingRate"`
	LowerFundingRate   string      `json:"lowerFundingRate"`
	IsPreListing       bool        `json:"isPreListing"`
	PreListingInfo     interface{} `json:"preListingInfo"`
	DisplayName        string      `json:"displayName"`
	LeverageFilter     struct {
		MinLeverage  string `json:"minLeverage"`  // Минимальное кредитное плечо.
		MaxLeverage  string `json:"maxLeverage"`  // Максимальное кредитное плечо.
		LeverageStep string `json:"leverageStep"` // Шаг изменения кредитного плеча.
	} `json:"leverageFilter"` // Объект, содержащий информацию о кредитном плече.
	PriceFilter struct {
		MinPrice string `json:"minPrice"` // Минимальная цена ордера.
		MaxPrice string `json:"maxPrice"` // Максимальная цена ордера.
		TickSize string `json:"tickSize"` // Минимальный шаг цены.
	} `json:"priceFilter"` // Объект, содержащий информацию о фильтре цены.
	LotSizeFilter struct {
		MaxOrderQty         string `json:"maxOrderQty"`         // Максимальное количество в одном ордере.
		MinOrderQty         string `json:"minOrderQty"`         // Минимальное количество в одном ордере.
		QtyStep             string `json:"qtyStep"`             // Шаг изменения количества.
		PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"` // Максимальное количество для Post-Only ордера.
		MaxMktOrderQty      string `json:"maxMktOrderQty"`
		MinNotionalValue    string `json:"minNotionalValue"`
	} `json:"lotSizeFilter"` // Объект, содержащий информацию о фильтре размера лота.
	RiskParameters struct {
		PriceLimitRatioX string `json:"priceLimitRatioX"`
		PriceLimitRatioY string `json:"priceLimitRatioY"`
	} `json:"riskParameters"`
}

type InstrumentsInfo struct {
	Category string           `json:"category"`
	List     []InstrumentInfo `json:"list"`
}

type InstrumentsInfoResponse struct {
	CommonResponse
	Result InstrumentsInfo `json:"result"`
}

type ClosedPnl struct {
	Symbol        string `json:"symbol"`
	OrderId       string `json:"orderId"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	OrderPrice    string `json:"orderPrice"`
	OrderType     string `json:"orderType"`
	ExecType      string `json:"execType"` //Exec type. Trade, BustTrade, SessionSettlePnL, Settle, MovePosition
	ClosedSize    string `json:"closedSize"`
	CumEntryValue string `json:"cumEntryValue"` //Cumulated Position value
	AvgEntryPrice string `json:"avgEntryPrice"`
	CumExitValue  string `json:"cumExitValue"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ClosedPnl     string `json:"closedPnl"`
	FillCount     string `json:"fillCount"`
	Leverage      string `json:"leverage"`
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}

type ListClosedPnl struct {
	Category       string      `json:"category"`
	NextPageCursor string      `json:"nextPageCursor"`
	List           []ClosedPnl `json:"list"`
}

type ClosedPnlResponse struct {
	CommonResponse
	Result ListClosedPnl `json:"result"`
}

type PositionInfo struct {
	Symbol                 string `json:"symbol"`
	Leverage               string `json:"leverage"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	AvgPrice               string `json:"avgPrice"`
	LiqPrice               string `json:"liqPrice"`
	RiskLimitValue         string `json:"riskLimitValue"`
	TakeProfit             string `json:"takeProfit"`
	PositionValue          string `json:"positionValue"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	TpslMode               string `json:"tpslMode"`
	RiskId                 int    `json:"riskId"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	MarkPrice              string `json:"markPrice"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	PositionMM             string `json:"positionMM"`
	CreatedTime            string `json:"createdTime"`
	PositionIdx            int    `json:"positionIdx"`
	PositionIM             string `json:"positionIM"`
	Seq                    int64  `json:"seq"`
	UpdatedTime            string `json:"updatedTime"`
	Side                   string `json:"side"`
	BustPrice              string `json:"bustPrice"`
	PositionBalance        string `json:"positionBalance"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	CurRealisedPnl         string `json:"curRealisedPnl"`
	Size                   string `json:"size"`
	PositionStatus         string `json:"positionStatus"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	StopLoss               string `json:"stopLoss"`
	TradeMode              int    `json:"tradeMode"`
	SessionAvgPrice        string `json:"sessionAvgPrice"`
}

type ListPositionInfo struct {
	Category       string         `json:"category"`
	NextPageCursor string         `json:"nextPageCursor"`
	List           []PositionInfo `json:"list"`
}

type PositionInfoResponse struct {
	CommonResponse
	Result ListPositionInfo `json:"result"`
}

type AccountWalletBalance struct {
	TotalEquity            string                     `json:"totalEquity"`
	AccountIMRate          string                     `json:"accountIMRate"`
	TotalMarginBalance     string                     `json:"totalMarginBalance"`
	TotalInitialMargin     string                     `json:"totalInitialMargin"`
	AccountType            string                     `json:"accountType"`
	TotalAvailableBalance  string                     `json:"totalAvailableBalance"`
	AccountMMRate          string                     `json:"accountMMRate"`
	TotalPerpUPL           string                     `json:"totalPerpUPL"`
	TotalWalletBalance     string                     `json:"totalWalletBalance"`
	AccountLTV             string                     `json:"accountLTV"`
	TotalMaintenanceMargin string                     `json:"totalMaintenanceMargin"`
	Coin                   []CoinAccountWalletBalance `json:"coin"`
}

type CoinAccountWalletBalance struct {
	AvailableToBorrow   string `json:"availableToBorrow"`
	Bonus               string `json:"bonus"`
	AccruedInterest     string `json:"accruedInterest"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	TotalOrderIM        string `json:"totalOrderIM"`
	Equity              string `json:"equity"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UsdValue            string `json:"usdValue"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	SpotHedgingQty      string `json:"spotHedgingQty"`
	BorrowAmount        string `json:"borrowAmount"`
	TotalPositionIM     string `json:"totalPositionIM"`
	WalletBalance       string `json:"walletBalance"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Locked              string `json:"locked"`
	MarginCollateral    bool   `json:"marginCollateral"`
	Coin                string `json:"coin"`
}

type AccountWalletBalanceResponse struct {
	CommonResponse
	Result struct {
		List []AccountWalletBalance `json:"list"`
	} `json:"result"`
}

type AccountFeeRate struct {
	Symbol       string `json:"symbol"`
	TakerFeeRate string `json:"takerFeeRate"`
	MakerFeeRate string `json:"makerFeeRate"`
}

type AccountFeeRateResponse struct {
	CommonResponse
	Result struct {
		List []AccountFeeRate `json:"list"`
	} `json:"result"`
}

type Candlestick struct {
	StartTime  string // Start time of the candle (ms)
	OpenPrice  string
	HighPrice  string
	LowPrice   string
	ClosePrice string // Close price. Is the last traded price when the candle is not closed
	Volume     string // Trade volume. Unit of contract: pieces of contract. Unit of spot: quantity of coins
	Turnover   string // Turnover. Unit of figure: quantity of quota coin
}

type ListKline struct {
	Symbol   string     `json:"symbol"`
	Category string     `json:"category"`
	List     [][]string `json:"list"`
}

func (lk ListKline) Candlesticks() []Candlestick {
	res := make([]Candlestick, 0, len(lk.List))
	for _, item := range lk.List {
		if len(item) < 7 {
			continue
		}
		res = append(res, Candlestick{
			StartTime:  item[0],
			OpenPrice:  item[1],
			HighPrice:  item[2],
			LowPrice:   item[3],
			ClosePrice: item[4],
			Volume:     item[5],
			Turnover:   item[6],
		})
	}
	return res
}

type GetKlineResponse struct {
	CommonResponse
	Result ListKline `json:"result"`
}

type Bid struct {
	Price string
	Size  string
}

type Ask struct {
	Price string
	Size  string
}

type MarketOrderBook struct {
	Symbol        string     `json:"s"`
	ListBid       [][]string `json:"b"`   //Bid, buyer. Sort by price desc, b[0] - Bid price, b[1] - Bid size
	ListAsk       [][]string `json:"a"`   //Ask, seller. Order by price asc, a[0] - Ask price, a[1] - Ask size
	Timestamp     int        `json:"ts"`  //The timestamp (ms) that the system generates the data
	UpdateID      int        `json:"u"`   //Update ID, is always in sequence
	CrossSequence int        `json:"seq"` //Cross sequence
}

func (o MarketOrderBook) Bids() []Bid {
	bids := make([]Bid, 0, len(o.ListBid))
	for i := 0; i < len(o.ListBid); i++ {
		bids = append(bids, Bid{
			Price: o.ListBid[i][0],
			Size:  o.ListBid[i][1],
		})
	}
	return bids
}
func (o MarketOrderBook) Asks() []Ask {
	asks := make([]Ask, 0, len(o.ListAsk))
	for i := 0; i < len(o.ListAsk); i++ {
		asks = append(asks, Ask{
			Price: o.ListAsk[i][0],
			Size:  o.ListAsk[i][1],
		})
	}
	return asks
}

type GetOrderBookResponse struct {
	CommonResponse
	Result MarketOrderBook `json:"result"`
}
type GetApiKeyInformation struct {
	Id          string `json:"id"`
	Note        string `json:"note"`
	ApiKey      string `json:"apiKey"`
	ReadOnly    int    `json:"readOnly"`
	Secret      string `json:"secret"`
	Permissions struct {
		ContractTrade []string      `json:"ContractTrade"`
		Spot          []string      `json:"Spot"`
		Wallet        []string      `json:"Wallet"`
		Options       []string      `json:"Options"`
		Derivatives   []string      `json:"Derivatives"`
		CopyTrading   []interface{} `json:"CopyTrading"`
		BlockTrade    []interface{} `json:"BlockTrade"`
		Exchange      []string      `json:"Exchange"`
		NFT           []interface{} `json:"NFT"`
		Affiliate     []interface{} `json:"Affiliate"`
	} `json:"permissions"`
	Ips           []string  `json:"ips"`
	Type          int       `json:"type"`
	DeadlineDay   int       `json:"deadlineDay"`
	ExpiredAt     time.Time `json:"expiredAt"`
	CreatedAt     time.Time `json:"createdAt"`
	Unified       int       `json:"unified"`
	Uta           int       `json:"uta"`
	UserID        int       `json:"userID"`
	InviterID     int       `json:"inviterID"`
	VipLevel      string    `json:"vipLevel"`
	MktMakerLevel string    `json:"mktMakerLevel"`
	AffiliateID   int       `json:"affiliateID"`
	RsaPublicKey  string    `json:"rsaPublicKey"`
	IsMaster      bool      `json:"isMaster"`
	ParentUid     string    `json:"parentUid"`
	KycLevel      string    `json:"kycLevel"`
	KycRegion     string    `json:"kycRegion"`
}

type GetApiKeyInformationResponse struct {
	CommonResponse
	Result GetApiKeyInformation `json:"result"`
}

type TradeAmendOrder struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type TradeAmendOrderResponse struct {
	CommonResponse
	Result TradeAmendOrder `json:"result"`
}

type CoinInfoRowChain struct {
	ChainType             string `json:"chainType"`             // Тип цепочки
	Confirmation          string `json:"confirmation"`          // Количество подтверждений для депозита: как только это количество будет достигнуто, ваши средства будут зачислены на счет и доступны для торговли.
	WithdrawFee           string `json:"withdrawFee"`           // Комиссия за вывод средств. Если пусто, это означает, что монета не поддерживает вывод.
	DepositMin            string `json:"depositMin"`            // Минимальная сумма депозита.
	WithdrawMin           string `json:"withdrawMin"`           // Минимальная сумма вывода.
	Chain                 string `json:"chain"`                 // Название цепочки/сети
	ChainDeposit          string `json:"chainDeposit"`          // Статус депозита для данной цепочки: 0 (приостановлен), 1 (нормальный).
	ChainWithdraw         string `json:"chainWithdraw"`         // Статус вывода для данной цепочки: 0 (приостановлен), 1 (нормальный).
	MinAccuracy           string `json:"minAccuracy"`           // Точность (количество знаков после запятой) для вывода или депозита.
	WithdrawPercentageFee string `json:"withdrawPercentageFee"` // Процентная комиссия за вывод. Это реальное число, например, 0.022 означает 2.2%.
	ContractAddress       string `json:"contractAddress"`       // Адрес контракта токена на данной цепочке. "" означает отсутствие адреса контракта (для нативных монет).
	SafeConfirmNumber     string `json:"safeConfirmNumber"`     // Количество подтверждений безопасности: как только это число будет достигнуто, ваши средства в эквиваленте USD будут полностью разблокированы и доступны для вывода.
}
type CoinInfoRow struct {
	Name         string             `json:"name"`         // Полное имя монеты
	Coin         string             `json:"coin"`         // Символ монеты
	RemainAmount string             `json:"remainAmount"` // Максимальная сумма вывода за одну транзакцию.
	Chains       []CoinInfoRowChain `json:"chains"`       // Массив объектов, содержащих информацию о различных цепочках (сетях), поддерживаемых для данной монеты.
}

type CoinInfoResponse struct {
	CommonResponse
	Result struct {
		Rows []CoinInfoRow `json:"rows"`
	} `json:"result"`
}
