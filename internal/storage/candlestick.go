package storage

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"sync"
	"time"
)

type Candlestick struct {
	muCandleEx     sync.Mutex
	candleExchange map[string]map[string]map[domain.CandlestickInterval][]domain.Candlestick

	muLastCandle sync.Mutex
	lastCandle   map[string]map[string]map[domain.CandlestickInterval]domain.Candlestick
}

func NewCandlestick() *Candlestick {
	return &Candlestick{
		candleExchange: make(map[string]map[string]map[domain.CandlestickInterval][]domain.Candlestick),
		lastCandle:     make(map[string]map[string]map[domain.CandlestickInterval]domain.Candlestick),
	}
}

func (c *Candlestick) Save(ctx context.Context, candlesticks []domain.Candlestick, exchange, symbol string) error {
	now := time.Now().In(time.UTC)
	lastTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	candleEx := make(map[domain.CandlestickInterval][]domain.Candlestick)
	var lastCandle domain.Candlestick
	for _, item := range candlesticks {
		candleEx[item.Interval] = append(candleEx[item.Interval], item)

		if lastCandle.CloseTime.Before(item.CloseTime) && lastTime.After(item.CloseTime) {
			lastCandle = item
		}
	}
	c.muCandleEx.Lock()
	if c.candleExchange[exchange] == nil {
		c.candleExchange[exchange] = map[string]map[domain.CandlestickInterval][]domain.Candlestick{}
	}
	if c.candleExchange[exchange][symbol] == nil {
		c.candleExchange[exchange][symbol] = map[domain.CandlestickInterval][]domain.Candlestick{}
	}
	c.candleExchange[exchange][symbol] = candleEx
	c.muCandleEx.Unlock()

	c.muLastCandle.Lock()
	if c.lastCandle[exchange] == nil {
		c.lastCandle[exchange] = map[string]map[domain.CandlestickInterval]domain.Candlestick{}
	}
	if c.lastCandle[exchange][symbol] == nil {
		c.lastCandle[exchange][symbol] = map[domain.CandlestickInterval]domain.Candlestick{}
	}
	c.lastCandle[exchange][symbol][lastCandle.Interval] = lastCandle
	c.muLastCandle.Unlock()
	return nil
}

func (c *Candlestick) LastCandlestick(
	ctx context.Context, exchange, symbol string, interval domain.CandlestickInterval,
) domain.Candlestick {
	if c.lastCandle[exchange] == nil || c.lastCandle[exchange][symbol] == nil {
		return domain.Candlestick{}
	}
	candle, has := c.lastCandle[exchange][symbol][interval]
	if !has {
		return domain.Candlestick{}
	}
	return candle
}

func (c *Candlestick) Candlestick(
	ctx context.Context, exchange, symbol string,
) map[domain.CandlestickInterval][]domain.Candlestick {
	if c.candleExchange[exchange] == nil || c.lastCandle[exchange][symbol] == nil {
		return nil
	}

	return c.candleExchange[exchange][symbol]
}
