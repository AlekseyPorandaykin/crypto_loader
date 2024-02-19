package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/clients"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/candlestick"
	http_server "github.com/AlekseyPorandaykin/crypto_loader/internal/component/controller"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/loaders"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/order"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/server/grpc"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/config"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage/repositories"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	binance_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bitget"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5"
	bybit_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/gateio"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kraken"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin"
	kukoin_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/mexc"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/server/http"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/shutdown"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use: "crypto-loader", Short: "Load prices from external sources",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		conf := config.Create()

		//Clients
		binanceClient, err := binance.NewManager(conf.BinanceSpotHost, conf.BinanceFutureHost)
		if err != nil {
			fmt.Println("Error init binanceClient: ", err.Error())
			return
		}
		defer binanceClient.Close()
		binanceClient.WithSender(binance_sender.NewBasic())
		byBitClient, err := v5.NewClient(conf.BybitHost, bybit_sender.NewBasic())
		if err != nil {
			fmt.Println("Error init byBitClient: ", err.Error())
			return
		}
		kukoinClient, err := kucoin.NewClient(conf.KucoinHost, kukoin_sender.New())
		if err != nil {
			fmt.Println("Error init kukoinClient: ", err.Error())
			return
		}
		okxClient, err := okx.NewClient(conf.OkxHost)
		if err != nil {
			fmt.Println("Error init okxClient: ", err.Error())
			return
		}
		gateIoClient, err := gateio.NewClient(conf.GateioHost)
		if err != nil {
			fmt.Println("Error init gateIoClient: ", err.Error())
			return
		}
		krakenClient, err := kraken.NewClient(conf.KrakenHost)
		if err != nil {
			fmt.Println("Error init krakenClient: ", err.Error())
			return
		}
		bitgetClient, err := bitget.NewClient(conf.BitgetHost)
		if err != nil {
			fmt.Println("Error init bitgetClient: ", err.Error())
			return
		}
		mexcClient, err := mexc.NewClient(conf.MexcHost)
		if err != nil {
			fmt.Println("Error init mexcClient: ", err.Error())
			return
		}

		//DB
		db, err := repositories.CreateDB(conf.ConfDB)
		if err != nil {
			fmt.Println("Error init database: ", err.Error())
			return
		}
		defer func() { _ = db.Close() }()

		//Repository
		priceRepo := repositories.NewPriceRepository(db)
		//Storage
		priceStorage := storage.NewPriceStorage(priceRepo)
		symbolStorage := storage.NewSymbol()
		candleStorage := storage.NewCandlestick()

		//Clients
		binanceAdapter := clients.NewBinance(binanceClient)

		//Application

		priceLoader := loaders.NewPrice(priceStorage, symbolStorage)
		priceLoader.AddClient("binance", binanceAdapter)
		priceLoader.AddClient("byBit", clients.NewByBit(byBitClient))
		priceLoader.AddClient("kukoin", clients.NewKucoin(kukoinClient))
		priceLoader.AddClient("okx", clients.NewOkx(okxClient))
		priceLoader.AddClient("gate.io", clients.NewGateIo(gateIoClient))
		priceLoader.AddClient("kraken", clients.NewKraken(krakenClient))
		priceLoader.AddClient("bitget", clients.NewBitget(bitgetClient))
		priceLoader.AddClient("mexc", clients.NewMexc(mexcClient))

		candleLoader := loaders.NewCandlestick(symbolStorage, candleStorage)
		candleLoader.AddLoader("binance", binanceAdapter)

		order := order.NewOrder()
		order.AddExchange("binance", binanceAdapter)

		agg := candlestick.NewCandlestick(candleStorage, priceStorage, symbolStorage)

		//Servers
		s := http.NewServer()
		controllerService := http_server.NewController(conf.HttpAddr, priceStorage, order, agg)
		s.RegistrationPage(controllerService)

		servGrpc := grpc.NewServer(priceStorage, conf.GrpcAddr)
		defer servGrpc.Close()

		//Runs
		go func() {
			defer shutdown.HandlePanic()
			priceStorage.Run(ctx)
		}()
		go func() {
			defer shutdown.HandlePanic()
			priceLoader.Run(ctx, conf.DurationPriceRequest)
		}()
		go func() {
			defer shutdown.HandlePanic()
			candleLoader.Run(ctx)
		}()
		go func() {
			defer shutdown.HandlePanic()
			defer cancel()
			if err := s.Run("localhost", conf.HttpAddr); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		}()
		go func() {
			defer shutdown.HandlePanic()
			defer cancel()
			if err := servGrpc.Start(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start server", zap.Error(err))
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
