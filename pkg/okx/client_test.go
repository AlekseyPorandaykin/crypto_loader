package okx

import (
	"context"
	"testing"
)

func TestClient_Tickers(t *testing.T) {
	c, err := NewClient("https://www.okx.com/")
	if err != nil {
		return
	}
	c.Tickers(context.TODO())
}
