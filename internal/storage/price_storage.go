package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/AlekseyPorandaykin/crypto_loader/internal/repositories"
	"github.com/VictoriaMetrics/fastcache"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"path"
)

var _ domain.PriceStorage = (*PriceStorage)(nil)

var keyPriceStorage = []byte("price_cache")

type PriceStorage struct {
	repo       domain.PriceStorage
	cache      *fastcache.Cache
	pathToFile string
}

func NewPriceStorage(repo *repositories.PriceRepository, pathToDir string) domain.PriceStorage {
	pathToFile := path.Join(pathToDir, "symbol_prices")
	cache := fastcache.LoadFromFileOrNew(pathToFile, 1024*1024*500)
	return &PriceStorage{repo: repo, cache: cache, pathToFile: pathToFile}
}

func (p *PriceStorage) SavePrices(ctx context.Context, prices []domain.SymbolPrice) error {
	if err := p.repo.SavePrices(ctx, prices); err != nil {
		return errors.Wrap(err, "save prices int repository")
	}
	if err := p.saveLastPrices(p.filterLastPrices(prices)); err != nil {
		return errors.Wrap(err, "save last price in cache")
	}
	return nil
}

func (p *PriceStorage) LastPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	var symbolPrices []domain.SymbolPrice
	cache := fastcache.LoadFromFileOrNew(p.pathToFile, 1024*1024*500)
	lpCache := cache.GetBig(nil, keyPriceStorage)
	if lpCache != nil {
		if err := json.Unmarshal(lpCache, &symbolPrices); err != nil {
			zap.L().Error("error unmarshal symbolPrices from cache", zap.Error(err))
		}
		if len(symbolPrices) > 0 {
			return symbolPrices, nil
		}
	}
	symbolPrices, err := p.repo.LastPrices(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "get last prices")
	}
	uniqPrices := p.filterLastPrices(symbolPrices)
	if err := p.saveLastPrices(uniqPrices); err != nil {
		return nil, errors.Wrap(err, "save last price in cache")
	}
	return uniqPrices, nil
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

func (p *PriceStorage) saveLastPrices(prices []domain.SymbolPrice) error {
	data, err := json.Marshal(prices)
	if err != nil {
		return errors.Wrap(err, "Marshal prices")
	}
	p.cache.SetBig(keyPriceStorage, data)
	if err := p.cache.SaveToFile(p.pathToFile); err != nil {
		return errors.Wrap(err, "save cache to file")
	}
	return nil
}
