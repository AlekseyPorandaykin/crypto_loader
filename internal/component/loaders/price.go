package loaders

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/metric"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/storage"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Client interface {
	Load(ctx context.Context) ([]domain.SymbolPrice, error)
}

type Price struct {
	clients map[string]Client

	muPrices sync.Mutex
	prices   []domain.SymbolPrice

	priceStorage  *storage.Price
	symbolStorage *storage.Symbol
}

func NewPrice(priceStorage *storage.Price, symbolStorage *storage.Symbol) *Price {
	return &Price{
		clients:       make(map[string]Client),
		prices:        make([]domain.SymbolPrice, 0, 10000),
		priceStorage:  priceStorage,
		symbolStorage: symbolStorage,
	}
}

func (p *Price) AddClient(name string, client Client) {
	p.clients[name] = client
}

func (p *Price) Run(ctx context.Context, d time.Duration) {
	for name, client := range p.clients {
		go func(name string, client Client) {
			p.loadPrices(ctx, name, client)
			ticker := time.NewTicker(d)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					p.loadPrices(ctx, name, client)
				}
			}
		}(name, client)
	}
}

func (p *Price) loadPrices(ctx context.Context, name string, client Client) {
	start := time.Now()
	prices, err := client.Load(ctx)
	if err != nil {
		metric.Errors.WithLabelValues("load_prices", name).Inc()
		zap.L().Error("error load price", zap.Error(err))
		return
	}
	metric.PriceLoadDuration.WithLabelValues(name).Add(float64(time.Since(start).Milliseconds()))
	metric.PriceLoaded.WithLabelValues(name).Add(float64(len(prices)))
	p.priceStorage.AddPrices(prices)

	symbols := make([]string, 0, len(prices))
	uniqSymbols := make(map[string]bool, len(prices))
	for _, price := range prices {
		if uniqSymbols[price.Symbol] {
			continue
		}
		symbols = append(symbols, price.Symbol)
		uniqSymbols[price.Symbol] = true
	}
	if err := p.symbolStorage.SaveSymbols(ctx, name, symbols); err != nil {
		zap.L().Error("error load symbols", zap.Error(err))
	}
}
