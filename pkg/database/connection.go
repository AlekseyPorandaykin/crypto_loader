package database

import (
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

type Config struct {
	Driver             string
	Username           string
	Password           string
	Host               string
	Port               string
	Database           string
	MaxOpenConnections int
	MaxIdleConnections int

	PathToDB string //for sqlite
}

func CreateConnection(conf Config) (*sqlx.DB, error) {
	switch conf.Driver {
	case "postgres":
		return CreatePostgresConnection(conf)
	case "sqlite":
		return CreateSqliteConnection(conf)

	default:
		return nil, fmt.Errorf("not found connection for driver: %s", conf.Driver)
	}
}

func CreatePostgresConnection(conf Config) (*sqlx.DB, error) {
	conn, err := sqlx.Connect(
		"pgx",
		fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s",
			conf.Driver,
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
		),
	)
	if err != nil {
		return nil, err
	}
	if conf.MaxOpenConnections > 0 {
		conn.SetMaxOpenConns(conf.MaxOpenConnections)
	}
	if conf.MaxIdleConnections > 0 {
		conn.SetMaxIdleConns(conf.MaxIdleConnections)
	}
	return conn, nil
}

func CreateSqliteConnection(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", filepath.Join(os.Getenv("APP_DIR"), conf.PathToDB))
	if err != nil {
		return nil, err
	}
	if errP := db.Ping(); errP != nil {
		_ = db.Close()
		return nil, errP
	}
	return db, nil
}
