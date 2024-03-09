package config

import (
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/database"
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
	ConfDB database.Config

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
		//ConfDB: database.Config{
		//	Driver:   "postgres",
		//	Username: "crypto_app",
		//	Password: "developer",
		//	Host:     "localhost",
		//	Port:     "5433",
		//	Database: "crypto_app",
		//},
		ConfDB: database.Config{
			Driver:   "sqlite",
			PathToDB: "/Users/alexey.porandaikin/Projects/go/projects/crypto_loader/storage/crypto_loader.db",
		},
		HttpAddr:   "8081",
		MetricAddr: ":9081",
		GrpcAddr:   ":50052",
	}
}
