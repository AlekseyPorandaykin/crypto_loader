CREATE TABLE IF NOT EXISTS prices
(
    price      VARCHAR(50) NOT NULL,
    symbol     VARCHAR(50) NOT NULL,
    exchange   VARCHAR(50) NOT NULL,
    datetime   TIMESTAMP   NOT NULL,
    updated_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX ON prices (symbol, exchange);

CREATE INDEX exchange_idx ON prices (exchange);
CREATE INDEX symbol_idx ON prices (symbol);

alter table prices
    owner to crypto_loader;
