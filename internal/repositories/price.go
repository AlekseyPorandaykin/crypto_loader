package repositories

import (
	"context"
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/domain"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
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
		"INSERT INTO prices(price, symbol,exchange,datetime) VALUES %s ON CONFLICT (price, symbol,exchange,datetime) DO NOTHING",
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
		query  = `SELECT * FROM prices ORDER BY datetime DESC LIMIT 300000`
	)
	if err := repo.db.SelectContext(ctx, &prices, query); err != nil {
		return nil, err
	}
	return prices, nil
}

func (repo *PriceRepository) FirstDatetime(ctx context.Context) (time.Time, error) {
	var (
		query     = `SELECT min(datetime) FROM prices`
		firstDate time.Time
	)
	if err := repo.db.GetContext(ctx, &firstDate, query); err != nil {
		return time.Time{}, err
	}
	return firstDate, nil
}

func (repo *PriceRepository) Exchanges(ctx context.Context) ([]string, error) {
	var (
		exchanges []string
		query     = `SELECT DISTINCT exchange FROM prices`
	)
	if err := repo.db.SelectContext(ctx, &exchanges, query); err != nil {
		return nil, err
	}
	return exchanges, nil
}

func (repo *PriceRepository) PopularSymbols(ctx context.Context, limit int) ([]string, error) {
	var (
		query = `
SELECT 
    symbol, count(exchange) 
FROM (
	SELECT DISTINCT symbol, exchange FROM prices
	 ) as temp_table 
GROUP BY symbol
`
		symbols []string
	)
	rows, err := repo.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var (
			symbol string
			count  int
		)
		if err := rows.Scan(&symbol, &count); err != nil {
			return nil, err
		}
		if count >= limit {
			symbols = append(symbols, symbol)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return symbols, nil
}

func (repo *PriceRepository) SymbolPrices(ctx context.Context, symbol string, from, to time.Time) ([]domain.SymbolPrice, error) {
	var (
		query = `
SELECT 
    price, symbol, exchange, datetime 
FROM prices 
WHERE symbol = $1 
  AND (datetime BETWEEN $2 AND $3)  
ORDER BY  datetime ASC
`
		result []domain.SymbolPrice
	)
	if err := repo.db.SelectContext(ctx, &result, query, symbol, from, to); err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *PriceRepository) DeleteOldData(ctx context.Context, to time.Time) error {
	var query = `
DELETE FROM prices WHERE datetime < $1
`
	_, err := repo.db.ExecContext(ctx, query, to)

	return err
}
