package exchange

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/system"
	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"time"
)

type SymbolInfoLoader interface {
	LoadSymbolInfo(ctx context.Context) ([]domain.SymbolInfo, error)
}

type Exchange struct {
	symbolInfoLoaders map[string]SymbolInfoLoader
	exchangeStorage   domain.ExchangeStorage
}

func NewExchange(exchangeStorage domain.ExchangeStorage) *Exchange {
	return &Exchange{
		symbolInfoLoaders: make(map[string]SymbolInfoLoader),
		exchangeStorage:   exchangeStorage,
	}
}

func (e *Exchange) AddSymbolInfoLoader(name string, loader SymbolInfoLoader) {
	e.symbolInfoLoaders[name] = loader
}

func (e *Exchange) Run(ctx context.Context) {
	for name, symbolInfoLoader := range e.symbolInfoLoaders {
		go func(name string, loader SymbolInfoLoader) {
			defer system.HandlePanic()
			e.loadSymbolInfo(ctx, name, loader)
			ticker := time.NewTicker(1 * time.Hour)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					e.loadSymbolInfo(ctx, name, loader)
				}
			}
		}(name, symbolInfoLoader)
	}
}

func (e *Exchange) loadSymbolInfo(ctx context.Context, name string, loader SymbolInfoLoader) {
	symbolInfo, err := loader.LoadSymbolInfo(ctx)
	if err != nil {
		zap.L().Error("load symbol info", zap.Error(err), zap.String("loader", name))
		return
	}
	errSave := backoff.Retry(func() error {
		return e.exchangeStorage.SaveSymbolInfo(ctx, symbolInfo)
	}, backoff.NewExponentialBackOff())
	if errSave != nil {
		zap.L().Error("save symbol info", zap.Error(err))
	}
}
