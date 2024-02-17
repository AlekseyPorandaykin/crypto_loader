package mempool_space

import "time"

type PriceDTO struct {
	CreatedAt  time.Time
	Currencies map[string]float64
}

// https://mempool.space/ru/docs/faq#looking-up-fee-estimates
type RecommendedFeesDTO struct {
	//sat/vB
	FastestFee  float64 `json:"fastestFee"`  //This figure is the median feerate of transactions in the first mempool block. Consider using this feerate if you want confirmation as soon as possible.
	HalfHourFee float64 `json:"halfHourFee"` //This figure is the average of the median feerate of the first mempool block and the median feerate of the second mempool block.
	HourFee     float64 `json:"hourFee"`     //This figure is the average of the Medium Priority feerate and the median feerate of the third mempool block. Consider using this feerate if you want confirmation soon but don't need it particularly quickly
	EconomyFee  float64 `json:"economyFee"`  //This figure is either 2x the minimum feerate, or the Low Priority feerate (whichever is lower). Consider using this feerate if you are in no rush and don't mind if confirmation takes a while.
	MinimumFee  float64 `json:"minimumFee"`
}
