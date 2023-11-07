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