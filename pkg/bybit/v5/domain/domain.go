package domain

//@doc https://bybit-exchange.github.io/docs/v5/enum

// AccountType @doc https://bybit-exchange.github.io/docs/v5/enum#accounttype
type AccountType string

const (
	ContractAccountType AccountType = "CONTRACT" //Derivatives Account
	UnifiedAccountType  AccountType = "UNIFIED"  //Unified Trading Account
	FundAccountType     AccountType = "FUND"     //Funding Account
	SpotAccountType     AccountType = "SPOT"     //Spot Account
	OptionAccountType   AccountType = "OPTION"   //USDC Derivatives
)

type AccountStatus string

const (
	AccountStatusNormal      = "ACCOUNT_STATUS_NORMAL"
	AccountStatusUnspecified = "ACCOUNT_STATUS_UNSPECIFIED"
)

type OrderCategory string

const (
	SpotOrderCategory    OrderCategory = "spot"
	LinearOrderCategory  OrderCategory = "linear"  //USDT perpetual, and USDC contract, including USDC perp, USDC futures
	InverseOrderCategory OrderCategory = "inverse" //Inverse contract, including Inverse perp, Inverse futures
	OptionOrderCategory  OrderCategory = "option"
)

type ExecType string

const (
	TradeExecType        ExecType = "Trade"
	AdlTradeExecType     ExecType = "AdlTrade" //Auto-Deleveraging
	FundingExecType      ExecType = "Funding"
	BustTradeExecType    ExecType = "BustTrade" //Liquidation
	DeliveryExecType     ExecType = "Delivery"  //USDC futures delivery
	SettleExecType       ExecType = "Settle"    //Inverse futures settlement
	BlockTradeExecType   ExecType = "BlockTrade"
	MovePositionExecType ExecType = "MovePosition"
)

type Side string

const (
	BuySide  Side = "Buy"
	SellSide Side = "Sell"
)

type OrderType string

const (
	MarketOrderType OrderType = "Market"
	LimitOrderType  OrderType = "Limit"
)
