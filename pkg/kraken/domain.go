package kraken

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type CommonResponse struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

type Tick struct {
	// https://docs.kraken.com/rest/#tag/Market-Data/operation/getTickerInformation
	Ask                        []string `json:"a"`
	Bid                        []string `json:"b"`
	LastTradeClosed            []string `json:"c"`
	Volume                     []string `json:"v"`
	VolumeWeightedAveragePrice []string `json:"p"`
	NumberOfTrades             []int64  `json:"t"`
	Low                        []string `json:"l"`
	High                       []string `json:"h"`
	TodayOpeningPrice          string   `json:"o"`
}

func (t Tick) AveragePrice() (string, error) {
	if len(t.Ask) == 0 || len(t.Bid) == 0 {
		return "", errors.New("empty price")
	}
	ask, err := strconv.ParseFloat(t.Ask[0], 64)
	if err != nil {
		return "", errors.Wrap(err, "parse ask")
	}
	bid, err := strconv.ParseFloat(t.Bid[0], 64)
	if err != nil {
		return "", errors.Wrap(err, "parse bid")
	}
	res := (ask + bid) / 2
	return fmt.Sprintf("%f", res), nil
}

type SymbolPair string

type TickerResponse struct {
	CommonResponse
	Result map[SymbolPair]Tick `json:"result"`
}
