package binance

import (
	"context"
	"testing"
)

func TestClient_GetPrice(t *testing.T) {
	c, err := NewClient("https://api.binance.com")
	if err != nil {
		return
	}
	c.GetPrice(context.TODO())
}
