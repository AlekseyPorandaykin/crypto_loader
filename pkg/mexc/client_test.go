package mexc

import (
	"context"
	"testing"
)

func TestClient_SymbolPriceTicker(t *testing.T) {
	c, err := NewClient("https://api.mexc.com/")
	if err != nil {
		return
	}
	c.SymbolPriceTicker(context.TODO())
}
