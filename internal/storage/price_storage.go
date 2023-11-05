package storage

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"path"
	"sync"
	"time"
)

var _ domain.PriceStorage = (*PriceStorage)(nil)

type PriceStorage struct {
	repo       domain.PriceStorage
	pathToFile string

	muPrices   sync.Mutex
	prices     []domain.SymbolPrice
	lastPrices map[string]domain.SymbolPrice
}

func NewPriceStorage(repo *repositories.PriceRepository, pathToDir string) *PriceStorage {
	pathToFile := path.Join(pathToDir, "symbol_prices")
	return &PriceStorage{
		repo:       repo,
		pathToFile: pathToFile,
		lastPrices: make(map[string]domain.SymbolPrice),
	}
}

func (p *PriceStorage) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.UpdatePrices(ctx)
		}
	}
}

func (p *PriceStorage) UpdatePrices(ctx context.Context) {
	p.muPrices.Lock()
	prices := p.prices
	p.prices = make([]domain.SymbolPrice, 0, len(prices))
	p.muPrices.Unlock()

	if err := p.repo.SavePrices(ctx, p.filterLastPrices(prices)); err != nil {
		zap.L().Error("error save prices", zap.Error(err))
	}
	zap.L().Debug("saved prices", zap.Int("count", len(prices)))
}

func (p *PriceStorage) AddPrices(prices []domain.SymbolPrice) {
	prices = p.mutatePrice(prices)
	p.muPrices.Lock()
	p.prices = append(p.prices, prices...)
	defer p.muPrices.Unlock()

	for _, price := range prices {
		key := fmt.Sprintf("%s-%s", price.Exchange, price.Symbol)
		p.lastPrices[key] = price
	}
}
func (p *PriceStorage) SavePrices(ctx context.Context, prices []domain.SymbolPrice) error {
	p.AddPrices(prices)
	if err := p.repo.SavePrices(ctx, p.mutatePrice(prices)); err != nil {
		zap.L().Error("save prices", zap.Error(err))
	}
	return nil
}

func (p *PriceStorage) LastPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	var symbolPrices []domain.SymbolPrice
	p.muPrices.Lock()
	for _, lp := range p.lastPrices {
		symbolPrices = append(symbolPrices, lp)
	}
	p.muPrices.Unlock()
	if len(symbolPrices) > 0 {
		return symbolPrices, nil
	}
	symbolPrices, err := p.repo.LastPrices(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get last prices")
	}
	return symbolPrices, nil
}

func (p *PriceStorage) filterLastPrices(prices []domain.SymbolPrice) []domain.SymbolPrice {
	uniq := make(map[string]domain.SymbolPrice)
	for _, price := range prices {
		key := fmt.Sprintf("%s-%s", price.Exchange, price.Symbol)
		sp, ok := uniq[key]
		if !ok {
			uniq[key] = price
			continue
		}
		if sp.Date.Before(price.Date) {
			uniq[key] = price
		}
	}

	result := make([]domain.SymbolPrice, 0, len(uniq))
	for _, val := range uniq {
		result = append(result, val)
	}
	return result
}

func (p *PriceStorage) mutatePrice(data []domain.SymbolPrice) []domain.SymbolPrice {
	result := make([]domain.SymbolPrice, 0, len(data))
	for _, item := range data {
		price := item
		price.Date = price.Date.In(time.UTC)
		result = append(result, price)
	}

	return result
}
