package binance

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance/sender"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
	"testing"
)

var (
	apiKeySpot    = "556O2t88HhtoXMbigHKKqeQ7UIAJZCrfYNLQxO2s9oEb5H1bd36ryVcqxydPUFII"
	secretKeySpot = "5gJhXukIYfZNVkiSUWIn0jUauU5VJuWmGmN14quB4WjGP2K9bg1JZiVpUlecDp29"

	apiKeyFuture    = "c63b09c49650da794de88d2388e7e70007a3aa3619be426b354d991af7bdea94"
	apiSecretFuture = "72bd120dfcaf91de3b483df4187c6132edb169f81156ba85ae6494f06174116e"
)

func TestClient_GetPrice(t *testing.T) {
	c, err := NewManager("https://api.binance.com", "futureHost")
	if err != nil {
		return
	}
	defer c.Close()
	c.WithSender(sender.NewLogger(zap.L(), sender.NewBasic()))
	c.GetPrice(context.TODO())
	data, err := c.PriceChangeStatisticsLastHour(context.TODO(), []string{"BTCUSDT", "BNBUSDT", "ETHUSDT"})
	if err != nil {
		return
	}
	_ = data
}
func TestClient_Ping(t *testing.T) {
	c, err := NewManager("https://api.binance.com", "https://testnet.binancefuture.com")
	if err != nil {
		return
	}
	defer c.Close()
	c.WithSender(sender.NewLogger(zap.L(), sender.NewBasic()))
	data, _ := c.FuturesExchangeInformation(context.TODO())
	saveToFile(data)
	//c.FutureCandlestickData(context.TODO(), "BTCUSDT", "1m")
	//c.FutureQueryOrder(
	//	context.TODO(),
	//	domain.CredentialDTO{APIKey: apiKeyFuture, ApiSecret: apiSecretFuture},
	//	"BTCUSDT",
	//	3543846157)
	//c.FutureCancelMultipleOrders(
	//	domain.CredentialDTO{APIKey: apiKeyFuture, ApiSecret: apiSecretFuture},
	//	"BTCUSDT",
	//	[]int{3543918994, 3543918996, 3543918995})
}

func saveToFile(data []byte) {
	fileName := uuid.New().String()
	file, err := os.Create(fmt.Sprintf("/Users/alexey.porandaikin/Projects/go/projects/crypto_loader/storage/temp/%s.json", fileName))
	if err != nil {
		return
	}
	defer file.Close()
	file.Write(data)
	fmt.Println("save file in : ", fileName)
}

func TestManager_FuturesPlaceMultipleOrders(t *testing.T) {
	c, err := NewManager("https://api.binance.com", "https://testnet.binancefuture.com")
	if err != nil {
		return
	}
	defer c.Close()
	c.WithSender(sender.NewLogger(zap.L(), sender.NewBasic()))
	order := domain.NewLimitFutureOrder(
		"BTCUSDT",
		domain.BuyOrderSide,
		domain.GtcTimeInForce,
		0.026,
		38200,
		false,
	)
	order.PositionSide = domain.BothPositionSide
	_, err = c.FuturesNewOrder(
		domain.CredentialDTO{APIKey: apiKeyFuture, ApiSecret: apiSecretFuture},
		order,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	takeProfitOrder := domain.NewTakeProfitMarketFutureOrder(
		"BTCUSDT", domain.SellOrderSide, 0.026, 38400, true,
	)
	takeProfitOrder.PositionSide = domain.BothPositionSide
	takeProfitOrder.PriceProtect = true
	takeProfitOrder.TimeInForce = domain.GtcTimeInForce
	takeProfitOrder.WorkingType = domain.MarkPriceWorkingType

	stopMarketOrder := domain.NewStopMarketFutureOrder(
		"BTCUSDT", domain.SellOrderSide, 0.026, 38000, true,
	)
	stopMarketOrder.PositionSide = domain.BothPositionSide
	stopMarketOrder.TimeInForce = domain.GtcTimeInForce
	stopMarketOrder.WorkingType = domain.MarkPriceWorkingType
	stopMarketOrder.PriceProtect = true

	orders := make([]domain.FutureOrder, 0, 10)
	orders = append(orders, takeProfitOrder, stopMarketOrder)
	_, err = c.FuturesPlaceMultipleOrders(
		domain.CredentialDTO{APIKey: apiKeyFuture, ApiSecret: apiSecretFuture},
		orders,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}
