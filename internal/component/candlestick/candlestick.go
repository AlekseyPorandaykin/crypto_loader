package candlestick

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/service/candlestick"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/system"
	"time"
)

type CandleLoader interface {
	FutureCandlestickOneHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
	FutureCandlestickFourHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
}

type Candlestick struct {
	candleService *candlestick.Candlestick
	exchanges     []string

	loaders map[string]CandleLoader
}

func NewCandlestick(candleService *candlestick.Candlestick) *Candlestick {
	return &Candlestick{
		loaders:       make(map[string]CandleLoader),
		candleService: candleService,
	}
}

func (c *Candlestick) AddLoader(exchange string, loader CandleLoader) {
	c.loaders[exchange] = loader
}

func (c *Candlestick) Run(ctx context.Context) {
	for ex := range c.loaders {
		go c.load(ctx, ex)
	}
}

func (c *Candlestick) load(ctx context.Context, exchange string) {
	loader, has := c.loaders[exchange]
	if !has {
		return
	}
	system.Go(func() {
		c.candleService.LoadFutureCandlestickOneHour(ctx, exchange, loader)
		ticker := time.NewTicker(100 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ticker.Stop()
				c.candleService.LoadFutureCandlestickOneHour(ctx, exchange, loader)
				ticker.Reset(durationToNextHour())
			}
		}
	})

	system.Go(func() {
		c.candleService.LoadFutureCandlestickFourHour(ctx, exchange, loader)
		ticker := time.NewTicker(100 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ticker.Stop()
				c.candleService.LoadFutureCandlestickFourHour(ctx, exchange, loader)
				ticker.Reset(durationToNextHour())
			}
		}
	})
}

func durationToNextHour() time.Duration {
	now := time.Now().In(time.UTC)
	nextExecute := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 10, 0, time.UTC)
	return nextExecute.Sub(now)
}
