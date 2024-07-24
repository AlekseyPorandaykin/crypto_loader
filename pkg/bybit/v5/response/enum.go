package response

type OrderType string

const (
	MarketOrderType  OrderType = "Market"
	LimitOrderType   OrderType = "Limit"
	UnknownOrderType OrderType = "UNKNOWN" //is not a valid request parameter value. Is only used in some responses. Mainly, it is used when execType is Funding.
)

type CancelType string

const (
	CancelByUserCancelType        CancelType = "CancelByUser"
	CancelByReduceOnlyCancelType  CancelType = "CancelByReduceOnly"
	CancelByPrepareLiqCancelType  CancelType = "CancelByPrepareLiq"
	CancelAllBeforeLiqCancelType  CancelType = "CancelAllBeforeLiq"
	CancelByPrepareAdlCancelType  CancelType = "CancelByPrepareAdl"
	CancelAllBeforeAdlCancelType  CancelType = "CancelAllBeforeAdl"
	CancelByAdminCancelType       CancelType = "CancelByAdmin"
	CancelBySettleCancelType      CancelType = "CancelBySettle"
	CancelByTpSlTsClearCancelType CancelType = "CancelByTpSlTsClear"
	CancelBySmpCancelType         CancelType = "CancelBySmp"
)

type StopOrderType string

const (
	TakeProfitStopOrderType             StopOrderType = "TakeProfit"
	StopLossStopOrderType               StopOrderType = "StopLoss"
	TrailingStopStopOrderType           StopOrderType = "TrailingStop"
	StopStopOrderType                   StopOrderType = "Stop"
	PartialTakeProfitStopOrderType      StopOrderType = "PartialTakeProfit"
	PartialStopLossStopOrderType        StopOrderType = "PartialStopLoss"
	TpslOrderStopOrderType              StopOrderType = "tpslOrder"              //spot TP/SL order
	OcoOrderStopOrderType               StopOrderType = "OcoOrder"               //spot Oco order
	MmRateCloseStopOrderType            StopOrderType = "MmRateClose"            //On web or app can set MMR to close position
	BidirectionalTpslOrderStopOrderType StopOrderType = "BidirectionalTpslOrder" //Spot bidirectional tpsl order
)

type OrderStatus string

const (
	NewOrderStatus             OrderStatus = "New" //order has been placed successfully
	PartiallyFilledOrderStatus OrderStatus = "PartiallyFilled"
	UntriggeredOrderStatus     OrderStatus = "Untriggered" //Conditional orders are created
)

type CreateType string

const (
	CreateByUserCreateType                              CreateType = "CreateByUser"
	CreateByAdminClosingCreateType                      CreateType = "CreateByAdminClosing"
	CreateBySettleCreateType                            CreateType = "CreateBySettle"                             //USDC Futures delivery; Position closed by contract delisted
	CreateByStopOrderCreateType                         CreateType = "CreateByStopOrder"                          //Futures conditional order
	CreateByTakeProfitCreateType                        CreateType = "CreateByTakeProfit"                         //Futures take profit order
	CreateByPartialTakeProfitCreateType                 CreateType = "CreateByPartialTakeProfit"                  //Futures partial take profit order
	CreateByStopLossCreateType                          CreateType = "CreateByStopLoss"                           //Futures stop loss order
	CreateByPartialStopLossCreateType                   CreateType = "CreateByPartialStopLoss"                    //Futures partial stop loss order
	CreateByTrailingStopCreateType                      CreateType = "CreateByTrailingStop"                       //Futures trailing stop order
	CreateByLiqCreateType                               CreateType = "CreateByLiq"                                //Laddered liquidation to reduce the required maintenance margin
	CreateByTakeOverPassThroughIfCreateType             CreateType = "CreateByTakeOver_PassThroughIf"             //If the position is still subject to liquidation (i.e., does not meet the required maintenance margin level), the position shall be taken over by the liquidation engine and closed at the bankruptcy price.
	CreateByAdlPassThroughCreateType                    CreateType = "CreateByAdl_PassThrough"                    //Auto-Deleveraging(ADL)[https://www.bybit.com/en/help-center/article/Auto-Deleveraging-ADL]
	CreateByBlockPassThroughCreateType                  CreateType = "CreateByBlock_PassThrough"                  //Order placed via Paradigm
	CreateByBlockTradeMovePositionPassThroughCreateType CreateType = "CreateByBlockTradeMovePosition_PassThrough" //Order created by move position
	CreateByClosingCreateType                           CreateType = "CreateByClosing"                            //The close order placed via web or app position area - web/app
	CreateByFGridBotCreateType                          CreateType = "CreateByFGridBot"                           //Order created via grid bot - web/app
	CloseByFGridBotCreateType                           CreateType = "CloseByFGridBot"                            //Order closed via grid bot - web/app
	CreateByTWAPCreateType                              CreateType = "CreateByTWAP"                               //Order created by TWAP - web/app
	CreateByTVSignalCreateType                          CreateType = "CreateByTVSignal"                           //Order created by TV webhook - web/app
	CreateByMmRateCloseCreateType                       CreateType = "CreateByMmRateClose"                        //Order created by Mm rate close function - web/app
	CreateByMartingaleBotCreateType                     CreateType = "CreateByMartingaleBot"                      //Order created by Martingale bot - web/app
	CloseByMartingaleBotCreateType                      CreateType = "CloseByMartingaleBot"                       //Order closed by Martingale bot - web/app
	CreateByIceBergCreateType                           CreateType = "CreateByIceBerg"                            //Order created by Ice berg strategy - web/app
	CreateByArbitrageCreateType                         CreateType = "CreateByArbitrage"                          //Order created by arbitrage - web/app
	CreateByDdhCreateType                               CreateType = "CreateByDdh"                                //Option dynamic delta hedge order - web/app
)

// TODO: пока непонятно зачем надо
type RejectReason string

type Side string

const (
	BuySide  Side = "Buy"
	SellSide Side = "Sell"
)

type TriggerBy string

const (
	LastPriceTriggerBy  TriggerBy = "LastPrice"
	IndexPriceTriggerBy TriggerBy = "IndexPrice"
	MarkPriceTriggerBy  TriggerBy = "MarkPrice"
)

type TimeInForce string

const (
	GTCTimeInForce      TimeInForce = "GTC"      //GoodTillCancel
	IOCTimeInForce      TimeInForce = "IOC"      //ImmediateOrCancel
	FOKTimeInForce      TimeInForce = "FOK"      //FillOrKill
	PostOnlyTimeInForce TimeInForce = "PostOnly" //https://www.bybit.com/en/help-center/article/Post-Only-Order
)

type TriggerDirection uint

const (
	RiseTriggerDirection TriggerDirection = 1
	FallTriggerDirection TriggerDirection = 2
)

type PositionIdx uint

const (
	OneWayPositionIdx   PositionIdx = 0 //one-way mode position
	BuySidePositionIdx  PositionIdx = 1 //Buy side of hedge-mode position
	SellSidePositionIdx PositionIdx = 2 //Sell side of hedge-mode position
)

type TpslMode string

const (
	FullTpslMode    TpslMode = "Full"    //entire position for TP/SL
	PartialTpslMode TpslMode = "Partial" // partial position tp/sl
)

type SmpType string

const (
	NoneSmpType        SmpType = "None" //default
	CancelMakerSmpType SmpType = "CancelMaker"
	CancelTakerSmpType SmpType = "CancelTaker"
	CancelBothSmpType  SmpType = "CancelBoth"
)
