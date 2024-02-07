package mempool_space

import "time"

type PriceDTO struct {
	CreatedAt  time.Time
	Currencies map[string]float64
}

type RecommendedFeesDTO struct {
	FastestFee  float64 `json:"fastestFee"`
	HalfHourFee float64 `json:"halfHourFee"`
	HourFee     float64 `json:"hourFee"`
	EconomyFee  float64 `json:"economyFee"`
	MinimumFee  float64 `json:"minimumFee"`
}
