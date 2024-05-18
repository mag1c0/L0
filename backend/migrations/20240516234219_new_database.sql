-- +goose Up
BEGIN;
DROP TABLE IF EXISTS order_deliveries CASCADE;
DROP TABLE IF EXISTS order_payments CASCADE;
DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;

DROP TYPE IF EXISTS locale_type;
DROP TYPE IF EXISTS currency_type;

CREATE TYPE locale_type AS ENUM (
    'ru',
    'en'
    );

CREATE TYPE currency_type AS ENUM (
    'RUB',
    'USD'
    );

create table orders
(
    order_uid          VARCHAR(50) PRIMARY KEY,
    track_number       VARCHAR(100),
    entry              VARCHAR(100),
    locale             locale_type,
    internal_signature VARCHAR(100),
    customer_id        VARCHAR(100),
    delivery_service   VARCHAR(100),
    shardkey           VARCHAR(100),
    sm_id              INTEGER,
    date_created       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    oof_shard          VARCHAR(100)
);

create table order_deliveries
(
    order_uid VARCHAR(50) PRIMARY KEY REFERENCES orders (order_uid),
    name      VARCHAR(100),
    phone     VARCHAR(100),
    zip       VARCHAR(100),
    city      VARCHAR(100),
    address   VARCHAR(100),
    region    VARCHAR(100),
    email     VARCHAR(100)
);

create table order_payments
(
    transaction   VARCHAR(50) PRIMARY KEY REFERENCES orders (order_uid),
    request_id    VARCHAR(100),
    currency      currency_type,
    provider      VARCHAR(100),
    amount        INTEGER,
    payment_dt    BIGINT,
    bank          VARCHAR(100),
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER
);

create table order_items
(
    id           SERIAL PRIMARY KEY,
    order_uid    VARCHAR(50) REFERENCES orders (order_uid),
    chrt_id      INTEGER,
    track_number VARCHAR(100),
    price        INTEGER,
    rid          VARCHAR(100),
    name         VARCHAR(100),
    sale         INTEGER,
    size         VARCHAR(100),
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR(100),
    status       INTEGER
);

COMMIT;

-- +goose Down
drop table order_deliveries;
drop table order_payments;
drop table order_items;
drop table orders;

DROP TYPE IF EXISTS locale_type;
DROP TYPE IF EXISTS currency_type;