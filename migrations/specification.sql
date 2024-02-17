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
--
-- alter table crypto_loader.prices
--     owner to crypto_app;

CREATE TABLE IF NOT EXISTS crypto_loader.first_symbol_price
(
    price      VARCHAR(50) NOT NULL,
    symbol     VARCHAR(50) NOT NULL,
    exchange   VARCHAR(50) NOT NULL,
    datetime   TIMESTAMP   NOT NULL,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON crypto_loader.first_symbol_price (symbol, exchange);
CREATE INDEX first_symbol_price_exchange_idx ON crypto_loader.first_symbol_price (exchange);
CREATE INDEX first_symbol_price_symbol_idx ON crypto_loader.first_symbol_price (symbol);

-- alter table crypto_loader.first_symbol_price
--     owner to crypto_app;