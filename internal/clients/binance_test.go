package clients

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/binance"
	"testing"
)

func TestBinance_LoadSymbolInfo(t *testing.T) {
	binanceManager, err := binance.NewManager("https://api.binance.com", "https://testnet.binancefuture.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, _ := NewBinance(binanceManager).LoadSymbolInfo(context.TODO())
	_ = res
}
func TestBinance_FutureCandlestick(t *testing.T) {
	binanceManager, err := binance.NewManager("https://api.binance.com", "https://testnet.binancefuture.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	res, _ := NewBinance(binanceManager).FutureCandlestickOneHour(context.TODO(), "BTCUSDT")
	_ = res
}
