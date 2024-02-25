package memory

import (
	"context"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"sync"
)

var _ domain.ExchangeStorage = (*ExchangeRepository)(nil)

type ExchangeRepository struct {
	symbolData   map[string]domain.SymbolInfo
	muSymbolData sync.Mutex
}

func NewExchangeRepository() *ExchangeRepository {
	return &ExchangeRepository{symbolData: make(map[string]domain.SymbolInfo)}
}

func (e *ExchangeRepository) SaveSymbolInfo(ctx context.Context, data []domain.SymbolInfo) error {
	e.muSymbolData.Lock()
	defer e.muSymbolData.Unlock()
	for _, item := range data {
		e.symbolData[item.Symbol] = item
	}

	return nil
}

func (e *ExchangeRepository) InfoBySymbol(ctx context.Context, symbol string) (domain.SymbolInfo, error) {
	e.muSymbolData.Lock()
	defer e.muSymbolData.Unlock()
	return e.symbolData[symbol], nil
}
