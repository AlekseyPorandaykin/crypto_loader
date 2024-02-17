package config

import (
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage/repositories"
	"time"
)

const defaultDurationPriceRequest = 20 * time.Second

type AppConfig struct {
	DurationPriceRequest time.Duration

	BinanceSpotHost   string
	BinanceFutureHost string
	BybitHost         string
	KucoinHost        string
	OkxHost           string
	GateioHost        string
	KrakenHost        string
	BitgetHost        string
	MexcHost          string

	//DB
	ConfDB repositories.Config

	HttpAddr   string
	GrpcAddr   string
	MetricAddr string
}

func Create() AppConfig {
	return AppConfig{
		DurationPriceRequest: defaultDurationPriceRequest,

		BinanceSpotHost:   "https://api.binance.com",
		BinanceFutureHost: "https://fapi.binance.com",
		//BinanceFutureHost: "https://testnet.binancefuture.com",

		BybitHost:  "https://api.bybit.com",
		KucoinHost: "https://api.kucoin.com/",
		OkxHost:    "https://www.okx.com/",
		GateioHost: "https://api.gateio.ws/",
		KrakenHost: "https://api.kraken.com/",
		BitgetHost: "https://api.bitget.com/",
		MexcHost:   "https://api.mexc.com/",
		ConfDB: repositories.Config{
			Driver:   "postgres",
			Username: "crypto_loader",
			Password: "developer",
			Host:     "localhost",
			Port:     "5433",
			Database: "crypto_loader",
		},
		HttpAddr:   ":8081",
		MetricAddr: ":9081",
		GrpcAddr:   ":50052",
	}
}
