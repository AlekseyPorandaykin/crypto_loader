package storage

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/metric"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sync"
	"time"
)

var _ domain.PriceStorage = (*Price)(nil)

type Price struct {
	repo domain.PriceStorage

	prices   []domain.SymbolPrice
	muPrices sync.Mutex

	lastPrices  map[string]map[string]domain.SymbolPrice
	muLastPrice sync.Mutex
}

func NewPriceStorage(repo domain.PriceStorage) *Price {
	return &Price{
		repo:       repo,
		lastPrices: make(map[string]map[string]domain.SymbolPrice),
	}
}

func (p *Price) Run(ctx context.Context) {
	go p.runUpdatePrices(ctx)
}

func (p *Price) runUpdatePrices(ctx context.Context) {
	p.UpdatePrices(ctx)
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

func (p *Price) UpdatePrices(ctx context.Context) {
	p.muPrices.Lock()
	prices := p.prices
	p.prices = make([]domain.SymbolPrice, 0, len(prices))
	p.muPrices.Unlock()

	if err := p.repo.SavePrices(ctx, p.filterLastPrices(prices)); err != nil {
		zap.L().Error("error save prices", zap.Error(err))
	}
	metric.PriceSaved.Add(float64(len(prices)))
}

func (p *Price) AddPrices(prices []domain.SymbolPrice) {
	prices = p.mutatePrice(prices)
	p.muPrices.Lock()
	p.prices = append(p.prices, prices...)
	p.muPrices.Unlock()

	p.muLastPrice.Lock()
	defer p.muLastPrice.Unlock()
	for _, price := range prices {
		if _, has := p.lastPrices[price.Exchange]; !has {
			p.lastPrices[price.Exchange] = make(map[string]domain.SymbolPrice)
		}
		p.lastPrices[price.Exchange][price.Symbol] = price
	}
}
func (p *Price) SavePrices(ctx context.Context, prices []domain.SymbolPrice) error {
	p.AddPrices(prices)
	if err := p.repo.SavePrices(ctx, p.mutatePrice(prices)); err != nil {
		zap.L().Error("save prices", zap.Error(err))
	}
	return nil
}

func (p *Price) LastPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	var symbolPrices []domain.SymbolPrice
	p.muLastPrice.Lock()
	for _, exchangePrices := range p.lastPrices {
		for _, lp := range exchangePrices {
			symbolPrices = append(symbolPrices, lp)
		}
	}
	p.muLastPrice.Unlock()
	if len(symbolPrices) > 0 {
		return symbolPrices, nil
	}
	symbolPrices, err := p.repo.LastPrices(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get last prices")
	}
	return symbolPrices, nil
}

func (p *Price) LastPrice(ctx context.Context, exchange, symbol string) (domain.SymbolPrice, error) {
	prices, err := p.LastPrices(ctx)
	if err != nil {
		return domain.SymbolPrice{}, err
	}
	for _, price := range prices {
		if price.Exchange == exchange && price.Symbol == symbol {
			return price, nil
		}
	}

	return domain.SymbolPrice{}, nil
}

func (p *Price) SymbolPrice(ctx context.Context, symbol string) ([]domain.SymbolPrice, error) {
	var symbolPrices []domain.SymbolPrice
	p.muLastPrice.Lock()
	for _, exchangePrices := range p.lastPrices {
		for _, price := range exchangePrices {
			if price.Symbol == symbol {
				symbolPrices = append(symbolPrices, price)
			}
		}
	}
	p.muLastPrice.Unlock()
	if len(symbolPrices) > 0 {
		return symbolPrices, nil
	}
	symbolPrices, err := p.repo.SymbolPrice(ctx, symbol)
	if err != nil {
		return nil, errors.Wrap(err, "get symbol prices")
	}
	return symbolPrices, nil
}

func (p *Price) ExchangePrice(ctx context.Context, exchange string) ([]domain.SymbolPrice, error) {
	var symbolPrices []domain.SymbolPrice
	p.muLastPrice.Lock()
	for _, price := range p.lastPrices[exchange] {
		symbolPrices = append(symbolPrices, price)
	}
	p.muLastPrice.Unlock()
	if len(symbolPrices) > 0 {
		return symbolPrices, nil
	}
	symbolPrices, err := p.repo.ExchangePrice(ctx, exchange)
	if err != nil {
		return nil, errors.Wrap(err, "get exchange prices")
	}
	return symbolPrices, nil
}

func (p *Price) filterLastPrices(prices []domain.SymbolPrice) []domain.SymbolPrice {
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

func (p *Price) mutatePrice(data []domain.SymbolPrice) []domain.SymbolPrice {
	result := make([]domain.SymbolPrice, 0, len(data))
	for _, item := range data {
		price := item
		price.Date = price.Date.In(time.UTC)
		price.UpdatedAt = time.Now().In(time.UTC)
		result = append(result, price)
	}

	return result
}
