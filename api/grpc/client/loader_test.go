package client

import (
	"context"
	"fmt"
	"testing"
)

func TestLoader_Start(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	l := DefaultLoader()

	go func() {
		defer cancel()
		if err := l.Start(ctx, 10); err != nil {
			t.Error(err)
			return
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case batch := <-l.Batch():
			fmt.Printf("SymbolPrices: %d \n", len(batch.Prices()))
		}
	}
}
