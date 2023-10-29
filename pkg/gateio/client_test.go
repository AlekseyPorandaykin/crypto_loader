package gateio

import (
	"context"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	c, err := NewClient("https://api.gateio.ws/")
	if err != nil {
		return
	}
	c.Ticker(context.TODO())
}
