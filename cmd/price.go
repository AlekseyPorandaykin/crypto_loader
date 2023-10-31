package cmd

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/clients"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/loaders"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
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
	"os/signal"
	"syscall"
	"time"
)

var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Load prices from external sources",
	Run: func(cmd *cobra.Command, args []string) {
		const defaultDurationRequest = time.Minute

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
			Username: "crypto_loader",
			Password: "developer",
			Host:     "localhost",
			Port:     "5433",
			Database: "crypto_loader",
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

		go priceLoader.Run(ctx, defaultDurationRequest)

		<-ctx.Done()
	},
}

func init() {
	LoaderCmd.AddCommand(priceCmd)
}
