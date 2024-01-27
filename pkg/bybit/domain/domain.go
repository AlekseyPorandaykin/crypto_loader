package domain

//@doc https://bybit-exchange.github.io/docs/v5/enum

// AccountType @doc https://bybit-exchange.github.io/docs/v5/enum#accounttype
type AccountType string

const (
	ContractAccountType AccountType = "CONTRACT" //Inverse Derivatives Account
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
	LinearOrderCategory  OrderCategory = "linear"
	InverseOrderCategory OrderCategory = "inverse"
	OptionOrderCategory  OrderCategory = "option"
)
