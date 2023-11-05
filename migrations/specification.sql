CREATE SCHEMA IF NOT EXISTS crypto_loader;

CREATE TABLE IF NOT EXISTS crypto_loader.prices
(
    price      VARCHAR(50) NOT NULL,
    symbol     VARCHAR(50) NOT NULL,
    exchange   VARCHAR(50) NOT NULL,
    datetime   TIMESTAMP   NOT NULL,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON crypto_loader.prices (symbol, exchange);

CREATE INDEX exchange_idx ON crypto_loader.prices (exchange);
CREATE INDEX symbol_idx ON crypto_loader.prices (symbol);

alter table crypto_loader.prices
    owner to crypto_app;
