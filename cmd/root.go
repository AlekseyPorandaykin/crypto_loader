package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/clients"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/loaders"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/grpc"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/server/http"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bitget"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/gateio"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kraken"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/mexc"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os/signal"
	"path"
	"syscall"
	"time"
)

var homeDir string = "./"

var cacheDir string = path.Join(homeDir, "storage/cache")

var rootCmd = &cobra.Command{
	Use: "crypto-loader", Short: "Load prices from external sources",
	Run: func(cmd *cobra.Command, args []string) {
		const defaultDurationRequest = 5 * time.Second

		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		//Clients
		binanceClient, err := binance.NewClient("https://api.binance.com")
		if err != nil {
			fmt.Println("Error init binanceClient: ", err.Error())
			return
		}
		byBitClient, err := bybit.NewClient("https://api.bybit.com")
		if err != nil {
			fmt.Println("Error init byBitClient: ", err.Error())
			return
		}
		kukoinClient, err := kucoin.NewClient("https://api.kucoin.com/")
		if err != nil {
			fmt.Println("Error init kukoinClient: ", err.Error())
			return
		}
		okxClient, err := okx.NewClient("https://www.okx.com/")
		if err != nil {
			fmt.Println("Error init okxClient: ", err.Error())
			return
		}
		gateIoClient, err := gateio.NewClient("https://api.gateio.ws/")
		if err != nil {
			fmt.Println("Error init gateIoClient: ", err.Error())
			return
		}
		krakenClient, err := kraken.NewClient("https://api.kraken.com/")
		if err != nil {
			fmt.Println("Error init krakenClient: ", err.Error())
			return
		}
		bitgetClient, err := bitget.NewClient("https://api.bitget.com/")
		if err != nil {
			fmt.Println("Error init bitgetClient: ", err.Error())
			return
		}
		mexcClient, err := mexc.NewClient("https://api.mexc.com/")
		if err != nil {
			fmt.Println("Error init mexcClient: ", err.Error())
			return
		}

		//DB
		db, err := repositories.CreateDB(repositories.Config{
			Driver:   "postgres",
			Username: "crypto_app",
			Password: "developer",
			Host:     "localhost",
			Port:     "5433",
			Database: "crypto_app",
		})
		if err != nil {
			fmt.Println("Error init database: ", err.Error())
			return
		}
		defer func() { _ = db.Close() }()

		//Repository
		priceRepo := repositories.NewPriceRepository(db)
		//Storage
		priceStorage := storage.NewPriceStorage(priceRepo, cacheDir)

		//Application

		priceLoader := loaders.NewPrice(priceStorage)
		priceLoader.AddClient("binance", clients.NewBinance(binanceClient))
		priceLoader.AddClient("byBit", clients.NewByBit(byBitClient))
		priceLoader.AddClient("kukoin", clients.NewKucoin(kukoinClient))
		priceLoader.AddClient("okx", clients.NewOkx(okxClient))
		priceLoader.AddClient("gate.io", clients.NewGateIo(gateIoClient))
		priceLoader.AddClient("kraken", clients.NewKraken(krakenClient))
		priceLoader.AddClient("bitget", clients.NewBitget(bitgetClient))
		priceLoader.AddClient("mexc", clients.NewMexc(mexcClient))

		//Servers
		servHTTP := http.NewServer(":8081", priceStorage)
		defer servHTTP.Close()

		servGrpc := grpc.NewServer(priceStorage, ":50052")
		defer servGrpc.Close()

		//Runs
		go priceStorage.Run(ctx)
		go priceLoader.Run(ctx, defaultDurationRequest)
		go func() {
			defer cancel()
			if err := servHTTP.Run(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		}()
		go func() {
			defer cancel()
			if err := servGrpc.Start(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		}()

		<-ctx.Done()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
		zap.L().Error("execute root cmd", zap.Error(err))
	}
}
