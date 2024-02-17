package client

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_SymbolPrices(t *testing.T) {
	prices, err := DefaultClient().SymbolPrices(context.TODO(), "BTCUSDT")
	require.NoError(t, err)
	require.NotEmpty(t, prices)
}
func TestClient_ExchangePrices(t *testing.T) {
	prices, err := DefaultClient().ExchangePrices(context.TODO(), "bybit")
	require.NoError(t, err)
	require.NotEmpty(t, prices)
}

func TestClient_AllSymbolPrices(t *testing.T) {
	prices, err := DefaultClient().AllSymbolPrices(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, prices)
}

func TestClient_SymbolSnapshot(t *testing.T) {
	snapshot, err := DefaultClient().SymbolSnapshot(context.TODO(), "binance", "BTCUSDC")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}

func TestClient_OneHourCandlesticks(t *testing.T) {
	snapshot, err := DefaultClient().Candlesticks(context.TODO(), "binance", "BTCUSDT", "1h")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}

func TestClient_FourHourCandlesticks(t *testing.T) {
	snapshot, err := DefaultClient().Candlesticks(context.TODO(), "binance", "BTCUSDT", "4h")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}
