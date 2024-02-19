package loaders

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/metric"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/shutdown"
	"go.uber.org/zap"
	"time"
)

type Loader interface {
	FutureCandlestickOneHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
	FutureCandlestickFourHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
}

type Candlestick struct {
	loaders            map[string]Loader
	symbolStorage      *storage.Symbol
	candlestickStorage *storage.Candlestick
}

func NewCandlestick(symbolStorage *storage.Symbol, candlestickStorage *storage.Candlestick) *Candlestick {
	return &Candlestick{
		symbolStorage:      symbolStorage,
		candlestickStorage: candlestickStorage,
		loaders:            make(map[string]Loader),
	}
}

func (c *Candlestick) AddLoader(exchange string, loader Loader) {
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
	go func() {
		defer shutdown.HandlePanic()
		c.loadFutureCandlestickOneHour(ctx, exchange, loader)
		ticker := time.NewTicker(100 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ticker.Stop()
				c.loadFutureCandlestickOneHour(ctx, exchange, loader)
				ticker.Reset(durationToNextHour())
			}
		}
	}()

	go func() {
		defer shutdown.HandlePanic()
		c.loadFutureCandlestickFourHour(ctx, exchange, loader)
		ticker := time.NewTicker(100 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				ticker.Stop()
				c.loadFutureCandlestickFourHour(ctx, exchange, loader)
				ticker.Reset(durationToNextHour())
			}
		}
	}()
}

func (c *Candlestick) loadFutureCandlestickOneHour(ctx context.Context, exchange string, loader Loader) {
	interval := domain.OneHourCandlestickInterval
	defer func(now time.Time) {
		metric.CandlestickLoadDuration.WithLabelValues(exchange, string(interval)).Add(float64(time.Since(now).Milliseconds()))
	}(time.Now())
	log := zap.L().With(zap.String("exchange", exchange))
	for _, symbol := range domain.SymbolCandlestick {
		candlesticks, err := loader.FutureCandlestickOneHour(ctx, symbol)
		if err != nil {
			log.Error("error load FutureCandlestickOneHour", zap.Error(err), zap.String("symbol", symbol))
			metric.CandlestickError.WithLabelValues(exchange, string(interval)).Inc()
			continue
		}
		if err := c.candlestickStorage.Save(ctx, candlesticks, exchange, symbol, interval); err != nil {
			log.Error("error save FutureCandlestickOneHour", zap.Error(err), zap.String("symbol", symbol))
		}
		metric.CandlestickSaved.WithLabelValues(exchange, string(interval)).Inc()
	}
}

func (c *Candlestick) loadFutureCandlestickFourHour(ctx context.Context, exchange string, loader Loader) {
	interval := domain.FourHourCandlestickInterval
	defer func(now time.Time) {
		metric.CandlestickLoadDuration.WithLabelValues(exchange, string(interval)).Add(float64(time.Since(now).Milliseconds()))
	}(time.Now())
	log := zap.L().With(zap.String("exchange", exchange))
	for _, symbol := range domain.SymbolCandlestick {
		candlesticks, err := loader.FutureCandlestickFourHour(ctx, symbol)
		if err != nil {
			log.Error("error load FutureCandlestickFourHour", zap.Error(err), zap.String("symbol", symbol))
			metric.CandlestickError.WithLabelValues(exchange, string(interval)).Inc()
			continue
		}
		if err := c.candlestickStorage.Save(ctx, candlesticks, exchange, symbol, interval); err != nil {
			log.Error("error save FutureCandlestickFourHour", zap.Error(err), zap.String("symbol", symbol))
		}
		metric.CandlestickSaved.WithLabelValues(exchange, string(interval)).Inc()
	}
}

func durationToNextHour() time.Duration {
	now := time.Now().In(time.UTC)
	nextExecute := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 10, 0, time.UTC)
	return nextExecute.Sub(now)
}
