package aggregator

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/dto"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"time"
)

type Aggregator struct {
	candlestickStorage *storage.Candlestick
	priceStorage       *storage.Price
	symbolStorage      *storage.Symbol
}

func NewAggregator(
	candlestickStorage *storage.Candlestick, priceStorage *storage.Price, symbolStorage *storage.Symbol,
) *Aggregator {
	return &Aggregator{
		candlestickStorage: candlestickStorage,
		priceStorage:       priceStorage,
		symbolStorage:      symbolStorage,
	}
}

func (a *Aggregator) SymbolSnapshot(ctx context.Context, exchange, symbol string) (domain.SymbolSnapshot, error) {
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

func candlesticksToDTO(data domain.Candlestick) dto.Candlestick {
	var openTime, closeTime, createdAt string
	if !data.OpenTime.IsZero() {
		openTime = data.OpenTime.Format(time.DateTime)
	}
	if !data.CloseTime.IsZero() {
		closeTime = data.CloseTime.Format(time.DateTime)
	}
	if !data.CreatedAt.IsZero() {
		createdAt = data.CreatedAt.Format(time.DateTime)
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
