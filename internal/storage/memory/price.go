package memory

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"sync"
)

const DatetimeFormat = "2006-01-02 15:04:05"

var _ domain.PriceStorage = (*PriceRepository)(nil)

type PriceRepository struct {
	prices []domain.SymbolPrice
	mu     sync.Mutex
}

func NewPriceRepository() *PriceRepository {
	return &PriceRepository{
		prices: make([]domain.SymbolPrice, 0, 300_000),
	}
}

func (repo *PriceRepository) SavePrices(ctx context.Context, prices []domain.SymbolPrice) error {
	repo.mu.Lock()
	repo.prices = prices
	repo.mu.Unlock()
	return nil
}

func (repo *PriceRepository) LastPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.prices, nil
}

func (repo *PriceRepository) SymbolPrice(ctx context.Context, symbol string) ([]domain.SymbolPrice, error) {
	var res []domain.SymbolPrice
	repo.mu.Lock()
	prices := repo.prices
	repo.mu.Unlock()
	for _, price := range prices {
		if price.Symbol == symbol {
			res = append(res, price)
		}
	}
	return res, nil
}
