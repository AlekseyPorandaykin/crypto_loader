package repositories

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/jmoiron/sqlx"
	"strings"
)

const DatetimeFormat = "2006-01-02 15:04:05"

var _ domain.PriceStorage = (*PriceRepository)(nil)

type PriceRepository struct {
	db *sqlx.DB
}

func NewPriceRepository(db *sqlx.DB) *PriceRepository {
	return &PriceRepository{
		db: db,
	}
}

func (repo *PriceRepository) SavePrices(ctx context.Context, prices []domain.SymbolPrice) error {
	var (
		values []string
	)

	if len(prices) == 0 {
		return nil
	}
	for _, price := range prices {
		values = append(
			values,
			fmt.Sprintf(
				"('%s','%s', '%s','%s')",
				price.Price, price.Symbol, price.Exchange, price.Date.Format(DatetimeFormat)),
		)
	}
	query := fmt.Sprintf(
		"INSERT INTO prices(price, symbol,exchange,datetime) VALUES %s ON CONFLICT (symbol, exchange) DO UPDATE SET price=EXCLUDED.price, datetime=EXCLUDED.datetime, updated_at=NOW()",
		strings.Join(values, ", "),
	)
	_, err := repo.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PriceRepository) LastPrices(ctx context.Context) ([]domain.SymbolPrice, error) {
	var (
		prices = make([]domain.SymbolPrice, 0, 300_000)
		query  = `SELECT * FROM prices ORDER BY datetime`
	)
	if err := repo.db.SelectContext(ctx, &prices, query); err != nil {
		return nil, err
	}
	return prices, nil
}
