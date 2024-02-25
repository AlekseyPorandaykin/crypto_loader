package database

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Database interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
}

type MetricDB struct {
	db Database
}

func NewMetricDB(db Database) *MetricDB {
	return &MetricDB{db: db}
}

func (d *MetricDB) NamedExecContext(ctx context.Context, name, query string, arg interface{}) (sql.Result, error) {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	res, err := d.db.NamedExecContext(ctx, query, arg)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return res, err
}

func (d *MetricDB) GetContext(ctx context.Context, name string, dest interface{}, query string, args ...interface{}) error {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	err := d.db.GetContext(ctx, dest, query, args...)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return err
}

func (d *MetricDB) SelectContext(ctx context.Context, name string, dest interface{}, query string, args ...interface{}) error {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	err := d.db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return err
}

func (d *MetricDB) ExecContext(ctx context.Context, name string, query string, args ...any) (sql.Result, error) {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	res, err := d.db.ExecContext(ctx, query, args...)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return res, err
}

func (d *MetricDB) QueryxContext(ctx context.Context, name string, query string, args ...interface{}) (*sqlx.Rows, error) {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	rows, err := d.db.QueryxContext(ctx, query, args...)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return rows, err
}

func (d *MetricDB) NamedQueryContext(ctx context.Context, name, query string, arg interface{}) (*sqlx.Rows, error) {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	rows, err := d.db.NamedQueryContext(ctx, query, arg)
	if err != nil {
		IncErrorQueryDB(name)
	}
	return rows, err
}

func (d *MetricDB) QueryRowxContext(ctx context.Context, name, query string, args ...interface{}) *sqlx.Row {
	defer func(start time.Time) {
		DurationExecuteQueryDB(name, time.Since(start))
	}(time.Now())
	IncCountQueryDB(name)
	return d.db.QueryRowxContext(ctx, query, args...)
}
