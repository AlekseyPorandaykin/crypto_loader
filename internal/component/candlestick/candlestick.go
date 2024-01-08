package candlestick

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"time"
)

type Candlestick struct {
	candlestickStorage *storage.Candlestick
	priceStorage       *storage.Price
	symbolStorage      *storage.Symbol
}

func NewCandlestick(
	candlestickStorage *storage.Candlestick, priceStorage *storage.Price, symbolStorage *storage.Symbol,
) *Candlestick {
	return &Candlestick{
		candlestickStorage: candlestickStorage,
		priceStorage:       priceStorage,
		symbolStorage:      symbolStorage,
	}
}

func (a *Candlestick) SymbolSnapshot(ctx context.Context, exchange, symbol string) (domain.SymbolSnapshot, error) {
	price, err := a.priceStorage.LastPrice(ctx, exchange, symbol)
	if err != nil {
		return domain.SymbolSnapshot{}, errors.Wrap(err, "error get lastPrice")
	}
	if !a.symbolStorage.HasExchange(exchange) {
		return domain.SymbolSnapshot{}, errors.New("not found exchange")
	}
	if !a.symbolStorage.HasSymbol(exchange, symbol) {
		return domain.SymbolSnapshot{}, errors.New("not found symbol")
	}

	return domain.SymbolSnapshot{
		Symbol:        symbol,
		Exchange:      exchange,
		Price:         price.Price,
		CreatedAt:     time.Now().In(time.UTC),
		Candlestick1H: a.candlestickStorage.LastCandlestick(ctx, exchange, symbol, domain.OneHourCandlestickInterval),
		Candlestick4H: a.candlestickStorage.LastCandlestick(ctx, exchange, symbol, domain.FourHourCandlestickInterval),
	}, nil
}

func (a *Candlestick) OneHourCandlesticks(ctx context.Context, exchange, symbol string) ([]dto.Candlestick, error) {
	if !a.symbolStorage.HasExchange(exchange) {
		return nil, errors.New("not found exchange")
	}
	if !a.symbolStorage.HasSymbol(exchange, symbol) {
		return nil, errors.New("not found symbol")
	}
	candles := a.candlestickStorage.Candlestick(ctx, exchange, symbol, domain.OneHourCandlestickInterval)
	res := make([]dto.Candlestick, 0, len(candles))
	for _, candle := range candles {
		res = append(res, candlesticksToDTO(candle))
	}
	return res, nil
}

func (a *Candlestick) FourHourCandlesticks(ctx context.Context, exchange, symbol string) ([]dto.Candlestick, error) {
	if !a.symbolStorage.HasExchange(exchange) {
		return nil, errors.New("not found exchange")
	}
	if !a.symbolStorage.HasSymbol(exchange, symbol) {
		return nil, errors.New("not found symbol")
	}
	candles := a.candlestickStorage.Candlestick(ctx, exchange, symbol, domain.FourHourCandlestickInterval)
	res := make([]dto.Candlestick, 0, len(candles))
	for _, candle := range candles {
		res = append(res, candlesticksToDTO(candle))
	}
	return res, nil
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
