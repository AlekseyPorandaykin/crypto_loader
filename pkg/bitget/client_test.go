package bitget

import (
	"context"
	"testing"
)

func TestClient_GetTicker(t *testing.T) {
	c, err := NewClient("https://api.bitget.com/")
	if err != nil {
		return
	}
	c.GetTicker(context.TODO())
}
