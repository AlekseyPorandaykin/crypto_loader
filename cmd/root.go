package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/clients"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/candlestick"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/exchange"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/price"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/component/server/grpc"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/config"
	http_server "github.com/AlekseyPorandaykin/crypto_loader/internal/controller"
	candlestick_service "github.com/AlekseyPorandaykin/crypto_loader/internal/service/candlestick"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/service/order"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/service/symbol"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage/memory"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage/repositories"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	binance_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bitget"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5"
	bybit_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/bybit/v5/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/database"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/gateio"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kraken"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin"
	kukoin_sender "github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/sender"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/mexc"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/okx"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/server/http"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/system"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os/signal"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use: "crypto-loader", Short: "LoadPrices prices from external sources",
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
		db, err := database.CreateConnection(conf.ConfDB)
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

		priceLoader := price.NewPrice(priceStorage, symbolStorage)
		priceLoader.AddClient(domain.Binance, binanceAdapter)
		priceLoader.AddClient(domain.ByBit, clients.NewByBit(byBitClient))
		priceLoader.AddClient(domain.KuKoin, clients.NewKucoin(kukoinClient))
		priceLoader.AddClient(domain.Okx, clients.NewOkx(okxClient))
		priceLoader.AddClient(domain.GateIo, clients.NewGateIo(gateIoClient))
		priceLoader.AddClient(domain.Kraken, clients.NewKraken(krakenClient))
		priceLoader.AddClient(domain.BitGet, clients.NewBitget(bitgetClient))
		priceLoader.AddClient(domain.Mexc, clients.NewMexc(mexcClient))

		exRepo := memory.NewExchangeRepository()

		candlestickLoader := candlestick_service.NewCandlestick(candleStorage, priceStorage, symbolStorage, exRepo)

		candleLoader := candlestick.NewCandlestick(candlestickLoader)
		candleLoader.AddLoader(domain.Binance, binanceAdapter)

		order := order.NewOrder()
		order.AddExchange(domain.Binance, binanceAdapter)
		ex := exchange.NewExchange(exRepo)
		ex.AddSymbolInfoLoader(domain.Binance, binanceAdapter)

		symbolService := symbol.NewSymbol(candleStorage, priceStorage, symbolStorage, exRepo)

		//Servers
		s := http.NewServer()
		controllerService := http_server.NewController(conf.HttpAddr, priceStorage, order, candlestickLoader, symbolService)
		s.RegistrationPage(controllerService)
		s.RegistrationApi(controllerService)

		servGrpc := grpc.NewServer(priceStorage, conf.GrpcAddr)
		defer servGrpc.Close()

		//Runs
		system.Go(func() {
			priceStorage.Run(ctx)
		})
		system.Go(func() {
			priceLoader.Run(ctx, conf.DurationPriceRequest)
		})
		system.Go(func() {
			candleLoader.Run(ctx)
		})
		system.Go(func() {
			ex.Run(ctx)
		})
		system.Go(func() {
			defer cancel()
			if err := s.Run("localhost", conf.HttpAddr); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start serve", zap.Error(err))
			}
		})
		system.Go(func() {
			defer cancel()
			if err := servGrpc.Start(); err != nil && !errors.Is(err, context.Canceled) {
				zap.L().Error("failed start server", zap.Error(err))
			}
		})
		<-ctx.Done()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil && !errors.Is(err, context.Canceled) {
		zap.L().Error("execute root cmd", zap.Error(err))
	}
}
