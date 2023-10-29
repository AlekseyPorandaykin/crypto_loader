package loaders

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
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

	storage domain.PriceStorage
}

func NewPrice(storage domain.PriceStorage) *Price {
	return &Price{
		clients: make(map[string]Client),
		prices:  make([]domain.SymbolPrice, 0, 10000),
		storage: storage,
	}
}

func (p *Price) AddClient(name string, client Client) {
	p.clients[name] = client
}

func (p *Price) Run(ctx context.Context, d time.Duration) {
	zap.L().Debug("Start loader price", zap.Int("count exchange", len(p.clients)))
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

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				p.savePrices(ctx)
			}
		}
	}()
}

func (p *Price) loadPrices(ctx context.Context, name string, client Client) {
	prices, err := client.Load(ctx)
	if err != nil {
		zap.L().Error("error load price", zap.Error(err))
		return
	}

	p.addPrices(prices)
	zap.L().Debug(
		"get prices",
		zap.String("exchange", name),
		zap.Int("count", len(prices)),
	)
}

func (p *Price) savePrices(ctx context.Context) {
	p.muPrices.Lock()
	prices := p.prices
	p.prices = make([]domain.SymbolPrice, 0, len(prices))
	p.muPrices.Unlock()

	if err := p.storage.SavePrices(ctx, prices); err != nil {
		zap.L().Error("error save prices", zap.Error(err))
	}
	zap.L().Debug("saved prices", zap.Int("count", len(prices)))
}

func (p *Price) addPrices(prices []domain.SymbolPrice) {
	p.muPrices.Lock()
	defer p.muPrices.Unlock()
	p.prices = append(p.prices, prices...)
}
