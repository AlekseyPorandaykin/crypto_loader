package kraken

import (
	"context"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	c, err := NewClient("https://api.kraken.com/")
	if err != nil {
		return
	}
	res, _ := c.Ticker(context.TODO())
	for symbolPair, items := range res.Result {
		if symbolPair != "XBTUSDT" {
			continue
		}
		b, _ := items.AveragePrice()
		//for a, b := range items {
		//	fmt.Println(a, b)
		//}
		_ = symbolPair
		_, _ = items, b
	}
}
