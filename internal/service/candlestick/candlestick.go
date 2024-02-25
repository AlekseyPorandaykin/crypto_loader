package candlestick

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/metric"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Loader interface {
	FutureCandlestickOneHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
	FutureCandlestickFourHour(ctx context.Context, symbol string) ([]domain.Candlestick, error)
}

type Candlestick struct {
	loaders            map[string]Loader
	candlestickStorage *storage.Candlestick
	priceStorage       *storage.Price
	exchangeStorage    domain.ExchangeStorage
	symbolStorage      *storage.Symbol
}

func NewCandlestick(
	candlestickStorage *storage.Candlestick, priceStorage *storage.Price, symbolStorage *storage.Symbol, exchangeStorage domain.ExchangeStorage,
) *Candlestick {
	return &Candlestick{
		candlestickStorage: candlestickStorage,
		priceStorage:       priceStorage,
		symbolStorage:      symbolStorage,
		exchangeStorage:    exchangeStorage,
		loaders:            make(map[string]Loader),
	}
}

func (c *Candlestick) AddLoader(exchange string, loader Loader) {
	c.loaders[exchange] = loader
}

func (c *Candlestick) Candlesticks(ctx context.Context, exchange, symbol string, interval domain.CandlestickInterval) (
	[]dto.Candlestick, error,
) {
	if !c.symbolStorage.HasExchange(exchange) {
		return nil, errors.New("not found exchange")
	}
	if !c.symbolStorage.HasSymbol(exchange, symbol) {
		return nil, errors.New("not found symbol")
	}
	if !domain.HasCandlestickInterval(interval) {
		return nil, errors.New("not found interval")
	}
	candles := c.candlestickStorage.Candlestick(ctx, exchange, symbol, interval)
	res := make([]dto.Candlestick, 0, len(candles))
	for _, candle := range candles {
		res = append(res, candlesticksToDTO(candle))
	}
	return res, nil
}

func (c *Candlestick) LoadFutureCandlestickOneHour(ctx context.Context, exchange string, loader Loader) {
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

func (c *Candlestick) LoadFutureCandlestickFourHour(ctx context.Context, exchange string, loader Loader) {
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

func candlesticksToDTO(data domain.Candlestick) dto.Candlestick {
	var openTime, closeTime, createdAt string
	if !data.OpenTime.IsZero() {
		openTime = data.OpenTime.In(time.UTC).Format(time.RFC3339)
	}
	if !data.CloseTime.IsZero() {
		closeTime = data.CloseTime.In(time.UTC).Format(time.RFC3339)
	}
	if !data.CreatedAt.IsZero() {
		createdAt = data.CreatedAt.In(time.UTC).Format(time.RFC3339)
	}
	return dto.Candlestick{
		Symbol:       data.Symbol,
		Exchange:     data.Exchange,
		OpenTime:     openTime,
		CloseTime:    closeTime,
		OpenPrice:    data.OpenPrice,
		HighPrice:    data.HighPrice,
		LowPrice:     data.LowPrice,
		ClosePrice:   data.ClosePrice,
		Volume:       data.Volume,
		NumberTrades: data.NumberTrades,
		Interval:     string(data.Interval),
		CreatedAt:    createdAt,
	}
}
