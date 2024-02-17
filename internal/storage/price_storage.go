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

type FirstPriceLoader interface {
	FirstSymbolPrices(ctx context.Context) ([]domain.SymbolPrice, error)
}

type Price struct {
	repo             domain.PriceStorage
	firstPriceLoader FirstPriceLoader

	prices   []domain.SymbolPrice
	muPrices sync.Mutex

	lastPrices map[string]domain.SymbolPrice

	firstPrices   map[string]map[string]domain.SymbolPrice
	muFirstPrices sync.Mutex

	newPricesCh chan domain.SymbolPrice
}

func NewPriceStorage(repo domain.PriceStorage, firstPriceRepo FirstPriceLoader) *Price {
	return &Price{
		repo:             repo,
		firstPriceLoader: firstPriceRepo,
		lastPrices:       make(map[string]domain.SymbolPrice),
		firstPrices:      make(map[string]map[string]domain.SymbolPrice),
		newPricesCh:      make(chan domain.SymbolPrice),
	}
}

func (p *Price) NewPrices() <-chan domain.SymbolPrice {
	return p.newPricesCh
}

func (p *Price) Run(ctx context.Context) {
	go p.runUpdatePrices(ctx)
	go p.runLoadFirstPrices(ctx)
}

func (p *Price) runUpdatePrices(ctx context.Context) {
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

func (p *Price) runLoadFirstPrices(ctx context.Context) {
	p.loadFirstPrices(ctx)
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.loadFirstPrices(ctx)
		}
	}
}

func (p *Price) loadFirstPrices(ctx context.Context) {
	prices, err := p.firstPriceLoader.FirstSymbolPrices(ctx)
	if err != nil {
		return
	}
	p.muFirstPrices.Lock()
	defer p.muFirstPrices.Unlock()
	for _, price := range prices {
		if _, has := p.firstPrices[price.Exchange]; !has {
			p.firstPrices[price.Exchange] = make(map[string]domain.SymbolPrice)
		}
		p.firstPrices[price.Exchange][price.Symbol] = price
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
	defer p.muPrices.Unlock()

	for _, price := range prices {
		key := fmt.Sprintf("%s-%s", price.Exchange, price.Symbol)
		p.lastPrices[key] = price
	}
	_ = p.SaveFirstSymbolPrices(context.TODO(), prices)
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
	p.muPrices.Lock()
	for _, lp := range p.lastPrices {
		if lp.Symbol == symbol {
			symbolPrices = append(symbolPrices, lp)
		}
	}
	p.muPrices.Unlock()
	if len(symbolPrices) > 0 {
		return symbolPrices, nil
	}
	symbolPrices, err := p.repo.SymbolPrice(ctx, symbol)
	if err != nil {
		return nil, errors.Wrap(err, "get symbol prices")
	}
	return symbolPrices, nil
}

func (p *Price) SaveFirstSymbolPrices(ctx context.Context, prices []domain.SymbolPrice) error {
	p.muFirstPrices.Lock()
	defer p.muFirstPrices.Unlock()
	for _, price := range prices {
		if _, has := p.firstPrices[price.Exchange]; !has {
			p.firstPrices[price.Exchange] = make(map[string]domain.SymbolPrice)
		}
		if _, has := p.firstPrices[price.Exchange][price.Symbol]; has {
			continue
		}
		p.firstPrices[price.Exchange][price.Symbol] = price
		select {
		case p.newPricesCh <- price:
		default:
			continue
		}
	}

	return p.repo.SaveFirstSymbolPrices(ctx, prices)
}

func (p *Price) FirstSymbolPrice(ctx context.Context, exchange, symbol string) (domain.SymbolPrice, error) {
	p.muFirstPrices.Lock()
	defer p.muFirstPrices.Unlock()
	if _, has := p.firstPrices[exchange]; !has {
		return domain.SymbolPrice{}, nil
	}
	if price, has := p.firstPrices[exchange][symbol]; has {
		return price, nil
	}

	return p.repo.FirstSymbolPrice(ctx, exchange, symbol)
}
func (p *Price) FirstSymbolPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	return p.firstPriceLoader.FirstSymbolPrices(ctx)
}

func (p *Price) HasFirstSymbolPrice(ctx context.Context, exchange, symbol string) bool {
	p.muFirstPrices.Lock()
	defer p.muFirstPrices.Unlock()
	if _, has := p.firstPrices[symbol]; has {
		return true
	}
	return false
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
