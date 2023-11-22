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
	c.FuturesNewOrder(
		domain.CredentialDTO{APIKey: apiKeyFuture, ApiSecret: apiSecretFuture},
		domain.NewStopMarketFutureOrder("BTCUSDT", "BUY", 0.01, 37200),
	)
}

func saveToFile(data []byte) {
	file, err := os.Create(fmt.Sprintf("/Users/alexey.porandaikin/Projects/go/projects/crypto_loader/storage/temp/%s.json", uuid.New().String()))
	if err != nil {
		return
	}
	defer file.Close()
	file.Write(data)
}
