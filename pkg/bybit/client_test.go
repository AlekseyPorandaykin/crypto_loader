package bybit

import (
	"context"
	"testing"
)

func TestClient_SpotTicker(t *testing.T) {
	c, err := NewClient("https://api.bybit.com/")
	if err != nil {
		return
	}
	c.SpotTicker(context.TODO())
}
