package kucoin

import (
	"context"
	"testing"
)

func TestClient_GetAllTickers(t *testing.T) {
	c, err := NewClient("https://api.kucoin.com/")
	if err != nil {
		return
	}
	c.GetAllTickers(context.TODO())
}
