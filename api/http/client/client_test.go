package client

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Prices(t *testing.T) {
	prices, err := DefaultClient().Prices(context.TODO(), "BTCUSDT")
	require.NoError(t, err)
	require.NotEmpty(t, prices)
}

func TestClient_AllSymbolPrices(t *testing.T) {
	prices, err := DefaultClient().AllSymbolPrices(context.TODO())
	require.NoError(t, err)
	require.NotEmpty(t, prices)
}

func TestClient_SymbolSnapshot(t *testing.T) {
	snapshot, err := DefaultClient().SymbolSnapshot(context.TODO(), "binance", "BTCUSDT")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}

func TestClient_OneHourCandlesticks(t *testing.T) {
	snapshot, err := DefaultClient().OneHourCandlesticks(context.TODO(), "binance", "BTCUSDT")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}

func TestClient_FourHourCandlesticks(t *testing.T) {
	snapshot, err := DefaultClient().FourHourCandlesticks(context.TODO(), "binance", "BTCUSDT")
	require.NoError(t, err)
	require.NotEmpty(t, snapshot)
}
