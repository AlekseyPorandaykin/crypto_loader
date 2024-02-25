package symbol

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"github.com/pkg/errors"
	"time"
)

type Symbol struct {
	candlestickStorage *storage.Candlestick
	priceStorage       *storage.Price
	exchangeStorage    domain.ExchangeStorage
	symbolStorage      *storage.Symbol
}

func NewSymbol(
	candlestickStorage *storage.Candlestick,
	priceStorage *storage.Price,
	symbolStorage *storage.Symbol,
	exchangeStorage domain.ExchangeStorage,
) *Symbol {
	return &Symbol{
		candlestickStorage: candlestickStorage,
		priceStorage:       priceStorage,
		exchangeStorage:    exchangeStorage,
		symbolStorage:      symbolStorage,
	}
}

func (c *Symbol) SymbolSnapshot(ctx context.Context, exchange, symbol string) (domain.SymbolSnapshot, error) {
	price, err := c.priceStorage.LastPrice(ctx, exchange, symbol)
	if err != nil {
		return domain.SymbolSnapshot{}, errors.Wrap(err, "error get lastPrice")
	}
	symbolInfo, err := c.exchangeStorage.InfoBySymbol(ctx, symbol)

	if !c.symbolStorage.HasExchange(exchange) {
		return domain.SymbolSnapshot{}, errors.New("not found exchange")
	}
	if !c.symbolStorage.HasSymbol(exchange, symbol) {
		return domain.SymbolSnapshot{}, errors.New("not found symbol")
	}

	return domain.SymbolSnapshot{
		Symbol:        symbol,
		BaseAsset:     symbolInfo.BaseAsset,
		QuoteAsset:    symbolInfo.QuoteAsset,
		Exchange:      exchange,
		Price:         price.Price,
		PriceUpdated:  price.UpdatedAt,
		CreatedAt:     time.Now().In(time.UTC),
		Candlestick1H: c.candlestickStorage.LastCandlestick(ctx, exchange, symbol, domain.OneHourCandlestickInterval),
		Candlestick4H: c.candlestickStorage.LastCandlestick(ctx, exchange, symbol, domain.FourHourCandlestickInterval),
	}, nil
}
