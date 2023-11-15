package config

import (
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"time"
)

const defaultDurationPriceRequest = 20 * time.Second

type AppConfig struct {
	DurationPriceRequest time.Duration

	BinanceHost string
	BybitHost   string
	KucoinHost  string
	OkxHost     string
	GateioHost  string
	KrakenHost  string
	BitgetHost  string
	MexcHost    string

	//DB
	ConfDB repositories.Config

	HttpAddr string
	GrpcAddr string
}

func Create() AppConfig {
	return AppConfig{
		DurationPriceRequest: defaultDurationPriceRequest,
		BinanceHost:          "https://api.binance.com",
		BybitHost:            "https://api.bybit.com",
		KucoinHost:           "https://api.kucoin.com/",
		OkxHost:              "https://www.okx.com/",
		GateioHost:           "https://api.gateio.ws/",
		KrakenHost:           "https://api.kraken.com/",
		BitgetHost:           "https://api.bitget.com/",
		MexcHost:             "https://api.mexc.com/",
		ConfDB: repositories.Config{
			Driver:   "postgres",
			Username: "crypto_app",
			Password: "developer",
			Host:     "localhost",
			Port:     "5433",
			Database: "crypto_app",
		},
		HttpAddr: ":8081",
		GrpcAddr: ":50052",
	}
}
